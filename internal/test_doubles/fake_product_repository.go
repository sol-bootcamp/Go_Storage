package test_doubles

import (
	"clases/internal/domain"
	"errors"
)

type FakeProductRepository struct {
	Products map[int]domain.Product
	lastID   int
}

func NewFakeProductRepository() *FakeProductRepository {
	return &FakeProductRepository{
		Products: make(map[int]domain.Product),
		lastID:   0,
	}
}

func (f *FakeProductRepository) GetAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	for _, product := range f.Products {
		products = append(products, product)
	}
	return products, nil
}

func (f *FakeProductRepository) GetProductByID(id int) (domain.Product, error) {
	product, exists := f.Products[id]
	if !exists {
		return domain.Product{}, errors.New("product not found")
	}
	return product, nil
}

func (f *FakeProductRepository) CreateProduct(product domain.Product) error {
	f.lastID++
	product.ID = f.lastID
	f.Products[f.lastID] = product
	return nil
}

func (f *FakeProductRepository) DeleteProduct(id int) error {
	if _, exists := f.Products[id]; !exists {
		return errors.New("product not found")
	}
	delete(f.Products, id)
	return nil
}

func (f *FakeProductRepository) UpdateProduct(id int, product domain.Product) error {
	if _, exists := f.Products[id]; !exists {
		return errors.New("product not found")
	}
	product.ID = id
	f.Products[id] = product
	return nil
}
