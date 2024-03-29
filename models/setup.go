package models

import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDatabase(){
	database, err:= gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi"))
	if err != nil{
		panic(err)
	}
	if err := database.AutoMigrate(&Pasien{}); err != nil {
		panic(err)
	}

	// Melakukan migrasi tabel RumahSakit
	if err := database.AutoMigrate(&RumahSakit{}); err != nil {
		panic(err)
	}

	DB = database
}