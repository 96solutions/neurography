// Package models contains representations of requests and events.
package models

// DeleteKnowledgeItemRequest represents request to delete models.KnowledgeItem.
type DeleteKnowledgeItemRequest struct {
	ID int64 `json:"id"`
}
