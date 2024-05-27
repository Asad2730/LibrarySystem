package model

import (
	"gorm.io/gorm"
)

type BorrowRecord struct {
	gorm.Model
	Book       *Book
	Patron     *Patron
	BorrowDate string
	ReturnDate string
}
