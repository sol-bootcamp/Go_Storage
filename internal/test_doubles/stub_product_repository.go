package test_doubles

import (
	"clases/internal/domain"
	"errors"
)

type StubProductRepository struct{}

func (s *StubProductRepository) GetAllProducts() ([]domain.Product, error) {
	return []domain.Product{}, errors.New("error")

}

func (d *StubProductRepository) GetProductByID(id int) (domain.Product, error) {
	return domain.Product{}, nil
}

func (d *StubProductRepository) CreateProduct(product domain.Product) error {
	return nil
}
func (d *StubProductRepository) DeleteProduct(id int) error {
	return nil
}
func (d *StubProductRepository) UpdateProduct(int, domain.Product) error {
	return nil
}
