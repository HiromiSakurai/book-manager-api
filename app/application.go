package app

import (
	"github.com/HiromiSakurai/book-manager-api/database"
	"github.com/HiromiSakurai/book-manager-api/logger"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartAppliction() {
	database.SetUp()

	mapUrls()

	router.Run("localhost:8080")

	logger.Info("start the application...")
}
