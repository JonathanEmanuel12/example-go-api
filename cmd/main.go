package main

import (
	"example-go-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	functionsController := controller.NewFunctionsController()

	server.GET("hello-world", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	server.GET("functions/calculate-imc", functionsController.GetIMC)

	server.Run(":8000")
}
