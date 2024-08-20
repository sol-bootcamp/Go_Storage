package repository_test

import (
	"clases/internal/domain"
	"clases/internal/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "count", "type", "price"}).
		AddRow(1, "Product 1", 10, "Type A", 9.99).
		AddRow(2, "Product 2", 20, "Type B", 19.99)

	mock.ExpectQuery("SELECT id, name, count, type, price FROM products").WillReturnRows(rows)

	repo := repository.NewProductRepository(db)

	products, err := repo.GetAllProducts()

	if err != nil {
		t.Errorf("error was not expected while getting products: %s", err)
	}

	if len(products) != 2 {
		t.Errorf("expected 2 products, got %d", len(products))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetProductByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "count", "type", "price"}).
		AddRow(1, "Product 1", 10, "Type A", 9.99)

	mock.ExpectQuery("SELECT id, name, count, type, price FROM products WHERE id = ?").
		WithArgs(1).
		WillReturnRows(rows)

	repo := repository.NewProductRepository(db)

	product, err := repo.GetProductByID(1)

	if err != nil {
		t.Errorf("error was not expected while getting product: %s", err)
	}

	if product.ID != 1 || product.Name != "Product 1" {
		t.Errorf("expected Product 1, got %v", product)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO products").
		WithArgs("New Product", 10, "Type C", 14.99).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewProductRepository(db)

	err = repo.CreateProduct(domain.Product{Name: "New Product", Count: 10, Type: "Type C", Price: 14.99})

	if err != nil {
		t.Errorf("error was not expected while creating product: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM products WHERE id = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repository.NewProductRepository(db)

	err = repo.DeleteProduct(1)

	if err != nil {
		t.Errorf("error was not expected while deleting product: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE products SET").
		WithArgs("Updated Product", 15, "Type D", 24.99, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repository.NewProductRepository(db)

	err = repo.UpdateProduct(1, domain.Product{Name: "Updated Product", Count: 15, Type: "Type D", Price: 24.99})

	if err != nil {
		t.Errorf("error was not expected while updating product: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
