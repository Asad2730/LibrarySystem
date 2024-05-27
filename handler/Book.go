package handler

import (
	"github.com/Asad2730/LibrarySystem/model"
	"github.com/Asad2730/LibrarySystem/repo"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book model.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := repo.CreateBook(&book); err != nil {
		c.JSON(401, err.Error())
		return
	}

	c.JSON(201, "Created!")
}

func GetBooks(c *gin.Context) {
	var book model.Book
	res, err := repo.GetBooks(&book)
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	c.JSON(200, res)

}
