package service

import (
	"github.com/rosivaldolucas/ecommerce/goapi/internal/database"
	"github.com/rosivaldolucas/ecommerce/goapi/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (categoryService *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := categoryService.CategoryDB.GetCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (categoryService *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := categoryService.CategoryDB.GetCategory(id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (categoryService *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)

	_, err := categoryService.CategoryDB.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	return category, nil
}
