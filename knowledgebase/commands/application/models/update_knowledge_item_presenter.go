// Package models contains representations of requests and events.
package models

import "github.com/96solutions/neurography/knowledgebase/commands/domain/models"

//go:generate mockgen -package=mock -destination=../../mock/mock_update_knowledge_item_presenter.go -source=update_knowledge_item_presenter.go UpdateKnowledgeItemPresenter

// UpdateKnowledgeItemPresenter represents output presenter of the update models.KnowledgeItem usecase.
type UpdateKnowledgeItemPresenter interface {
	SetResult(item *models.KnowledgeItem)
}
