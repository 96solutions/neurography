// Package models contains representations of requests and events.
package models

// UpdateKnowledgeItemCommand represents input of the update models.KnowledgeItem usecase.
type UpdateKnowledgeItemCommand struct {
	ID         int64    `json:"id"`
	Title      string   `json:"title"`
	Anchor     string   `json:"anchor"`
	Data       string   `json:"data"`
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
}
