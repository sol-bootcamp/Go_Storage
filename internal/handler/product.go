package handler

import (
	"clases/internal/domain"
	"clases/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var ErrInvalidID = "invalid id"

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (ph *ProductHandler) GetAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		products, err := ph.service.GetAllProducts()
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)

	}
}

func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, ErrInvalidID, 400)
		return
	}

	product, err := ph.service.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct domain.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := newProduct.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = ph.service.CreateProduct(newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, ErrInvalidID, 400)
		return
	}

	err = ph.service.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "product deleted"}`))
}

func (ph *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, ErrInvalidID, 400)
		return
	}

	var newProduct domain.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = newProduct.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = ph.service.UpdateProduct(id, newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
