package controllers

import (
	"github.com/binhosemcrause/webapi-with-go/models"
	"github.com/gin-gonic/gin"
)

func AllBooks(c *gin.Context) {
	allBooks := models.ShowAllBooks()
	c.JSON(200, allBooks)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	book := models.ShowBook(id)
	c.JSON(200, book)
}

/*
func NewBook(c *gin.Context) {
	var b models.Book

	err := c.ShouldBindJSON(&b)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	println("teste do teste", c.Request.Body)

	//newBook := models.CreateBook()
}
*/
