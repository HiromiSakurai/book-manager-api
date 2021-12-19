package controller

import (
	"net/http"
	"strconv"

	"github.com/HiromiSakurai/book-manager-api/model"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusNotImplemented, gin.H{"message": "[Get]/books Not implemented"})
}

func CreateBooks(c *gin.Context) {
	book := &model.Book{}
	var err error

	if err = c.BindJSON(book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err = book.Save(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusCreated, book)
	}
}

func GetBookByID(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.ParseUint(idstr, 10, 64)

	book := &model.Book{ID: id}

	if err := book.Get(); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	} else {
		c.IndentedJSON((http.StatusOK), book)
	}
}
