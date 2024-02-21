package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/dto"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/entity"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/infra/database"
	entityPkg "github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/pkg/entity"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

// CreateProduct godoc
// @Summary 		Create a product
// @Description 	Create a product
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			product body dto.CreateProductInput true "product request"
// @Success 		201 {object} string
// @Failure 		400 {object} string
// @Failure 		500 {object} string
// @Router 			/products [post]
// @Security 		ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetProduct godoc
// @Summary 		Get a product
// @Description 	Get a product
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "product id"
// @Success 		200 {object} entity.Product
// @Failure 		400 {object} string
// @Failure 		404 {object} string
// @Failure 		500 {object} string
// @Router 			/products/{id} [get]
// @Security 		ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(p)
}

// UpdateProduct godoc
// @Summary 		Update a product
// @Description 	Update a product
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "product id"
// @Param 			product body entity.Product true "product request"
// @Success 		204 {object} string
// @Failure 		400 {object} string
// @Failure 		404 {object} string
// @Failure 		500 {object} string
// @Router 			/products/{id} [put]
// @Security 		ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteProduct godoc
// @Summary 		Delete a product
// @Description 	Delete a product
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "product id"
// @Success 		204 {object} string
// @Failure 		400 {object} string
// @Failure 		500 {object} string
// @Router 			/products/{id} [delete]
// @Security 		ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetProducts godoc
// @Summary 		Get all products
// @Description 	Get all products
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			page query int false "page"
// @Param 			limit query int false "limit"
// @Param 			sort query string false "sort"
// @Success 		200 {object} []entity.Product
// @Failure 		400 {object} string
// @Failure 		500 {object} string
// @Router 			/products [get]
// @Security 		ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
