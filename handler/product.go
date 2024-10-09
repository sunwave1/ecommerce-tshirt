package handler

import (
	"EcommerceShirt/database"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Weight    float32 `json:"weight"`
	Width     float32 `json:"width"`
	Height    float32 `json:"height"`
	Price     float32 `json:"price"`
	DeletedAt []uint8 `json:"deleted_at,omitempty"`
	CreatedAt []uint8 `json:"created_at,omitempty"`
	UpdatedAt []uint8 `json:"updated_at,omitempty"`
}

func GetProducts(context *gin.Context) {

	builder := new(database.Builder)

	var products []Product

	result := builder.Raw("SELECT * FROM products").Get(func(db *sql.DB, row *sql.Rows, columns []string) error {

		var id int
		var name string
		var weight, price, height, width float32
		var deletedAt, createdAt, updatedAt []uint8

		if errorScan := row.Scan(&id, &name, &weight, &width, &height, &price, &deletedAt, &createdAt, &updatedAt); errorScan != nil {
			return errorScan
		}

		products = append(products, Product{
			Id:        id,
			Name:      name,
			Weight:    weight,
			Width:     width,
			Height:    height,
			Price:     price,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})

		return nil
	})

	if result != nil {

		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on search",
		})

		return
	}

	context.JSON(200, products)
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
