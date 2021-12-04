package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/HiromiSakurai/book-manager-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := database()
	if err != nil {
		fmt.Println("DB接続失敗")
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}

	db.Logger.LogMode(logger.LogLevel(4))

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

var books = []model.Book{
	{ID: 1, Title: "hoge"},
	{ID: 2, Title: "fuga"},
	{ID: 3, Title: "sakurai"},
	{ID: 4, Title: "hiromi"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var newBook model.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new album to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
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
