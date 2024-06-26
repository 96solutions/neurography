package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/96solutions/neurography/knowledgebase/commands/application/models"
	"github.com/96solutions/neurography/knowledgebase/commands/application/usecases"
	"github.com/96solutions/neurography/knowledgebase/commands/mock"
	"go.uber.org/mock/gomock"
)

func TestDeleteKnowledgeItem_Do(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedItemID := int64(5)

	service := mock.NewMockKnowledgeItemService(ctrl)
	service.EXPECT().DeleteItem(expectedItemID).Return(nil)

	req := &models.DeleteKnowledgeItemCommand{ID: expectedItemID}

	presenter := mock.NewMockDeleteKnowledgeItemPresenter(ctrl)
	presenter.EXPECT().SetResult(true)

	uc := usecases.NewDeleteKnowledgeItem(service, presenter)

	ctx := context.Background()

	err := uc.Handle(ctx, req)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
}

func TestDeleteKnowledgeItem_Do_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedItemID := int64(5)
	expectedError := errors.New("expected error")

	service := mock.NewMockKnowledgeItemService(ctrl)
	service.EXPECT().DeleteItem(expectedItemID).Return(expectedError)

	req := &models.DeleteKnowledgeItemCommand{ID: expectedItemID}

	presenter := mock.NewMockDeleteKnowledgeItemPresenter(ctrl)

	uc := usecases.NewDeleteKnowledgeItem(service, presenter)

	ctx := context.Background()

	err := uc.Handle(ctx, req)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	if err.Error() != expectedError.Error() {
		t.Fatalf("Expected: %s, got: %s", expectedError.Error(), err.Error())
	}
}
