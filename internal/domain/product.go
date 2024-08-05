package domain

import "errors"

// Product is an struct that represents a product.
type Product struct {
	// ID is the unique identifier of the product.
	ID int
	// Name is the name of the product.
	Name string
	// Type is the type of the product.
	Type string
	// Count is the number of products.
	Count int
	// Price is the price of the product.
	Price float64
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Count <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	if p.Type == "" {
		return errors.New("type is required")
	}

	if p.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}
