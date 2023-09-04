package services

import (
	"fmt" // formatting dan mencetak teks

	"restapi-golang/helpers"      //package yang berisi utility dan fungsi-fungsi bantuan (helpers) untuk aplikasi.
	"restapi-golang/models"       // definisi model (struktur data)
	"restapi-golang/repositories" //definisi repositori, yang digunakan untuk mengakses dan memanipulasi data dalam database.

	"gorm.io/gorm" //berinteraksi dengan database menggunakan GORM, yang merupakan ORM (Object-Relational Mapping) untuk Go.
)

// struct itemService memiliki satu field yaitu itemRepo yang merupakan objek dari repositories.ItemRepository.
// itemService akan digunakan untuk mengimplementasikan layanan terkait item.
type itemService struct {
	itemRepo repositories.ItemRepository
}

// Create implements ItemService.
// Method ini mengimplementasikan struct dari ItemService.
// Menerima input models.Item (sebuah struktur data yang merupakan representasi dari item) dan mengembalikan objek helpers.Response yang berisi status dan pesan.
func (service *itemService) Create(item models.Item) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.Create(item); err != nil {
		response.Status = 500
		response.Messages = "Failed to create a new item " + err.Error()
	} else {
		response.Status = 200
		response.Messages = "Success to create a new item"
	}
	return response
}

// Delete implements ItemService.
// Method ini menerima ID item sebagai input dan mengembalikan objek helpers.Response.
func (service *itemService) Delete(idItem int) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.Delete(idItem); err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to delete item : ", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success to delete item"
	}
	return response
}

// GetAll implements ItemService.
// Untuk mendapatkan semua item dalam aplikasi dan mengembalikan objek helpers.Response yang berisi data item jika berhasil.
func (service *itemService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.itemRepo.GetAll()
	if err != nil {
		response.Status = 500
		response.Messages = "Failed to get items"
	} else {
		response.Status = 200
		response.Messages = "Success to get items"
		response.Data = data
	}
	return response
}

// GetById implements ItemService.
// mendapatkan item berdasarkan ID dan mengembalikan objek helpers.Response yang berisi data item jika berhasil.
func (service *itemService) GetById(idItem int) helpers.Response {
	var response helpers.Response
	data, err := service.itemRepo.GetById(idItem)
	if err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to get item : ", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success to get item"
		response.Data = data
	}
	return response
}

// Update implements ItemService.
// memperbarui item dengan ID tertentu dan mengembalikan objek helpers.Response yang mencerminkan hasil operasi pembaruan.
func (service *itemService) Update(idItem int, item models.Item) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.Update(idItem, item); err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to update item ", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success to update item"
	}
	return response
}

type ItemService interface {
	Create(item models.Item) helpers.Response
	Update(idItem int, item models.Item) helpers.Response
	Delete(idItem int) helpers.Response
	GetById(idItem int) helpers.Response
	GetAll() helpers.Response
}

// digunakan untuk membuat objek layanan item baru.
// Fungsi ini menerima objek gorm.DB yang akan digunakan untuk berinteraksi dengan database dan mengembalikan objek ItemService yang sesuai dengan implementasi itemService.
// Ini menginisialisasi itemRepo dengan repositori item yang sesuai.
func NewItemService(db *gorm.DB) ItemService {
	return &itemService{itemRepo: repositories.NewItemRepository(db)}
}
