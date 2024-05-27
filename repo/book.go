package repo

import (
	"github.com/Asad2730/LibrarySystem/internal"
	"github.com/Asad2730/LibrarySystem/model"
)

func CreateBook(book *model.Book) error {

	if err := internal.Db.Create(&book).Error; err != nil {
		return err
	}

	return nil
}

func GetBooks(book *model.Book) (books *[]model.Book, err error) {

	if err := internal.Db.Find(book).Error; err != nil {
		return nil, err
	}

	return books, nil
}
