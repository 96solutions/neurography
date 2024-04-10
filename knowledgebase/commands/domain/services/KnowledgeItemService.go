// Package services contains domain business rules.
package services

import (
	"errors"
	"time"

	"github.com/96solutions/neurography/knowledgebase/commands/domain/models"
)

// KnowledgeItemService is a scope of business rules & actions related to the Knowledge Item.
type KnowledgeItemService struct {
	//
}

// NewKnowledgeItemService function makes new instance of KnowledgeItemService.
func NewKnowledgeItemService() *KnowledgeItemService {
	return &KnowledgeItemService{}
}

// BuildNewItem function builds new models.KnowledgeItem instance.
func (s *KnowledgeItemService) BuildNewItem(
	title, anchor, data string,
	tags []string,
	categories []*models.Category,
) (*models.KnowledgeItem, error) {
	err := s.validate(title, anchor, data, tags, categories)
	if err != nil {
		return nil, err //TODO:
	}

	createdAt := time.Now()

	item := &models.KnowledgeItem{
		Title:      title,
		Anchor:     anchor,
		Data:       data,
		Categories: categories,
		Tags:       tags,
		CreatedAt:  &createdAt,
	}

	return item, nil
}

// UpdateItem function updates existing models.KnowledgeItem instance.
func (s *KnowledgeItemService) UpdateItem(
	item *models.KnowledgeItem,
	title, anchor, data string,
	tags []string,
	categories []*models.Category,
) error {
	err := s.validate(title, anchor, data, tags, categories)
	if err != nil {
		return err //TODO:
	}

	if item.ID == 0 {
		return errors.New("provided Knowledge Item doesn't exist")
	}

	item.Title = title
	item.Anchor = anchor
	item.Data = data
	item.Tags = tags
	item.Categories = categories

	return nil
}

func (s *KnowledgeItemService) validate(
	title, anchor, data string,
	tags []string,
	categories []*models.Category,
) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}

	if anchor == "" {
		return errors.New("anchor cannot be empty")
	}

	if data == "" {
		return errors.New("data cannot be empty")
	}

	for _, tag := range tags {
		if tag == "" {
			return errors.New("tag cannot be empty")
		}
	}

	for _, category := range categories {
		if category == nil {
			return errors.New("category cannot be empty")
		}
		if category.ID == 0 {
			return errors.New("category doesn't exist")
		}
	}

	//TODO: improve validation

	return nil
}
