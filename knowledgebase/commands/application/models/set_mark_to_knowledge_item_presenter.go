// Package models contains representations of requests and events.
package models

import "github.com/96solutions/neurography/knowledgebase/commands/domain/models"

//go:generate mockgen -package=mock -destination=../../mock/mock_set_mark_to_knowledge_item_presenter.go -source=set_mark_to_knowledge_item_presenter.go SetMarkToKnowledgeItemPresenter

// SetMarkToKnowledgeItemPresenter represents output presenter of the set new mark to models.KnowledgeItem usecase.
type SetMarkToKnowledgeItemPresenter interface {
	SetResult(item *models.KnowledgeItem)
}
