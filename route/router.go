package router

import "github.com/gin-gonic/gin"

func InitializeRoutes(engine *gin.Engine) {

	groupRoutes := engine.Group("/api/v1")

	groupRoutes.GET("/products", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Ok",
		})
	})

	groupRoutes.POST("/product/:id", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Ok",
		})
	})

	groupRoutes.DELETE("/product/:id", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Ok",
		})
	})

	groupRoutes.PUT("/product", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Ok",
		})
	})

}
