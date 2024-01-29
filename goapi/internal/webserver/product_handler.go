package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rosivaldolucas/ecommerce/goapi/internal/entity"
	"github.com/rosivaldolucas/ecommerce/goapi/internal/service"
)

type WebProductHandler struct {
	ProductService service.ProductService
}

func NewWebProductHandler(productService service.ProductService) *WebProductHandler {
	return &WebProductHandler{
		ProductService: productService,
	}
}

func (webProductHandler *WebProductHandler) GetProducts(response http.ResponseWriter, request *http.Request) {
	products, err := webProductHandler.ProductService.GetProducts()

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(response).Encode(products)
}

func (webProductHandler *WebProductHandler) GetProductByCategoryID(response http.ResponseWriter, request *http.Request) {
	categoryID := chi.URLParam(request, "categoryID")

	if categoryID == "" {
		http.Error(response, "categoryID is required", http.StatusBadRequest)

		return
	}

	products, err := webProductHandler.ProductService.GetProductByCategoryID(categoryID)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(response).Encode(products)
}

func (webProductHandler *WebProductHandler) GetProduct(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		http.Error(response, "id is required", http.StatusBadRequest)

		return
	}

	product, err := webProductHandler.ProductService.GetProduct(id)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(response).Encode(product)
}

func (webProductHandler *WebProductHandler) CreateCategory(response http.ResponseWriter, request *http.Request) {
	var product entity.Product

	err := json.NewDecoder(request.Body).Decode(&product)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	resultProduct, err := webProductHandler.ProductService.CreateProduct(product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(response).Encode(resultProduct)
}
