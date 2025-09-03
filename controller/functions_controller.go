package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type FunctionsController struct{}

func NewFunctionsController() FunctionsController {
	return FunctionsController{}
}

func (fc FunctionsController) GetIMC(ctx *gin.Context) {
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

	imc := fc.calculateIMC(height, weight)

	ctx.JSON(200, gin.H{
		"result": imc,
	})
}

func (fc FunctionsController) calculateIMC(height, weight float64) float64 {
	return weight / (height * height)
}
