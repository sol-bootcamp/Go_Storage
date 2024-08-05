package repository

import (
	"clases/internal/domain"
	"database/sql"
	"errors"
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
	row := pr.db.QueryRow("SELECT id, name, count, type, price FROM products WHERE id = ?", id)
	var product domain.Product
	err := row.Scan(&product.ID, &product.Name, &product.Count, &product.Type, &product.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, errors.New("product not found")
		}
		return domain.Product{}, err
	}
	return product, nil

}

func (pr *productRepository) CreateProduct(product domain.Product) error {
	result, err := pr.db.Exec("INSERT INTO products (name, count, type, price) VALUES (?, ?, ?, ?)", product.Name, product.Count, product.Type, product.Price)
	if err != nil {
		return err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = int(lastId)

	return nil

}

func (pr *productRepository) DeleteProduct(id int) error {
	exist, err := pr.productExists(id)
	if err != nil {
		return errors.New("error checking if product exists")
	}
	if !exist {
		return errors.New("product not found")
	}
	_, err = pr.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) UpdateProduct(id int, product domain.Product) error {
	exist, err := pr.productExists(id)
	if err != nil {
		return errors.New("error checking if product exists")
	}
	if !exist {
		return errors.New("product not found")
	}

	_, err = pr.db.Exec("UPDATE products SET name = ?, count = ?, type = ?, price = ? WHERE id = ?", product.Name, product.Count, product.Type, product.Price, id)
	if err != nil {
		return err
	}
	return nil

}

func (pr *productRepository) productExists(id int) (bool, error) {
	var exist bool
	query := "SELECT EXISTS(SELECT 1 FROM products WHERE id = ?)"
	err := pr.db.QueryRow(query, id).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}
