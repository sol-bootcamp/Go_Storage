package service_test

import (
	"clases/internal/domain"
	"clases/internal/service"
	"clases/internal/test_doubles"
	"testing"
)

func TestProductService_dummy(t *testing.T) {
	dummyRepo := &test_doubles.DummyProductRepository{}
	service := service.NewProductService(dummyRepo)

	_, err := service.GetAllProducts()
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestProductServiceCreate_dummy(t *testing.T) {
	dummyRepo := &test_doubles.DummyProductRepository{}
	service := service.NewProductService(dummyRepo)

	product := domain.Product{ID: 1, Name: "test", Type: "test", Count: 1, Price: 1}
	err := service.CreateProduct(product)
	if err != nil {
		t.Fatalf("error should be nil, got %v", err)
	}
}

// STUB

func TestProductService_stub(t *testing.T) {
	stubRepo := &test_doubles.StubProductRepository{}
	service := service.NewProductService(stubRepo)

	products, err := service.GetAllProducts()
	if err != nil {
		t.Error("error should be nil")
	}
	if len(products) != 1 {
		t.Errorf("expected 1 products, got %d", len(products))
	}

	if products[0].Name != "test" {
		t.Errorf("expected product name test, got %s", products[0].Name)
	}
}

//MOCK

func TestProductService_mock(t *testing.T) {
	mockRepo := &test_doubles.MockProductRepository{}
	service := service.NewProductService(mockRepo)

	newProduct := domain.Product{ID: 1, Name: "Test Mock", Type: "test", Count: 1, Price: 1}
	err := service.CreateProduct(newProduct)

	if err != nil {
		t.Fatalf("error should be nil, got %v", err)
	}

	if mockRepo.CreatedProducts[0].Name != "Test Mock" {
		t.Errorf("expected product name Test Mock, got %s", mockRepo.CreatedProducts[0].Name)
	}

	if len(mockRepo.CreatedProducts) != 1 {
		t.Errorf("expected 1 product, got %d", len(mockRepo.CreatedProducts))
	}

}

// FAKE

func TestProductService_fake(t *testing.T) {
	fakeRepo := &test_doubles.FakeProductRepository{
		Products: map[int]domain.Product{
			1: {ID: 1, Name: "Test Fake", Type: "test", Count: 1, Price: 1},
			2: {ID: 2, Name: "Test Fake 2", Type: "test", Count: 1, Price: 1},
		},
	}

	service := service.NewProductService(fakeRepo)
	products, err := service.GetAllProducts()

	if err != nil {
		t.Fatalf("error should be nil, got %v", err)
	}

	if len(products) != 2 {
		t.Errorf("expected 2 products, got %d", len(products))
	}

	if products[0].Name != "Test Fake" {
		t.Errorf("expected product name Test Fake, got %s", products[0].Name)
	}

}
