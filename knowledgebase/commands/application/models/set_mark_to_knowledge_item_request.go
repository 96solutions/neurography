// Package models contains representations of requests and events.
package models

// SetMarkToKnowledgeItemRequest represents request to set new mark to models.KnowledgeItem.
type SetMarkToKnowledgeItemRequest struct {
	ID   int64 `json:"id"`
	Mark int64 `json:"mark"`
}
