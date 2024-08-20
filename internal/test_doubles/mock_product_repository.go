package test_doubles

import "clases/internal/domain"

type MockProductRepository struct {
	CreatedProducts []domain.Product
}

func (m *MockProductRepository) GetAllProducts() ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (m *MockProductRepository) GetProductByID(id int) (domain.Product, error) {
	return domain.Product{}, nil
}

func (m *MockProductRepository) CreateProduct(product domain.Product) error {
	m.CreatedProducts = append(m.CreatedProducts, product)
	return nil
}

func (m *MockProductRepository) DeleteProduct(id int) error {
	return nil
}

func (m *MockProductRepository) UpdateProduct(id int, product domain.Product) error {
	return nil
}
