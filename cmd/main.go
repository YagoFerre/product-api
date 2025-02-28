package main

import (
	"product-api/controller"
	"product-api/db"
	"product-api/repository"
	"product-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	// camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// camada use cases
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)
	// camada controller
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product/create", ProductController.CreateProduct)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000")
}
