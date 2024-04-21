package services_test

import (
	"errors"
	"testing"

	"github.com/96solutions/neurography/knowledgebase/commands/domain/models"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/services"
	"github.com/96solutions/neurography/knowledgebase/commands/mock"
	"go.uber.org/mock/gomock"
)

func TestCategoryService_CreateOrGetCategory_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCategoriesRepo(ctrl)
	s := services.NewCategoryService(repo)

	expectedCategoryName := "expectedCategoryName"
	expectedCategory := &models.Category{
		ID:   15,
		Name: expectedCategoryName,
	}
	repo.EXPECT().FindByName(expectedCategoryName).Return(expectedCategory, nil)

	cat, err := s.CreateOrGetCategory(expectedCategoryName)
	if err != nil {
		t.Fatal(err)
	}
	if cat.Name != expectedCategoryName {
		t.Errorf("Category name: expected %s, got %s", expectedCategoryName, cat.Name)
	}
	if cat.ID != expectedCategory.ID {
		t.Errorf("Category ID: expected %d, got %d", expectedCategory.ID, cat.ID)
	}
}

func TestCategoryService_CreateOrGetCategory_Existing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCategoriesRepo(ctrl)
	s := services.NewCategoryService(repo)

	expectedCategoryName := "expectedCategoryName"
	var expectedCategoryID int64 = 15

	repo.EXPECT().FindByName(expectedCategoryName).Return(nil, nil)
	repo.EXPECT().Create(gomock.Any()).DoAndReturn(func(category *models.Category) (int64, error) {
		if category.Name != expectedCategoryName {
			t.Errorf("Category name: expected %s, got %s", expectedCategoryName, category.Name)
		}

		category.ID = expectedCategoryID

		return expectedCategoryID, nil
	})

	cat, err := s.CreateOrGetCategory(expectedCategoryName)
	if err != nil {
		t.Fatal(err)
	}
	if cat.Name != expectedCategoryName {
		t.Errorf("Category name: expected %s, got %s", expectedCategoryName, cat.Name)
	}
	if cat.ID != expectedCategoryID {
		t.Errorf("Category ID: expected %d, got %d", expectedCategoryID, cat.ID)
	}
}

func TestCategoryService_CreateOrGetCategory_TooShortName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCategoriesRepo(ctrl)
	s := services.NewCategoryService(repo)

	expectedErrorName := "category name is too short"
	expectedCategoryName := "s"

	repo.EXPECT().FindByName(expectedCategoryName).Return(nil, nil)
	_, err := s.CreateOrGetCategory(expectedCategoryName)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != expectedErrorName {
		t.Errorf("Category name: expected %s, got %s", expectedErrorName, err.Error())
	}
}

func TestCategoryService_CreateOrGetCategory_RepoFindByNameError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCategoriesRepo(ctrl)
	s := services.NewCategoryService(repo)

	expectedError := errors.New("expected error")
	expectedCategoryName := "s"

	repo.EXPECT().FindByName(expectedCategoryName).Return(nil, expectedError)
	_, err := s.CreateOrGetCategory(expectedCategoryName)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, expectedError) {
		t.Errorf("Category name: expected %s, got %s", expectedError.Error(), err.Error())
	}
}

func TestCategoryService_CreateOrGetCategory_RepoCreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCategoriesRepo(ctrl)
	s := services.NewCategoryService(repo)

	expectedError := errors.New("expected error")
	expectedCategoryName := "expectedCategoryName"
	var expectedCategoryID int64 = 15

	repo.EXPECT().FindByName(expectedCategoryName).Return(nil, nil)
	repo.EXPECT().Create(gomock.Any()).DoAndReturn(func(category *models.Category) (int64, error) {
		if category.Name != expectedCategoryName {
			t.Errorf("Category name: expected %s, got %s", expectedCategoryName, category.Name)
		}

		category.ID = expectedCategoryID

		return expectedCategoryID, expectedError
	})

	_, err := s.CreateOrGetCategory(expectedCategoryName)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, expectedError) {
		t.Errorf("Category name: expected %s, got %s", expectedError.Error(), err.Error())
	}
}

func TestCategoryService_DeleteCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCategoriesRepo(ctrl)
	s := services.NewCategoryService(repo)

	expectedCategoryName := "expectedCategoryName"
	expectedCategory := &models.Category{
		ID:   15,
		Name: expectedCategoryName,
	}

	repo.EXPECT().FindByName(expectedCategoryName).Return(expectedCategory, nil)
	repo.EXPECT().Delete(expectedCategory).Return(nil)

	err := s.DeleteCategory(expectedCategoryName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCategoryService_DeleteCategory_NotExisting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCategoriesRepo(ctrl)
	s := services.NewCategoryService(repo)

	expectedCategoryName := "expectedCategoryName"
	expectedErrorMessage := "category not exists"

	repo.EXPECT().FindByName(expectedCategoryName).Return(nil, nil)

	err := s.DeleteCategory(expectedCategoryName)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != expectedErrorMessage {
		t.Errorf("Category name: expected %s, got %s", expectedErrorMessage, err.Error())
	}
}

func TestCategoryService_DeleteCategory_RepoFindByNameError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCategoriesRepo(ctrl)
	s := services.NewCategoryService(repo)

	expectedCategoryName := "expectedCategoryName"
	expectedError := errors.New("expected error")

	repo.EXPECT().FindByName(expectedCategoryName).Return(nil, expectedError)

	err := s.DeleteCategory(expectedCategoryName)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, expectedError) {
		t.Errorf("Category name: expected %s, got %s", expectedError.Error(), err.Error())
	}
}

func TestCategoryService_DeleteCategory_RepoDeleteError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCategoriesRepo(ctrl)
	s := services.NewCategoryService(repo)

	expectedCategoryName := "expectedCategoryName"
	expectedError := errors.New("expected error")
	expectedCategory := &models.Category{
		ID:   15,
		Name: expectedCategoryName,
	}

	repo.EXPECT().FindByName(expectedCategoryName).Return(expectedCategory, nil)
	repo.EXPECT().Delete(expectedCategory).Return(expectedError)

	err := s.DeleteCategory(expectedCategoryName)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, expectedError) {
		t.Errorf("Category name: expected %s, got %s", expectedError.Error(), err.Error())
	}
}
