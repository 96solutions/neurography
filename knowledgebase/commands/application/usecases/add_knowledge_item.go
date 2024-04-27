// Package usecases contains a set of sequences for interactions between services and users.
package usecases

import (
	"context"

	"github.com/96solutions/neurography/knowledgebase/commands/application/models"
	domain "github.com/96solutions/neurography/knowledgebase/commands/domain/models"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/services"
)

// AddKnowledgeItem type represents usecase that has sequence of actions to create new models.KnowledgeItem.
type AddKnowledgeItem struct {
	categoryService      services.CategoryService
	knowledgeItemService services.KnowledgeItemService
	presenter            models.AddKnowledgeItemPresenter
}

// NewAddKnowledgeItem function builds new instance of AddKnowledgeItem usecase.
func NewAddKnowledgeItem(
	categoryService services.CategoryService,
	knowledgeItemService services.KnowledgeItemService,
	presenter models.AddKnowledgeItemPresenter,
) *AddKnowledgeItem {
	return &AddKnowledgeItem{
		categoryService:      categoryService,
		knowledgeItemService: knowledgeItemService,
		presenter:            presenter,
	}
}

// Handle function performs usecase actions.
func (uc *AddKnowledgeItem) Handle(_ context.Context, cmd *models.AddKnowledgeItemCommand) error {
	var categories []*domain.Category
	for _, categoryName := range cmd.Categories {
		cat, err := uc.categoryService.CreateOrGetCategory(categoryName)
		if err != nil {
			return err
		}

		categories = append(categories, cat)
	}

	item, err := uc.knowledgeItemService.NewItem(cmd.Title, cmd.Anchor, cmd.Data, cmd.Tags, categories)
	if err != nil {
		return err
	}

	uc.presenter.SetResult(item)

	return nil
}
