package internal

import (
	"github.com/Asad2730/LibrarySystem/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=123 dbname= port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&model.Book{}, &model.Patron{})
	db.AutoMigrate(&model.BorrowRecord{})

	Db = db
}
