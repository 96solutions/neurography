// Package usecases contains a set of sequences for interactions between services and users.
package usecases

import (
	"github.com/96solutions/neurography/knowledgebase/commands/application/models"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/services"
)

// DeleteKnowledgeItem type represents usecase that has sequence of actions to delete new models.KnowledgeItem.
type DeleteKnowledgeItem struct {
	knowledgeItemService services.KnowledgeItemService
}

// NewDeleteKnowledgeItem function builds new instance of DeleteKnowledgeItem usecase.
func NewDeleteKnowledgeItem(
	knowledgeItemService services.KnowledgeItemService,
) *DeleteKnowledgeItem {
	return &DeleteKnowledgeItem{
		knowledgeItemService: knowledgeItemService,
	}
}

// Handle function performs usecase actions.
func (uc *DeleteKnowledgeItem) Handle(request *models.DeleteKnowledgeItemRequest) error {
	return uc.knowledgeItemService.DeleteItem(request.ID)
}
