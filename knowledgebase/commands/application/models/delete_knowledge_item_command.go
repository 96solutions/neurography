// Package models contains representations of requests and events.
package models

// DeleteKnowledgeItemCommand represents input of the delete models.KnowledgeItem usecase.
type DeleteKnowledgeItemCommand struct {
	ID int64 `json:"id"`
}
