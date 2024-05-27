package model

import "gorm.io/gorm"

type Patron struct {
	gorm.Model
	Name         string
	ConstactInfo string
}
