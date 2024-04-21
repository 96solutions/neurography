package usecases_test

import (
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

	req := models.SetMarkToKnowledgeItemRequest{
		ID:   expectedItemID,
		Mark: expectedMark,
	}

	item := &domain.KnowledgeItem{
		ID:       expectedItemID,
		LastMark: expectedMark,
	}

	knowledgeItemsService := mock.NewMockKnowledgeItemService(ctrl)
	knowledgeItemsService.EXPECT().SetLatestMark(expectedItemID, expectedMark).Return(item, nil)

	uc := usecases.NewSetMarkToKnowledgeItem(knowledgeItemsService)

	resultItem, err := uc.Handle(req)
	if err != nil {
		t.Error(err)
	}
	if resultItem.ID != expectedItemID {
		t.Errorf("got item ID %d, want %d", resultItem.ID, expectedItemID)
	}
	if resultItem.LastMark != expectedMark {
		t.Errorf("got item LastMark %d, want %d", resultItem.LastMark, expectedMark)
	}
}
