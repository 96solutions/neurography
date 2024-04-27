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

func TestSetMarkToKnowledgeItem_Do(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedItemID := int64(5)
	expectedMark := int64(8)

	req := models.SetMarkToKnowledgeItemCommand{
		ID:   expectedItemID,
		Mark: expectedMark,
	}

	item := &domain.KnowledgeItem{
		ID:       expectedItemID,
		LastMark: expectedMark,
	}

	knowledgeItemsService := mock.NewMockKnowledgeItemService(ctrl)
	knowledgeItemsService.EXPECT().SetLatestMark(expectedItemID, expectedMark).Return(item, nil)

	presenter := mock.NewMockSetMarkToKnowledgeItemPresenter(ctrl)
	presenter.EXPECT().SetResult(gomock.Any()).Do(func(resultItem *domain.KnowledgeItem) {
		if resultItem.ID != expectedItemID {
			t.Errorf("got item ID %d, want %d", resultItem.ID, expectedItemID)
		}
		if resultItem.LastMark != expectedMark {
			t.Errorf("got item LastMark %d, want %d", resultItem.LastMark, expectedMark)
		}
	})

	uc := usecases.NewSetMarkToKnowledgeItem(knowledgeItemsService, presenter)

	ctx := context.Background()

	err := uc.Handle(ctx, req)
	if err != nil {
		t.Error(err)
	}
}

func TestSetMarkToKnowledgeItem_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedItemID := int64(5)
	expectedMark := int64(8)

	req := models.SetMarkToKnowledgeItemCommand{
		ID:   expectedItemID,
		Mark: expectedMark,
	}

	expectedError := errors.New("expected error")

	knowledgeItemsService := mock.NewMockKnowledgeItemService(ctrl)
	knowledgeItemsService.EXPECT().SetLatestMark(expectedItemID, expectedMark).Return(nil, expectedError)

	presenter := mock.NewMockSetMarkToKnowledgeItemPresenter(ctrl)

	uc := usecases.NewSetMarkToKnowledgeItem(knowledgeItemsService, presenter)

	ctx := context.Background()

	err := uc.Handle(ctx, req)
	if err == nil {
		t.Error("expected error")
	}
	if !errors.Is(err, expectedError) {
		t.Errorf("got error %v, want %v", err, expectedError)
	}
}
