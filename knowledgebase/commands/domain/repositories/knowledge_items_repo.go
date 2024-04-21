// Package repositories contains list of interfaces required for domain services to provide them with data.
package repositories

import "github.com/96solutions/neurography/knowledgebase/commands/domain/models"

//go:generate mockgen -package=mock -destination=../../mock/mock_knowledge_items_repo.go -source=knowledge_items_repo.go KnowledgeItemsRepo

// KnowledgeItemsRepo interface represents a list of functions required for domain services
// to work with storage.
type KnowledgeItemsRepo interface {
	Create(item *models.KnowledgeItem) (int64, error)
	Save(item *models.KnowledgeItem) error
	Delete(item *models.KnowledgeItem) error
	FindByID(id int64) (*models.KnowledgeItem, error)
}
