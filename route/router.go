package router

import (
	"EcommerceShirt/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(engine *gin.Engine) {

	groupRoutes := engine.Group("/api/v1")

	groupRoutes.GET("/products", handler.GetProducts)
	groupRoutes.POST("/product/:id", handler.GetProduct)
	groupRoutes.DELETE("/product/:id", handler.DeleteProduct)
	groupRoutes.PUT("/product", handler.UpdateProduct)

}
