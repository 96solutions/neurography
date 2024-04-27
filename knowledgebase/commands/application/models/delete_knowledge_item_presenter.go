// Package models contains representations of requests and events.
package models

//go:generate mockgen -package=mock -destination=../../mock/mock_delete_knowledge_item_presenter.go -source=delete_knowledge_item_presenter.go DeleteKnowledgeItemPresenter

// DeleteKnowledgeItemPresenter represents output of the delete models.KnowledgeItem usecase.
type DeleteKnowledgeItemPresenter interface {
	SetResult(bool)
}
