package usecase

import (
	"product-api/model"
	"product-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repository repository.ProductRepository) ProductUsecase {
	return ProductUsecase{repository: repository}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	return pu.repository.CreateProduct(product)
}

func (pu *ProductUsecase) GetProductDetails(id int) (*model.Product, error) {
	return pu.repository.GetProductDetails(id)
}
