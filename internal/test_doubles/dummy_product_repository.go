package test_doubles

import "clases/internal/domain"

type DummyProductRepository struct{}

func (d *DummyProductRepository) GetAllProducts() ([]domain.Product, error) {
	return []domain.Product{}, nil
}
func (d *DummyProductRepository) GetProductByID(id int) (domain.Product, error) {
	return domain.Product{}, nil
}

func (d *DummyProductRepository) CreateProduct(product domain.Product) error {
	return nil
}
func (d *DummyProductRepository) DeleteProduct(id int) error {
	return nil
}
func (d *DummyProductRepository) UpdateProduct(int, domain.Product) error {
	return nil
}
