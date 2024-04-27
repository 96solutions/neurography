// Package usecases contains a set of sequences for interactions between services and users.
package usecases

import (
	"context"

	"github.com/96solutions/neurography/knowledgebase/commands/application/models"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/services"
)

// DeleteKnowledgeItem type represents usecase that has sequence of actions to delete new models.KnowledgeItem.
type DeleteKnowledgeItem struct {
	knowledgeItemService services.KnowledgeItemService
	presenter            models.DeleteKnowledgeItemPresenter
}

// NewDeleteKnowledgeItem function builds new instance of DeleteKnowledgeItem usecase.
func NewDeleteKnowledgeItem(
	knowledgeItemService services.KnowledgeItemService,
	presenter models.DeleteKnowledgeItemPresenter,
) *DeleteKnowledgeItem {
	return &DeleteKnowledgeItem{
		knowledgeItemService: knowledgeItemService,
		presenter:            presenter,
	}
}

// Handle function performs usecase actions.
func (uc *DeleteKnowledgeItem) Handle(_ context.Context, cmd *models.DeleteKnowledgeItemCommand) error {
	err := uc.knowledgeItemService.DeleteItem(cmd.ID)
	if err != nil {
		return err
	}

	uc.presenter.SetResult(true)
	return nil
}
