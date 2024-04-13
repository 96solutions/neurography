// Package contracts contains list of interfaces required for domain services.
package contracts

import "github.com/96solutions/neurography/knowledgebase/commands/domain/models"

//go:generate mockgen -package=mock -destination=../../mock/mock_categories_repo.go -source=categories_repo.go CategoriesRepo

// CategoriesRepo interface is a set of methods required
// for services to work with models.Category and storage.
type CategoriesRepo interface {
	FindByName(name string) (*models.Category, error)
	Create(category *models.Category) (int, error)
	Delete(category *models.Category) error
}
