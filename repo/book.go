package repo

import (
	"github.com/Asad2730/LibrarySystem/internal"
	"github.com/Asad2730/LibrarySystem/model"
	"gorm.io/gorm"
)

func CreateBook(book *model.Book) error {

	if err := internal.Db.Create(&book).Error; err != nil {
		return err
	}

	return nil
}

func GetBooks(book *model.Book) (books *[]model.Book, err error) {
	// condition := model.Book{
	// 	Model: gorm.Model{
	// 		DeletedAt: ,
	// 	},
	// }
	if err := internal.Db.Find(book).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func GetBook(id uint) (book *model.Book, err error) {

	condition := model.Book{
		Model: gorm.Model{
			ID: id,
			//	DeletedAt:,
		},
	}
	if err := internal.Db.First(&book, condition).Error; err != nil {
		return nil, err
	}

	return book, nil
}
