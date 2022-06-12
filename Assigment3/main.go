package main

import (
	"Assignment3/config"
	"Assignment3/controllers"
	"Assignment3/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	webService := services.NewWebService()
	webController := controllers.NewWebController(webService)

	router.GET("/status", webController.GetStatus)
	router.Static("/js", "./js")

	router.Run(config.APP_PORT)
}
