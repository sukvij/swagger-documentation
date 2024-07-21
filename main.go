package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "vijju/docs"
	"vijju/user/controller"

	"github.com/gin-gonic/gin"
)

// @title Gin Swagger Example API
func main() {
	app := gin.Default()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	controller.NewController(app)
	app.Run(":8080")
}
