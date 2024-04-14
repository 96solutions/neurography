// Package services contains domain business rules.
package services

import (
	"errors"

	"github.com/96solutions/neurography/knowledgebase/commands/domain/models"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/repositories"
)

const minCategoryNameLength = 1

//go:generate mockgen -package=mock -destination=../../mock/mock_category_service.go -source=category_service.go CategoryService

// CategoryService represents a service that provides functionality related to the models.Category.
type CategoryService interface {
	CreateOrGetCategory(name string) (*models.Category, error)
	DeleteCategory(name string) error
}

// categoryService is a set of business rules & actions related to the Category.
type categoryService struct {
	repo repositories.CategoriesRepo
}

// NewCategoryService function makes new instance of CategoryService.
func NewCategoryService(repo repositories.CategoriesRepo) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

// CreateOrGetCategory functions creates new models.Category or returns existing.
func (s *categoryService) CreateOrGetCategory(name string) (*models.Category, error) {
	cat, err := s.repo.FindByName(name)
	if err != nil {
		return nil, err
	}

	if cat != nil {
		return cat, nil
	}

	if len(name) <= minCategoryNameLength {
		return nil, errors.New("category name is too short")
	}

	cat = &models.Category{
		Name: name,
	}

	cat.ID, err = s.repo.Create(cat)
	if err != nil {
		return nil, err
	}

	return cat, nil
}

// DeleteCategory function deletes models.Category.
func (s *categoryService) DeleteCategory(name string) error {
	cat, err := s.repo.FindByName(name)
	if err != nil {
		return err
	}

	if cat == nil {
		return errors.New("category not exists")
	}

	err = s.repo.Delete(cat)
	if err != nil {
		return err
	}

	return nil
}
