// Package usecases contains a set of sequences for interactions between services and users.
package usecases

import (
	"context"

	"github.com/96solutions/neurography/knowledgebase/commands/application/models"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/services"
)

// SetMarkToKnowledgeItem type represents usecase that has sequence of actions
// to set/update models.KnowledgeItem`s LastMark.
type SetMarkToKnowledgeItem struct {
	knowledgeItemService services.KnowledgeItemService
	presenter            models.SetMarkToKnowledgeItemPresenter
}

// NewSetMarkToKnowledgeItem function builds new instance of SetMarkToKnowledgeItem usecase.
func NewSetMarkToKnowledgeItem(
	service services.KnowledgeItemService,
	presenter models.SetMarkToKnowledgeItemPresenter,
) *SetMarkToKnowledgeItem {
	return &SetMarkToKnowledgeItem{
		knowledgeItemService: service,
		presenter:            presenter,
	}
}

// Handle function performs usecase actions.
func (uc *SetMarkToKnowledgeItem) Handle(_ context.Context, cmd models.SetMarkToKnowledgeItemCommand) error {
	item, err := uc.knowledgeItemService.SetLatestMark(cmd.ID, cmd.Mark)
	if err != nil {
		return err
	}

	uc.presenter.SetResult(item)
	return nil
}
