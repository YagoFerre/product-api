package controller

import (
	"net/http"
	"product-api/model"
	"product-api/usecase"

	"github.com/gin-gonic/gin"
)

type productController struct {
	// use cases
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Product 1",
			Price: 25,
		},
		{
			ID:    2,
			Name:  "Product 2",
			Price: 40,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
