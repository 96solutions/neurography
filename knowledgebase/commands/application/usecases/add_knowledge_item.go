// Package usecases contains a set of possible sequences of interactions between services and users.
package usecases

import (
	"github.com/96solutions/neurography/knowledgebase/commands/application/models"
	domain "github.com/96solutions/neurography/knowledgebase/commands/domain/models"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/services"
)

// AddKnowledgeItem type represents usecase that has sequence of actions to create new models.KnowledgeItem.
type AddKnowledgeItem struct {
	categoryService      services.CategoryService
	knowledgeItemService services.KnowledgeItemService
}

// NewAddKnowledgeItem function builds new instance of AddKnowledgeItem usecase.
func NewAddKnowledgeItem(
	categoryService services.CategoryService,
	knowledgeItemService services.KnowledgeItemService,
) *AddKnowledgeItem {
	return &AddKnowledgeItem{
		categoryService:      categoryService,
		knowledgeItemService: knowledgeItemService,
	}
}

// Do function performs usecase actions.
func (uc *AddKnowledgeItem) Do(request *models.AddKnowledgeItemRequest) (*domain.KnowledgeItem, error) {
	var categories []*domain.Category
	for _, categoryName := range request.Categories {
		cat, err := uc.categoryService.CreateOrGetCategory(categoryName)
		if err != nil {
			return nil, err
		}

		categories = append(categories, cat)
	}

	item, err := uc.knowledgeItemService.NewItem(request.Title, request.Anchor, request.Data, request.Tags, categories)
	if err != nil {
		return nil, err
	}

	return item, nil
}
