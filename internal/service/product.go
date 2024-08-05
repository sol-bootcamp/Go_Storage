package service

import (
	"clases/internal/domain"
	"clases/internal/repository"
	"fmt"
)

// ProductService is the interface that provides product methods
type ProductService interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(int) (domain.Product, error)
	CreateProduct(domain.Product) error
	DeleteProduct(int) error
	UpdateProduct(int, domain.Product) error
}

// productService is a concrete implementation of ProductService
type productService struct {
	repository repository.ProductRepository
}

// NewProductService creates a new ProductService with the necessary dependencies
func NewProductService(repository repository.ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}

func (ps *productService) GetAllProducts() ([]domain.Product, error) {
	product, err := ps.repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}
	return product, nil
}

func (ps *productService) GetProductByID(id int) (domain.Product, error) {
	product, err := ps.repository.GetProductByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (ps *productService) CreateProduct(product domain.Product) error {
	return ps.repository.CreateProduct(product)
}

func (ps *productService) DeleteProduct(id int) error {
	err := ps.repository.DeleteProduct(id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}
	return nil
}

func (ps *productService) UpdateProduct(id int, product domain.Product) error {
	return ps.repository.UpdateProduct(id, product)
}
