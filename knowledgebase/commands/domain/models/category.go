// Package models contains types that represent entities of business logic.
package models

// Category type represents category data
// which is used to structure knowledge items.
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
