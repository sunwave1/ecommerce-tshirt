package main

import (
	router "EcommerceShirt/route"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	router.InitializeRoutes(engine)

	engine.Run()
}
