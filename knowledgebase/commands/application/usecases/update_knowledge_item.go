// Package usecases contains a set of sequences for interactions between services and users.
package usecases

import (
	"context"

	"github.com/96solutions/neurography/knowledgebase/commands/application/models"
	domain "github.com/96solutions/neurography/knowledgebase/commands/domain/models"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/services"
)

// UpdateKnowledgeItem type represents usecase that has sequence of actions to update new models.KnowledgeItem.
type UpdateKnowledgeItem struct {
	categoryService      services.CategoryService
	knowledgeItemService services.KnowledgeItemService
	presenter            models.UpdateKnowledgeItemPresenter
}

// NewUpdateKnowledgeItem function builds new instance of UpdateKnowledgeItem usecase.
func NewUpdateKnowledgeItem(
	categoryService services.CategoryService,
	knowledgeItemService services.KnowledgeItemService,
	presenter models.UpdateKnowledgeItemPresenter,
) *UpdateKnowledgeItem {
	return &UpdateKnowledgeItem{
		categoryService:      categoryService,
		knowledgeItemService: knowledgeItemService,
		presenter:            presenter,
	}
}

// Handle function performs usecase actions.
func (uc *UpdateKnowledgeItem) Handle(_ context.Context, cmd *models.UpdateKnowledgeItemCommand) error {
	var categories []*domain.Category
	for _, categoryName := range cmd.Categories {
		cat, err := uc.categoryService.CreateOrGetCategory(categoryName)
		if err != nil {
			return err
		}

		categories = append(categories, cat)
	}

	item, err := uc.knowledgeItemService.UpdateItem(
		cmd.ID, cmd.Title, cmd.Anchor,
		cmd.Data, cmd.Tags, categories)
	if err != nil {
		return err
	}

	uc.presenter.SetResult(item)
	return nil
}
