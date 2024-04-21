// Package models contains representations of requests and events.
package models

// AddKnowledgeItemRequest represents request to create new models.KnowledgeItem.
type AddKnowledgeItemRequest struct {
	Title      string   `json:"title"`
	Anchor     string   `json:"anchor"`
	Data       string   `json:"data"`
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
}
