package repositories

import (
	"restapi-golang/models" //berisi definisi model data

	"gorm.io/gorm" //untuk berinteraksi dengan database menggunakan ORM
)

// deklarasi tipe data baru dengan nama "dbItem".
// Tipe data ini memiliki satu field yaitu "Conn" yang merupakan pointer ke objek dari tipe "gorm.DB".
// Tipe ini akan digunakan untuk mengimplementasikan sebuah interface bernama "ItemRepository".
type dbItem struct {
	Conn *gorm.DB
}

// Create implements ItemRepository.
// implementasi dari metode "Create" yang mengambil objek dari tipe "models.Item" dan mencoba membuatnya di database menggunakan objek "gorm.DB".
// Metode ini mengembalikan error jika operasi tidak berhasil.
func (db *dbItem) Create(item models.Item) error {
	return db.Conn.Create(&item).Error
}

// Delete implements ItemRepository.
// implementasi dari metode "Delete" yang menghapus item berdasarkan ID yang diberikan dari database.
// Metode ini mengembalikan error jika operasi tidak berhasil.
func (db *dbItem) Delete(idItem int) error {
	return db.Conn.Delete(&models.Item{IdItem: idItem}).Error
}

// GetAll implements ItemRepository.
// Implementasi dari metode "GetAll" yang mengambil semua item dari database dan mengembalikannya dalam bentuk slice ([]models.Item).
// Metode ini juga mengembalikan error jika operasi tidak berhasil.
func (db *dbItem) GetAll() ([]models.Item, error) {
	var data []models.Item
	result := db.Conn.Find(&data)
	return data, result.Error
}

// GetById implements ItemRepository.
// implementasi dari metode "GetById" yang mengambil item berdasarkan ID yang diberikan dari database dan mengembalikannya dalam bentuk objek "models.Item".
// Metode ini juga mengembalikan error jika operasi tidak berhasil.
func (db *dbItem) GetById(idItem int) (models.Item, error) {
	var data models.Item
	result := db.Conn.Where("id_item", idItem).First(&data)
	return data, result.Error
}

// Update implements ItemRepository.
// implementasi dari metode "Update" yang mengupdate item berdasarkan ID yang diberikan dengan data yang baru dari objek "models.Item".
// Metode ini mengembalikan error jika operasi tidak berhasil.
func (db *dbItem) Update(idItem int, item models.Item) error {
	return db.Conn.Where("id_item", idItem).Updates(item).Error
}

// interface "ItemRepository" mendefinisikan lima method yang harus diimplementasikan oleh tipe data lain.
type ItemRepository interface {
	Create(item models.Item) error
	Update(idItem int, item models.Item) error
	Delete(idItem int) error
	GetById(idItem int) (models.Item, error)
	GetAll() ([]models.Item, error)
}

// digunakan untuk membuat sebuah objek dari tipe "ItemRepository".
// Fungsi ini menerima satu parameter yaitu pointer ke objek "gorm.DB" dan mengembalikan objek "ItemRepository" yang diinisialisasi dengan objek "dbItem"
// yang memiliki koneksi database yang diberikan.
// Dengan cara ini, dapat membuat instance dari "ItemRepository" yang sesuai dengan jenis database yang digunakan.
func NewItemRepository(Conn *gorm.DB) ItemRepository {
	return &dbItem{Conn: Conn}
}
