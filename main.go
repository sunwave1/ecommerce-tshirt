package main

import (
	"EcommerceShirt/database"
	router "EcommerceShirt/route"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	engine := gin.Default()

	router.InitializeRoutes(engine)

	result := database.InitializeDatabase()

	if result != nil {
		log.Fatal(result)
	}

	engine.Run()

	defer database.GetConnection().Close()

}
