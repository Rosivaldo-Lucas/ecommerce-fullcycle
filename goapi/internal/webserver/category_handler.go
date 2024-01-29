package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rosivaldolucas/ecommerce/goapi/internal/entity"
	"github.com/rosivaldolucas/ecommerce/goapi/internal/service"
)

type WebCategoryHandler struct {
	CategoryService service.CategoryService
}

func NewWebCategoryHandler(categoryService service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{
		CategoryService: categoryService,
	}
}

func (webCategoryHandler *WebCategoryHandler) GetCategories(response http.ResponseWriter, request *http.Request) {
	categories, err := webCategoryHandler.CategoryService.GetCategories()

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(response).Encode(categories)
}

func (webCategoryHandler *WebCategoryHandler) GetCategory(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		http.Error(response, "id is required", http.StatusBadRequest)

		return
	}

	category, err := webCategoryHandler.CategoryService.GetCategory(id)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(response).Encode(category)
}

func (webCategoryHandler *WebCategoryHandler) CreateCategory(response http.ResponseWriter, request *http.Request) {
	var category entity.Category

	err := json.NewDecoder(request.Body).Decode(&category)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	resultCategory, err := webCategoryHandler.CategoryService.CreateCategory(category.Name)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(response).Encode(resultCategory)
}
