package repository

import (
	"clases/internal/domain"
	"database/sql"
	"log"
)

// ProductRepository is the interface that provides product methods
type ProductRepository interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(id int) (domain.Product, error)
	CreateProduct(product domain.Product) error
	DeleteProduct(id int) error
	UpdateProduct(int, domain.Product) error
}

// productRepository is a concrete implementation of ProductRepository
type productRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new ProductRepository with the necessary dependencies
func NewProductRepository(db *sql.DB) ProductRepository {
	repo := &productRepository{
		db: db,
	}
	return repo

}

func (pr *productRepository) GetAllProducts() ([]domain.Product, error) {
	rows, err := pr.db.Query("SELECT id, name, count,type, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Count, &product.Type, &product.Price)
		if err != nil {
			log.Println(err)
			continue
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil

}

func (pr *productRepository) GetProductByID(id int) (domain.Product, error) {
	return domain.Product{}, nil

}

func (pr *productRepository) CreateProduct(product domain.Product) error {
	return nil
}

func (pr *productRepository) DeleteProduct(id int) error {
	return nil
}

func (pr *productRepository) UpdateProduct(id int, product domain.Product) error {
	return nil
}
