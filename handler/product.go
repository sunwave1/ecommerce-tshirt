package handler

import (
	"github.com/gin-gonic/gin"
)

func GetProducts(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Error on search",
	})
}

func GetProduct(context *gin.Context) {

	context.JSON(200, gin.H{
		"message": "Ok",
	})
}

func UpdateProduct(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Ok",
	})
}

func DeleteProduct(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Ok",
	})
}
