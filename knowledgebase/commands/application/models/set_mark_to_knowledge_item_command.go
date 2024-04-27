// Package models contains representations of requests and events.
package models

// SetMarkToKnowledgeItemCommand represents input of the set new mark to models.KnowledgeItem usecase.
type SetMarkToKnowledgeItemCommand struct {
	ID   int64 `json:"id"`
	Mark int64 `json:"mark"`
}
