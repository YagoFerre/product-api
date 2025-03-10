package repository

import (
	"database/sql"
	"fmt"
	"product-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{connection: connection}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price FROM product"

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObject model.Product

	for rows.Next() {
		err = rows.Scan(
			// endereço de memoria &
			&productObject.ID,
			&productObject.Name,
			&productObject.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObject)
	}

	rows.Close()
	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO product (name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id) // primeiro valor que retorna do banco atribui a var id int line 50
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	product.ID = id

	query.Close()
	return product, nil
}

func (pr *ProductRepository) GetProductDetails(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var productObject model.Product

	err = query.QueryRow(id).Scan(
		&productObject.ID,
		&productObject.Name,
		&productObject.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	query.Close()
	return &productObject, nil
}
