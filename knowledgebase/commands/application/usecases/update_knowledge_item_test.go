package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/96solutions/neurography/knowledgebase/commands/application/models"
	"github.com/96solutions/neurography/knowledgebase/commands/application/usecases"
	domain "github.com/96solutions/neurography/knowledgebase/commands/domain/models"
	"github.com/96solutions/neurography/knowledgebase/commands/mock"
	"go.uber.org/mock/gomock"
)

func TestUpdateKnowledgeItem_Do_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedItemID := int64(5)
	cmd := &models.UpdateKnowledgeItemCommand{
		ID:         expectedItemID,
		Title:      "expectedTitle",
		Anchor:     "expectedAnchor",
		Data:       "expectedData and more",
		Tags:       []string{"tag1", "tag2"},
		Categories: []string{"category1", "category2"},
	}

	expectedCategories := []*domain.Category{
		&domain.Category{
			ID:   1,
			Name: cmd.Categories[0],
		},
		&domain.Category{
			ID:   1,
			Name: cmd.Categories[1],
		},
	}

	expectedItem := &domain.KnowledgeItem{
		ID:         expectedItemID,
		Title:      cmd.Title,
		Anchor:     cmd.Anchor,
		Data:       cmd.Data,
		Categories: expectedCategories,
		Tags:       cmd.Tags,
	}

	catService := mock.NewMockCategoryService(ctrl)
	catService.EXPECT().CreateOrGetCategory(cmd.Categories[0]).DoAndReturn(func(_ string) (*domain.Category, error) {
		return expectedCategories[0], nil
	})
	catService.EXPECT().CreateOrGetCategory(cmd.Categories[1]).DoAndReturn(func(_ string) (*domain.Category, error) {
		return expectedCategories[1], nil
	})

	itemService := mock.NewMockKnowledgeItemService(ctrl)
	itemService.EXPECT().UpdateItem(
		expectedItemID, cmd.Title, cmd.Anchor,
		cmd.Data, cmd.Tags, expectedCategories).
		Return(expectedItem, nil)

	presenter := mock.NewMockUpdateKnowledgeItemPresenter(ctrl)
	presenter.EXPECT().SetResult(expectedItem).Do(func(item *domain.KnowledgeItem) {
		if item.ID != expectedItem.ID {
			t.Errorf("expected ID %d, got %d", expectedItem.ID, item.ID)
		}
		if item.Title != expectedItem.Title {
			t.Errorf("expected Title %s, got %s", expectedItem.Title, item.Title)
		}
		if item.Anchor != expectedItem.Anchor {
			t.Errorf("expected Anchor %s, got %s", expectedItem.Anchor, item.Anchor)
		}
		if item.Data != expectedItem.Data {
			t.Errorf("expected Data %s, got %s", expectedItem.Data, item.Data)
		}
		if len(item.Categories) != len(expectedItem.Categories) {
			t.Errorf("expected %d categories, got %d", len(expectedItem.Categories), len(item.Categories))
		} else {
			for i := 0; i < len(expectedItem.Categories); i++ {
				if item.Categories[i].Name != expectedItem.Categories[i].Name {
					t.Errorf("expected Categories %d, got %d", expectedItem.Categories[i].ID, item.Categories[i].ID)
				}
			}
		}

		if len(item.Tags) != len(expectedItem.Tags) {
			t.Errorf("expected %d tags, got %d", len(expectedItem.Tags), len(item.Tags))
		} else {
			for i := 0; i < len(expectedItem.Tags); i++ {
				if item.Tags[i] != expectedItem.Tags[i] {
					t.Errorf("expected Tags %s, got %s", expectedItem.Tags[i], item.Tags[i])
				}
			}
		}
	})

	ctx := context.Background()

	uc := usecases.NewUpdateKnowledgeItem(catService, itemService, presenter)

	err := uc.Handle(ctx, cmd)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateKnowledgeItem_Do_CategoryServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedItemID := int64(5)
	cmd := &models.UpdateKnowledgeItemCommand{
		ID:         expectedItemID,
		Title:      "expectedTitle",
		Anchor:     "expectedAnchor",
		Data:       "expectedData and more",
		Tags:       []string{"tag1", "tag2"},
		Categories: []string{"category1", "category2"},
	}
	expectedError := errors.New("expected error")

	catService := mock.NewMockCategoryService(ctrl)
	catService.EXPECT().CreateOrGetCategory(cmd.Categories[0]).DoAndReturn(func(_ string) (*domain.Category, error) {
		return nil, expectedError
	})

	presenter := mock.NewMockUpdateKnowledgeItemPresenter(ctrl)

	itemService := mock.NewMockKnowledgeItemService(ctrl)

	uc := usecases.NewUpdateKnowledgeItem(catService, itemService, presenter)

	ctx := context.Background()

	err := uc.Handle(ctx, cmd)
	if !errors.Is(err, expectedError) {
		t.Errorf("expected error %s, got %s", expectedError, err)
	}
}

func TestUpdateKnowledgeItem_Do_KnowledgeItemServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedItemID := int64(5)
	cmd := &models.UpdateKnowledgeItemCommand{
		ID:         expectedItemID,
		Title:      "expectedTitle",
		Anchor:     "expectedAnchor",
		Data:       "expectedData and more",
		Tags:       []string{"tag1", "tag2"},
		Categories: []string{"category1"},
	}
	expectedCategory := &domain.Category{
		ID:   1,
		Name: cmd.Categories[0],
	}
	expectedError := errors.New("expected error")

	catService := mock.NewMockCategoryService(ctrl)
	catService.EXPECT().CreateOrGetCategory(cmd.Categories[0]).DoAndReturn(func(_ string) (*domain.Category, error) {
		return expectedCategory, nil
	})

	itemService := mock.NewMockKnowledgeItemService(ctrl)
	itemService.EXPECT().
		UpdateItem(expectedItemID, cmd.Title, cmd.Anchor, cmd.Data, cmd.Tags, []*domain.Category{expectedCategory}).
		Return(nil, expectedError)

	presenter := mock.NewMockUpdateKnowledgeItemPresenter(ctrl)

	uc := usecases.NewUpdateKnowledgeItem(catService, itemService, presenter)

	ctx := context.Background()

	err := uc.Handle(ctx, cmd)
	if !errors.Is(err, expectedError) {
		t.Errorf("expected error %s, got %s", expectedError, err)
	}
}
