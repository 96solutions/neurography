// Package services contains domain business rules.
package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/96solutions/neurography/knowledgebase/commands/domain/contracts"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/models"
)

const minTitleLength = 3
const minAnchorLength = 3
const minDataLength = 15
const minTagLength = 1
const minScore = 0
const maxScore = 100
const minMark = 0
const maxMark = 10

// KnowledgeItemService is a scope of business rules & actions related to the Knowledge Item.
type KnowledgeItemService struct {
	repo contracts.KnowledgeItemsRepo
}

// NewKnowledgeItemService function makes new instance of KnowledgeItemService.
func NewKnowledgeItemService(repo contracts.KnowledgeItemsRepo) *KnowledgeItemService {
	return &KnowledgeItemService{
		repo: repo,
	}
}

// NewItem function builds new models.KnowledgeItem instance.
func (s *KnowledgeItemService) NewItem(
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

	item.ID, err = s.repo.Create(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateItem function updates existing models.KnowledgeItem instance.
func (s *KnowledgeItemService) UpdateItem(
	item *models.KnowledgeItem,
	title, anchor, data string,
	tags []string,
	categories []*models.Category,
) (*models.KnowledgeItem, error) {
	err := s.validate(title, anchor, data, tags, categories)
	if err != nil {
		return nil, err //TODO:
	}

	if item.ID == 0 {
		return nil, errors.New("provided Knowledge Item doesn't exist")
	}

	item.Title = title
	item.Anchor = anchor
	item.Data = data
	item.Tags = tags
	item.Categories = categories

	updatedAt := time.Now()
	item.UpdatedAt = &updatedAt

	err = s.repo.Update(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteItem function deletes existing models.KnowledgeItem.
func (s *KnowledgeItemService) DeleteItem(item *models.KnowledgeItem) error {
	if item.ID == 0 {
		return errors.New("provided Knowledge Item doesn't exist")
	}

	return s.repo.Delete(item)
}

func (s *KnowledgeItemService) validate(
	title, anchor, data string,
	tags []string,
	categories []*models.Category,
) error {
	if len(title) <= minTitleLength {
		return errors.New("title is too short")
	}

	if len(anchor) <= minAnchorLength {
		return errors.New("anchor is too short")
	}

	if len(data) <= minDataLength {
		return errors.New("data is too short")
	}

	for _, tag := range tags {
		if len(tag) <= minTagLength {
			return errors.New("tag is too short")
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

func (s *KnowledgeItemService) validateMark(mark int) error {
	if mark < minMark {
		return fmt.Errorf("mark cannot be less than %d", minMark)
	}
	if mark > maxMark {
		return fmt.Errorf("mark cannot be more than %d", maxMark)
	}

	return nil
}

// SetLatestMark sets last testing result to the knowledge item and updates score.
func (s *KnowledgeItemService) SetLatestMark(item *models.KnowledgeItem, mark int) error {
	if item.ID == 0 {
		return errors.New("provided Knowledge Item doesn't exist")
	}

	if err := s.validateMark(mark); err != nil {
		return err
	}

	lastCheckAt := time.Now()
	item.LastCheckAt = &lastCheckAt

	// wipe Score in case of worst mark.
	// means knowledge item has been completely forgotten.
	if mark == minMark {
		item.Score = minMark
	}

	// reduce score if current testing result is worse than previous.
	if item.LastMark > mark {
		item.Score += mark - item.LastMark
	}

	// add mark to the score if current testing result better than previous.
	if item.LastMark <= mark {
		item.Score += mark
	}

	if item.Score < minScore {
		item.Score = minScore
	}
	if item.Score > maxScore {
		item.Score = maxScore
	}

	item.LastMark = mark

	return s.repo.Update(item)
}
