// Package usecases contains a set of sequences for interactions between services and users.
package usecases

import (
	"github.com/96solutions/neurography/knowledgebase/commands/application/models"
	domain "github.com/96solutions/neurography/knowledgebase/commands/domain/models"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/services"
)

// SetMarkToKnowledgeItem type represents usecase that has sequence of actions
// to set/update models.KnowledgeItem`s LastMark.
type SetMarkToKnowledgeItem struct {
	knowledgeItemService services.KnowledgeItemService
}

// NewSetMarkToKnowledgeItem function builds new instance of SetMarkToKnowledgeItem usecase.
func NewSetMarkToKnowledgeItem(service services.KnowledgeItemService) *SetMarkToKnowledgeItem {
	return &SetMarkToKnowledgeItem{
		knowledgeItemService: service,
	}
}

// Handle function performs usecase actions.
func (uc *SetMarkToKnowledgeItem) Handle(request models.SetMarkToKnowledgeItemRequest) (*domain.KnowledgeItem, error) {
	return uc.knowledgeItemService.SetLatestMark(request.ID, request.Mark)
}
