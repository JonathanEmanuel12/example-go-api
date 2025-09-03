package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("hello-world", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	server.GET("functions/calculate-imc", func(ctx *gin.Context) {
		height, err := strconv.ParseFloat(ctx.Query("height"), 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid height parameter",
			})
			return
		}

		weight, err := strconv.ParseFloat(ctx.Query("weight"), 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid weight parameter",
			})
			return
		}

		imc := weight / (height * height)

		ctx.JSON(200, gin.H{
			"result": imc,
		})
	})

	server.Run(":8000")
}
