// Package usecases contains a set of sequences for interactions between services and users.
package usecases

import (
	"github.com/96solutions/neurography/knowledgebase/commands/application/models"
	domain "github.com/96solutions/neurography/knowledgebase/commands/domain/models"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/services"
)

// UpdateKnowledgeItem type represents usecase that has sequence of actions to update new models.KnowledgeItem.
type UpdateKnowledgeItem struct {
	categoryService      services.CategoryService
	knowledgeItemService services.KnowledgeItemService
}

// NewUpdateKnowledgeItem function builds new instance of UpdateKnowledgeItem usecase.
func NewUpdateKnowledgeItem(
	categoryService services.CategoryService,
	knowledgeItemService services.KnowledgeItemService,
) *UpdateKnowledgeItem {
	return &UpdateKnowledgeItem{
		categoryService:      categoryService,
		knowledgeItemService: knowledgeItemService,
	}
}

// Handle function performs usecase actions.
func (uc *UpdateKnowledgeItem) Handle(request *models.UpdateKnowledgeItemRequest) (*domain.KnowledgeItem, error) {
	var categories []*domain.Category
	for _, categoryName := range request.Categories {
		cat, err := uc.categoryService.CreateOrGetCategory(categoryName)
		if err != nil {
			return nil, err
		}

		categories = append(categories, cat)
	}

	item, err := uc.knowledgeItemService.UpdateItem(
		request.ID, request.Title, request.Anchor,
		request.Data, request.Tags, categories)
	if err != nil {
		return nil, err
	}

	return item, nil
}
