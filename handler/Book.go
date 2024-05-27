package handler

import (
	"strconv"

	"github.com/Asad2730/LibrarySystem/model"
	"github.com/Asad2730/LibrarySystem/repo"
	"github.com/gin-gonic/gin"
)

func stringToUint(str string) uint {
	// Convert the string to uint
	num, err := strconv.ParseUint(str, 10, 0)
	if err != nil {

		return 0
	}

	// Return the uint value
	return uint(num)
}

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

func GetBook(c *gin.Context) {
	id := c.Param("id")
	uid := stringToUint(id)
	if uid == 0 {
		c.JSON(401, "Id type missMacth")
	}
	res, err := repo.GetBook(uid)
	if err != nil {
		c.JSON(402, err.Error())
		return
	}

	c.JSON(200, res)

}
