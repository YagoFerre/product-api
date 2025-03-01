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

	product := server.Group("/api/v1")
	{
		product.GET("/products", ProductController.GetProducts)
		product.GET("/products/:id", ProductController.GetProductDetails)
		product.POST("/product/create", ProductController.CreateProduct)
	}

	server.Run(":8000")
}
