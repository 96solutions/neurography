package models

import (
	"github.com/96solutions/neurography/knowledgebase/commands/domain/models"
)

//go:generate mockgen -package=mock -destination=../../mock/mock_add_knowledge_item_presenter.go -source=add_knowledge_item_presenter.go AddKnowledgeItemPresenter

// AddKnowledgeItemPresenter represents output presenter of the add models.KnowledgeItem usecase.
type AddKnowledgeItemPresenter interface {
	SetResult(item *models.KnowledgeItem)
}
