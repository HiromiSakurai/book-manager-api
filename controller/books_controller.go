package controller

import (
	"net/http"
	"strconv"

	"github.com/HiromiSakurai/book-manager-api/database"
	"github.com/HiromiSakurai/book-manager-api/logger"
	"github.com/HiromiSakurai/book-manager-api/model"
	"github.com/gin-gonic/gin"
)

var books = []model.Book{
	{ID: 1, Title: "hoge"},
	{ID: 2, Title: "fuga"},
	{ID: 3, Title: "sakurai"},
	{ID: 4, Title: "hiromi"},
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func CreateBooks(c *gin.Context) {
	var newBook model.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	result := database.DB.Create(&newBook)

	if result.Error != nil {
		logger.Error("aaaas", result.Error)
	} else {
		c.IndentedJSON(http.StatusCreated, newBook)
	}
}

func GetBookByID(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.ParseUint(idstr, 10, 64)

	for _, b := range books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
