package config

// mengimpor dua package
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// deklarasi sebuah fungsi bernama InitDB.
// Fungsi ini akan mengembalikan objek *gorm.DB yang akan digunakan untuk berinteraksi dengan database.
func InitDB() *gorm.DB {

	//deklarasi dan inisialisasi variabel yang akan digunakan untuk mengonfigurasi koneksi ke database MySQL.
	host := "localhost"
	port := "3306"
	dbname := "restapi-golang"
	username := "root"
	password := ""

	//string yang berisi Data Source Name (DSN) untuk koneksi ke database MySQL
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname

	// Kode ini menggunakan gorm.Open untuk membuka koneksi ke database MySQL dengan menggunakan driver MySQL yang telah diimpor sebelumnya.
	//Juga, ini mengkonfigurasi GORM dengan &gorm.Config{SkipDefaultTransaction: true}.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	////Jika ada kesalahan dalam koneksi, program akan panic dan keluar.
	if err != nil {
		panic("Can't connect to database")
	}

	//Jika koneksi ke database berhasil, fungsi ini akan mengembalikan objek *gorm.DB yang dapat digunakan dalam aplikasi untuk menjalankan query dan berinteraksi dengan database MySQL.
	return db

}
