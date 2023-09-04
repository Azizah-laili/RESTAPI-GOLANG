// program utama
package main

//program mengimpor tiga package
import (
	"restapi-golang/config"
	"restapi-golang/controllers"

	"github.com/labstack/echo/v4"
)

// fungsi utama dari program.
// Semua kode yang akan dijalankan saat program dimulai ditempatkan di dalam fungsi ini.
func main() {
	// panggilan fungsi InitDB() dari package config yang digunakan untuk menginisialisasi koneksi database.
	//Hasilnya disimpan dalam variabel db.
	db := config.InitDB()

	//penginisialisasian router Echo baru yang akan digunakan untuk menangani permintaan HTTP.
	route := echo.New()

	//penginisialisasian grup route dengan prefiks
	apiV1 := route.Group("api/v1/")

	// pembuatan objek itemController yang merupakan instance dari kontroler, digunakan untuk mengelola operasi CRUD (Create, Read, Update, Delete) pada entitas "item".
	itemController := controllers.NewItemController(db) //menerima koneksi database (db) sebagai argumen

	//Kemudian, beberapa route ditentukan dengan menggunakan apiV1 dan menghubungkannya dengan metode yang sesuai dari itemController.
	apiV1.POST("item/create", itemController.Create)            // route POST untuk membuat item baru.
	apiV1.PUT("item/update/:id_item", itemController.Update)    //route PUT untuk memperbarui item dengan ID tertentu.
	apiV1.DELETE("item/delete/:id_item", itemController.Delete) //route DELETE untuk menghapus item dengan ID tertentu.
	apiV1.GET("item/get_all", itemController.GetAll)            //route GET untuk mengambil semua item.
	apiV1.GET("item/detail", itemController.GetById)            //oute GET untuk mengambil item berdasarkan ID.

	//memulai server dengan port 8080
	route.Start(":8080")
}
