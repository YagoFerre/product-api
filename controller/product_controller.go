package controller

import (
	"net/http"
	"product-api/model"
	"product-api/usecase"
	"strconv"

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
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	createdProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdProduct)
}

func (p *productController) GetProductDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Invalid id",
		})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Id should be a number",
		})
		return
	}

	product, err := p.productUseCase.GetProductDetails(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "Product not found in database",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
