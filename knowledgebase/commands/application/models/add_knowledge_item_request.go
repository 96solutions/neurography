// Package models contains base components of the application level.
package models

// AddKnowledgeItemRequest represents request to create new models.KnowledgeItem.
type AddKnowledgeItemRequest struct {
	Title, Anchor, Data string
	Tags, Categories    []string
}
