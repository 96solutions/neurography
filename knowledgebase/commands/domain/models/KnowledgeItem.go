// Package models contains types that represent entities of business logic.
package models

import "time"

// KnowledgeItem represents one particular piece of knowledge.
type KnowledgeItem struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Anchor string `json:"anchor"`
	Data   string `json:"description"`

	Categories []*Category `json:"categories"`

	Tags []string `json:"tags,omitempty"`

	Score     int `json:"score"`
	LastScore int `json:"last_score"`

	LastCheckAt *time.Time `json:"last_check_at"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
