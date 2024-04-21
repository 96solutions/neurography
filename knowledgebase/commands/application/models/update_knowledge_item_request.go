// Package models contains representations of requests and events.
package models

// UpdateKnowledgeItemRequest represents request to update models.KnowledgeItem.
type UpdateKnowledgeItemRequest struct {
	ID         int64    `json:"id"`
	Title      string   `json:"title"`
	Anchor     string   `json:"anchor"`
	Data       string   `json:"data"`
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
}
