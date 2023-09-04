package controllers

import (
	"net/http"
	"restapi-golang/models"
	"restapi-golang/services"
	"strconv"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// mendefinisikan sebuah struct bernama ItemController yang memiliki dua atribut: itemService dan validate.
type ItemController struct {
	//Ini adalah controller untuk entitas "Item", dan itemService digunakan untuk berinteraksi dengan layanan yang terkait dengan Item,
	itemService services.ItemService

	//sedangkan validate adalah objek untuk melakukan validasi input.
	validate vl.Validate
}

// Method ini digunakan untuk menangani permintaan pembuatan item baru (HTTP POST).
func (controller ItemController) Create(c echo.Context) error {
	//Dalam metode ini, mendefinisikan struktur payload yang digunakan untuk mem-parse dan memvalidasi data yang diterima dari permintaan JSON.
	type payload struct {
		NamaItem    string  `json:"nama_item" validate:"required"`
		Unit        string  `json:"unit" validate:"required"`
		Stok        int     `json:"stok" validate:"required"`
		HargaSatuan float64 `json:"harga_satuan" validate:"required"`
	}

	//Kemudian, data yang diterima dari permintaan ke payloadValidator, memeriksa kesalahan dalam binding, dan melakukan validasi struktur data.
	//Jika semuanya berhasil, maka bisa menggunakan itemService untuk membuat item baru dalam basis data dan mengembalikan respons HTTP dengan hasilnya.
	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return err
	}

	result := controller.itemService.Create(models.Item{NamaItem: payloadValidator.NamaItem, Unit: payloadValidator.Unit, Stok: payloadValidator.Stok, HargaSatuan: payloadValidator.HargaSatuan})

	return c.JSON(http.StatusOK, result)
}

// method untuk meng-handle permintaan pembaruan item (HTTP PUT).
// Seperti Create, metode ini juga mengikuti langkah-langkah serupa, termasuk parsing data JSON, validasi, dan pembaruan item yang sesuai dalam basis data.
func (controller ItemController) Update(c echo.Context) error {
	type payload struct {
		NamaItem    string  `json:"nama_item" validate:"required"`
		Unit        string  `json:"unit" validate:"required"`
		Stok        int     `json:"stok" validate:"required"`
		HargaSatuan float64 `json:"harga_satuan" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	idItem, _ := strconv.Atoi(c.Param("id_item"))
	result := controller.itemService.Update(idItem, models.Item{NamaItem: payloadValidator.NamaItem, Unit: payloadValidator.Unit, Stok: payloadValidator.Stok, HargaSatuan: payloadValidator.HargaSatuan})

	return c.JSON(http.StatusOK, result)
}

// Mehod delete digunakan untuk menghapus item berdasarkan ID yang diberikan dalam permintaan (HTTP DELETE).
// ID diambil dari parameter URL, diubah menjadi integer, dan digunakan untuk memanggil itemService untuk menghapus item.
func (controller ItemController) Delete(c echo.Context) error {
	idItem, _ := strconv.Atoi(c.Param("id_item"))
	result := controller.itemService.Delete(idItem)

	return c.JSON(http.StatusOK, result)
}

// method yang digunakan untuk mengambil daftar semua item (HTTP GET). Metode ini memanggil itemService untuk mengambil semua item dari basis data.
func (controller ItemController) GetAll(c echo.Context) error {
	result := controller.itemService.GetAll()
	return c.JSON(http.StatusOK, result)
}

// Method digunakan untuk mengambil item berdasarkan ID yang diberikan dalam permintaan (HTTP GET).
// ID diambil dari parameter query URL, diubah menjadi integer, dan digunakan untuk memanggil itemService untuk mendapatkan item yang sesuai.
func (controller ItemController) GetById(c echo.Context) error {
	idItem, _ := strconv.Atoi(c.QueryParam("id_item"))
	result := controller.itemService.GetById(idItem)
	return c.JSON(http.StatusOK, result)
}

// fungsi pembuatan objek ItemController dimana menerima parameter berupa objek GORM database (db) yang akan digunakan untuk berinteraksi dengan basis data.
func NewItemController(db *gorm.DB) ItemController {

	//Dalam fungsi ini, membuat sebuah objek ItemService yang menggunakan db, dan kemudian objek ItemController yang baru
	//dengan objek ItemService dan objek validasi yang sesuai.
	service := services.NewItemService(db)
	controller := ItemController{
		itemService: service,
		validate:    *vl.New(),
	}

	return controller
}
