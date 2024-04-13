// Package services contains domain business rules.
package services

import (
	"errors"

	"github.com/96solutions/neurography/knowledgebase/commands/domain/contracts"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/models"
)

const minCategoryNameLength = 1

// CategoryService is a set of business rules & actions related to the Category.
type CategoryService struct {
	repo contracts.CategoriesRepo
}

// NewCategoryService function makes new instance of CategoryService.
func NewCategoryService(repo contracts.CategoriesRepo) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

// CreateOrGetCategory functions creates new models.Category or returns existing.
func (s *CategoryService) CreateOrGetCategory(name string) (*models.Category, error) {
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
func (s *CategoryService) DeleteCategory(name string) error {
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
