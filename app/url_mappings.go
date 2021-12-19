package app

import "github.com/HiromiSakurai/book-manager-api/controller"

func mapUrls() {
	router.GET("/ping", controller.Ping)
	router.GET("/books", controller.GetBooks)
	router.GET("/books/:id", controller.GetBookByID)
	router.POST("/books", controller.CreateBooks)
}
