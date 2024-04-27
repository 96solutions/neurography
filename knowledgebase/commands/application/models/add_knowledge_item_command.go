// Package models contains representations of requests and events.
package models

// AddKnowledgeItemCommand represents input of the add new models.KnowledgeItem usecase.
type AddKnowledgeItemCommand struct {
	Title      string   `json:"title"`
	Anchor     string   `json:"anchor"`
	Data       string   `json:"data"`
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
}
