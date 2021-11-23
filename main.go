package main

import (
	"fmt"
	"net/http"

	"github.com/HiromiSakurai/book-manager-api/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_, err := database()
	if err != nil {
		fmt.Println("DB接続失敗")
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}

	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")
}

func database() (database *gorm.DB, err error) {
	dsn := "root:@tcp(localhost:3306)/book_manager_api?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

var books = []entity.Book{
	{ID: "1", Title: "hoge"},
	{ID: "2", Title: "fuga"},
	{ID: "3", Title: "sakurai"},
	{ID: "4", Title: "hiromi"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var newBook entity.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new album to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, b := range books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
