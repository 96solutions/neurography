package services_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/96solutions/neurography/knowledgebase/commands/domain/models"
	"github.com/96solutions/neurography/knowledgebase/commands/domain/services"
	"github.com/96solutions/neurography/knowledgebase/commands/mock"
	"go.uber.org/mock/gomock"
)

func TestKnowledgeItemService_NewItem_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)

	var expectedItemID int64 = 11
	expectedTitle := "expectedTitle"
	expectedAnchor := "expectedAnchor"
	expectedData := "expectedData and something else"
	expectedTags := []string{"expectedTag1", "expectedTag2", "expectedTag3"}
	expectedCategories := []*models.Category{
		&models.Category{
			ID:   1,
			Name: "expectedCategory1",
		},
		&models.Category{
			ID:   2,
			Name: "expectedCategory2",
		},
		&models.Category{
			ID:   3,
			Name: "expectedCategory3",
		},
	}

	repo.EXPECT().Create(gomock.Any()).DoAndReturn(func(item *models.KnowledgeItem) (int64, error) {
		if item.Title != expectedTitle {
			return 0, errors.New("expected title to be: " + expectedTitle)
		}
		if item.Anchor != expectedAnchor {
			return 0, errors.New("expected anchor to: " + expectedAnchor)
		}
		if item.Data != expectedData {
			return 0, errors.New("expected data to be: " + expectedData)
		}
		if len(item.Tags) != len(expectedTags) {
			return 0, errors.New("expected tags to be equal")
		}
		for i, tag := range item.Tags {
			if tag != expectedTags[i] {
				return 0, errors.New("expected tags to be equal")
			}
		}
		if len(item.Categories) != len(expectedCategories) {
			return 0, errors.New("expected categories to be equal")
		}
		for i, category := range item.Categories {
			if category.ID != expectedCategories[i].ID {
				return 0, errors.New("expected categories IDs to be equal")
			}
			if category.Name != expectedCategories[i].Name {
				return 0, errors.New("expected categories Names to be equal")
			}
		}

		return expectedItemID, nil
	})

	s := services.NewKnowledgeItemService(repo)
	item, err := s.NewItem(expectedTitle, expectedAnchor, expectedData, expectedTags, expectedCategories)
	if err != nil {
		t.Fatal(err)
	}
	if item.ID != expectedItemID {
		t.Errorf("expected item.ID: %+v, got item.ID: %+v", expectedItemID, item.ID)
	}
	if item.CreatedAt == nil {
		t.Error("expected item.CreatedAt")
	}
	if item.UpdatedAt != nil {
		t.Error("expected item.UpdatedAt to be nil")
	}
	if item.Score != 0 {
		t.Error("expected item.Score to be 0")
	}
	if item.LastMark != 0 {
		t.Error("expected item.LastMark to be 0")
	}
	if item.LastCheckAt != nil {
		t.Error("expected item.LastCheckAt to be nil")
	}
}

func TestKnowledgeItemService_NewItem_RepoCreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)

	var expectedItemID int64 = 11
	expectedTitle := "expectedTitle"
	expectedAnchor := "expectedAnchor"
	expectedData := "expectedData and something else"
	expectedTags := []string{"expectedTag1", "expectedTag2", "expectedTag3"}
	expectedCategories := []*models.Category{
		&models.Category{
			ID:   1,
			Name: "expectedCategory1",
		},
		&models.Category{
			ID:   2,
			Name: "expectedCategory2",
		},
		&models.Category{
			ID:   3,
			Name: "expectedCategory3",
		},
	}
	expectedError := errors.New("expected error")

	repo.EXPECT().Create(gomock.Any()).DoAndReturn(func(item *models.KnowledgeItem) (int64, error) {
		if item.Title != expectedTitle {
			return 0, errors.New("expected title to be: " + expectedTitle)
		}
		if item.Anchor != expectedAnchor {
			return 0, errors.New("expected anchor to: " + expectedAnchor)
		}
		if item.Data != expectedData {
			return 0, errors.New("expected data to be: " + expectedData)
		}
		if len(item.Tags) != len(expectedTags) {
			return 0, errors.New("expected tags to be equal")
		}
		for i, tag := range item.Tags {
			if tag != expectedTags[i] {
				return 0, errors.New("expected tags to be equal")
			}
		}
		if len(item.Categories) != len(expectedCategories) {
			return 0, errors.New("expected categories to be equal")
		}
		for i, category := range item.Categories {
			if category.ID != expectedCategories[i].ID {
				return 0, errors.New("expected categories IDs to be equal")
			}
			if category.Name != expectedCategories[i].Name {
				return 0, errors.New("expected categories Names to be equal")
			}
		}

		return expectedItemID, expectedError
	})

	s := services.NewKnowledgeItemService(repo)
	_, err := s.NewItem(expectedTitle, expectedAnchor, expectedData, expectedTags, expectedCategories)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != expectedError.Error() {
		t.Fatalf("expected error: %s, got: %s", expectedError.Error(), err.Error())
	}
}

func TestKnowledgeItemService_NewItem_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)

	testCases := []struct {
		name               string
		expectedItemID     int
		expectedTitle      string
		expectedAnchor     string
		expectedData       string
		expectedTags       []string
		expectedCategories []*models.Category
		expectedError      error
	}{
		{
			name:               "title too short",
			expectedItemID:     15,
			expectedTitle:      "e", // too short
			expectedAnchor:     "expectedAnchor",
			expectedData:       "expectedData and something",
			expectedTags:       []string{"expectedTag1", "expectedTag2", "expectedTag3"},
			expectedCategories: make([]*models.Category, 0),
			expectedError:      errors.New("title is too short"),
		},
		{
			name:               "anchor too short",
			expectedItemID:     15,
			expectedTitle:      "expectedTitle",
			expectedAnchor:     "e", // too short
			expectedData:       "expectedData and something",
			expectedTags:       []string{"expectedTag1", "expectedTag2", "expectedTag3"},
			expectedCategories: make([]*models.Category, 0),
			expectedError:      errors.New("anchor is too short"),
		},
		{
			name:               "data too short",
			expectedItemID:     15,
			expectedTitle:      "expectedTitle",
			expectedAnchor:     "expectedAnchor",
			expectedData:       "expectedData", // too short
			expectedTags:       []string{"expectedTag1", "expectedTag2", "expectedTag3"},
			expectedCategories: make([]*models.Category, 0),
			expectedError:      errors.New("data is too short"),
		},
		{
			name:               "tag too short",
			expectedItemID:     15,
			expectedTitle:      "expectedTitle",
			expectedAnchor:     "expectedAnchor",
			expectedData:       "expectedData and something",
			expectedTags:       []string{"e", "expectedTag2", "expectedTag3"}, // too short
			expectedCategories: make([]*models.Category, 0),
			expectedError:      errors.New("tag is too short"),
		},
		{
			name:               "category cannot be empty",
			expectedItemID:     15,
			expectedTitle:      "expectedTitle",
			expectedAnchor:     "expectedAnchor",
			expectedData:       "expectedData and something",
			expectedTags:       []string{"expectedTag1", "expectedTag2", "expectedTag3"},
			expectedCategories: []*models.Category{nil}, // empty category
			expectedError:      errors.New("category cannot be empty"),
		},
		{
			name:           "category doesn't exist",
			expectedItemID: 15,
			expectedTitle:  "expectedTitle",
			expectedAnchor: "expectedAnchor",
			expectedData:   "expectedData and something",
			expectedTags:   []string{"expectedTag1", "expectedTag2", "expectedTag3"},
			expectedCategories: []*models.Category{&models.Category{
				ID:   0, // not exists
				Name: "Category Name",
			}},
			expectedError: errors.New("category doesn't exist"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := services.NewKnowledgeItemService(repo)
			_, err := s.NewItem(tc.expectedTitle, tc.expectedAnchor, tc.expectedData, tc.expectedTags, tc.expectedCategories)
			if err.Error() != tc.expectedError.Error() {
				t.Fatalf("expected error: %s, got: %s", tc.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestKnowledgeItemService_UpdateItem_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expectedItemID int64 = 5
	item := &models.KnowledgeItem{
		ID:     expectedItemID,
		Title:  "Test Item",
		Anchor: "Test Anchor",
		Data:   "Test Data and Something more",
		Tags:   []string{"tag1", "tag2"},
	}
	expectedTitle := "expected title"
	expectedAnchor := "expected anchor"
	expectedData := "expected data and something more"
	categories := []*models.Category{
		&models.Category{
			ID:   1,
			Name: "Category Name1",
		},
	}
	tags := []string{"tag1", "tag2", "tag3"}

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)
	repo.EXPECT().FindByID(expectedItemID).Return(item, nil)
	repo.EXPECT().Save(item).Return(nil)

	s := services.NewKnowledgeItemService(repo)
	result, err := s.UpdateItem(expectedItemID, expectedTitle, expectedAnchor, expectedData, tags, categories)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != item.ID {
		t.Fatalf("expected ID: %d, got: %d", item.ID, result.ID)
	}
	if result.Title != expectedTitle {
		t.Fatalf("expected Title: %s, got: %s", expectedTitle, result.Title)
	}
	if result.Anchor != expectedAnchor {
		t.Fatalf("expected Anchor: %s, got: %s", expectedAnchor, result.Anchor)
	}
	if result.Data != expectedData {
		t.Fatalf("expected Data: %s, got: %s", expectedData, result.Data)
	}
	if len(result.Tags) == len(tags) {
		for i := range result.Tags {
			if result.Tags[i] != tags[i] {
				t.Fatalf("expected Tag: %s, got: %s", tags[i], result.Tags[i])
			}
		}
	} else {
		t.Fatalf("expected number of Tags: %v, got: %v", len(tags), len(result.Tags))
	}
	if len(result.Categories) == len(categories) {
		for i := range result.Categories {
			if result.Categories[i] != categories[i] {
				t.Fatalf("expected Category: %d, got: %d", categories[i].ID, result.Categories[i].ID)
			}
		}
	} else {
		t.Fatalf("expected number of Categorys: %v, got: %v", len(categories), len(result.Categories))
	}
	if result.UpdatedAt == nil {
		t.Error("expected UpdatedAt to be set")
	}
}

func TestKnowledgeItemService_UpdateItem_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expectedItemID int64 = 5
	item := &models.KnowledgeItem{
		ID:     expectedItemID,
		Title:  "Test Item",
		Anchor: "Test Anchor",
		Data:   "Test Data and Something more",
		Tags:   []string{"tag1", "tag2"},
	}
	expectedTitle := "expected title"
	expectedAnchor := "expected anchor"
	expectedData := "expected data and something more"
	categories := []*models.Category{
		&models.Category{
			ID:   1,
			Name: "Category Name1",
		},
	}
	tags := []string{"tag1", "tag2", "tag3"}
	expectedError := errors.New("expected error")

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)
	repo.EXPECT().FindByID(expectedItemID).Return(item, nil)
	repo.EXPECT().Save(item).Return(expectedError)

	s := services.NewKnowledgeItemService(repo)
	_, err := s.UpdateItem(expectedItemID, expectedTitle, expectedAnchor, expectedData, tags, categories)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != expectedError.Error() {
		t.Fatalf("expected error: %s, got: %s", expectedError.Error(), err.Error())
	}
}

func TestKnowledgeItemService_UpdateItem_RepoNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expectedItemID int64 = 5
	item := &models.KnowledgeItem{
		ID:     expectedItemID,
		Title:  "Test Item",
		Anchor: "Test Anchor",
		Data:   "Test Data and Something more",
		Tags:   []string{"tag1", "tag2"},
	}
	expectedTitle := "expected title"
	expectedAnchor := "expected anchor"
	expectedData := "expected data and something more"
	categories := []*models.Category{
		&models.Category{
			ID:   1,
			Name: "Category Name1",
		},
	}
	tags := []string{"tag1", "tag2", "tag3"}
	expectedError := errors.New("expected not found error")

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)
	repo.EXPECT().FindByID(expectedItemID).Return(item, expectedError)

	s := services.NewKnowledgeItemService(repo)
	_, err := s.UpdateItem(expectedItemID, expectedTitle, expectedAnchor, expectedData, tags, categories)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != expectedError.Error() {
		t.Fatalf("expected error: %s, got: %s", expectedError.Error(), err.Error())
	}
}

func TestKnowledgeItemService_UpdateItem_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	item := &models.KnowledgeItem{
		Title:  "Test Item",
		Anchor: "Test Anchor",
		Data:   "Test Data and Something more",
		Tags:   []string{"tag1", "tag2"},
	}

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)

	testCases := []struct {
		name               string
		item               *models.KnowledgeItem
		expectedItemID     int64
		expectedTitle      string
		expectedAnchor     string
		expectedData       string
		expectedTags       []string
		expectedCategories []*models.Category
		expectedError      error
	}{
		{
			name:               "title too short",
			item:               item,
			expectedItemID:     15,
			expectedTitle:      "e", // too short
			expectedAnchor:     "expectedAnchor",
			expectedData:       "expectedData and something",
			expectedTags:       []string{"expectedTag1", "expectedTag2", "expectedTag3"},
			expectedCategories: make([]*models.Category, 0),
			expectedError:      errors.New("title is too short"),
		},
		{
			name:               "anchor too short",
			item:               item,
			expectedItemID:     15,
			expectedTitle:      "expectedTitle",
			expectedAnchor:     "e", // too short
			expectedData:       "expectedData and something",
			expectedTags:       []string{"expectedTag1", "expectedTag2", "expectedTag3"},
			expectedCategories: make([]*models.Category, 0),
			expectedError:      errors.New("anchor is too short"),
		},
		{
			name:               "data too short",
			item:               item,
			expectedItemID:     15,
			expectedTitle:      "expectedTitle",
			expectedAnchor:     "expectedAnchor",
			expectedData:       "expectedData", // too short
			expectedTags:       []string{"expectedTag1", "expectedTag2", "expectedTag3"},
			expectedCategories: make([]*models.Category, 0),
			expectedError:      errors.New("data is too short"),
		},
		{
			name:               "tag too short",
			item:               item,
			expectedItemID:     15,
			expectedTitle:      "expectedTitle",
			expectedAnchor:     "expectedAnchor",
			expectedData:       "expectedData and something",
			expectedTags:       []string{"e", "expectedTag2", "expectedTag3"}, // too short
			expectedCategories: make([]*models.Category, 0),
			expectedError:      errors.New("tag is too short"),
		},
		{
			name:               "category cannot be empty",
			item:               item,
			expectedItemID:     15,
			expectedTitle:      "expectedTitle",
			expectedAnchor:     "expectedAnchor",
			expectedData:       "expectedData and something",
			expectedTags:       []string{"expectedTag1", "expectedTag2", "expectedTag3"},
			expectedCategories: []*models.Category{nil}, // empty category
			expectedError:      errors.New("category cannot be empty"),
		},
		{
			name:           "category doesn't exist",
			item:           item,
			expectedItemID: 15,
			expectedTitle:  "expectedTitle",
			expectedAnchor: "expectedAnchor",
			expectedData:   "expectedData and something",
			expectedTags:   []string{"expectedTag1", "expectedTag2", "expectedTag3"},
			expectedCategories: []*models.Category{&models.Category{
				ID:   0, // not exists
				Name: "Category Name",
			}},
			expectedError: errors.New("category doesn't exist"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo.EXPECT().FindByID(tc.expectedItemID).Return(tc.item, nil)
			s := services.NewKnowledgeItemService(repo)
			_, err := s.UpdateItem(
				tc.expectedItemID, tc.expectedTitle, tc.expectedAnchor,
				tc.expectedData, tc.expectedTags, tc.expectedCategories)
			if err.Error() != tc.expectedError.Error() {
				t.Fatalf("expected error: %s, got: %s", tc.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestKnowledgeItemService_DeleteItem_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expectedItemID int64 = 5
	item := &models.KnowledgeItem{
		ID:     expectedItemID,
		Title:  "Test Item",
		Anchor: "Test Anchor",
		Data:   "Test Data and Something more",
		Tags:   []string{"tag1", "tag2"},
	}

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)
	repo.EXPECT().FindByID(expectedItemID).Return(item, nil)
	repo.EXPECT().Delete(item).Return(nil)

	s := services.NewKnowledgeItemService(repo)
	err := s.DeleteItem(expectedItemID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKnowledgeItemService_DeleteItem_ItemNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expectedItemID int64 = 5
	expectedError := errors.New("not found error")

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)
	repo.EXPECT().FindByID(expectedItemID).Return(nil, expectedError)

	s := services.NewKnowledgeItemService(repo)
	err := s.DeleteItem(expectedItemID)
	if err == nil {
		t.Fatal(err)
	}
	if err.Error() != expectedError.Error() {
		t.Fatalf("expected error: %s, got: %s", expectedError.Error(), err.Error())
	}
}

func TestKnowledgeItemService_DeleteItem_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expectedItemID int64 = 5
	item := &models.KnowledgeItem{
		ID:     expectedItemID,
		Title:  "Test Item",
		Anchor: "Test Anchor",
		Data:   "Test Data and Something more",
		Tags:   []string{"tag1", "tag2"},
	}
	expectedError := errors.New("expected error")

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)
	repo.EXPECT().FindByID(expectedItemID).Return(item, nil)
	repo.EXPECT().Delete(item).Return(expectedError)

	s := services.NewKnowledgeItemService(repo)
	err := s.DeleteItem(expectedItemID)
	if err == nil {
		t.Fatal(err)
	}
	if err.Error() != expectedError.Error() {
		t.Fatalf("expected error: %s, got: %s", expectedError.Error(), err.Error())
	}
}

func TestKnowledgeItemService_SetLatestMark(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	item := &models.KnowledgeItem{
		ID:     0,
		Title:  "Test Item",
		Anchor: "Test Anchor",
		Data:   "Test Data and Something more",
		Tags:   []string{"tag1", "tag2"},
	}

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)
	s := services.NewKnowledgeItemService(repo)

	testCases := []struct {
		name          string
		item          *models.KnowledgeItem
		itemID        int64
		mark          int64
		exScore       int64
		exMark        int64
		expectedMark  int64
		expectedScore int64
		expectedError error
	}{
		{
			name:          "mark less than min 0",
			item:          item,
			itemID:        1,
			mark:          -1,
			exScore:       0,
			exMark:        0,
			expectedMark:  0,
			expectedScore: 0,
			expectedError: fmt.Errorf("mark cannot be less than %d", 0),
		},
		{
			name:          "mark is higher than max 10",
			item:          item,
			itemID:        1,
			mark:          11,
			exScore:       0,
			exMark:        0,
			expectedMark:  0,
			expectedScore: 0,
			expectedError: fmt.Errorf("mark cannot be more than %d", 10),
		},
		{
			name:          "first mark",
			item:          item,
			itemID:        1,
			mark:          6,
			exScore:       0,
			exMark:        0,
			expectedMark:  6,
			expectedScore: 6,
			expectedError: nil,
		},
		{
			name:          "mark higher than previous",
			item:          item,
			itemID:        1,
			mark:          6,
			exScore:       25,
			exMark:        5,
			expectedMark:  6,
			expectedScore: 31,
			expectedError: nil,
		},
		{
			name:          "mark less than previous",
			item:          item,
			itemID:        1,
			mark:          4,
			exScore:       25,
			exMark:        5,
			expectedMark:  4,
			expectedScore: 24,
			expectedError: nil,
		},
		{
			name:          "mark eq previous",
			item:          item,
			itemID:        1,
			mark:          5,
			exScore:       25,
			exMark:        5,
			expectedMark:  5,
			expectedScore: 30,
			expectedError: nil,
		},
		{
			name:          "completely forgotten",
			item:          item,
			itemID:        1,
			mark:          0,
			exScore:       25,
			exMark:        5,
			expectedMark:  0,
			expectedScore: 0,
			expectedError: nil,
		},
		{
			name:          "score >= 0",
			item:          item,
			itemID:        1,
			mark:          1,
			exScore:       1,
			exMark:        4,
			expectedMark:  1,
			expectedScore: 0,
			expectedError: nil,
		},
		{
			name:          "score <= 100",
			item:          item,
			itemID:        1,
			mark:          10,
			exScore:       95,
			exMark:        7,
			expectedMark:  10,
			expectedScore: 100,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.item.ID = tc.itemID
			tc.item.Score = tc.exScore
			tc.item.LastMark = tc.exMark

			repo.EXPECT().FindByID(tc.itemID).Return(tc.item, nil)
			if tc.expectedError == nil {
				repo.EXPECT().Save(tc.item).Return(nil)
			}

			resultItem, err := s.SetLatestMark(tc.itemID, tc.mark)
			// error expected
			if tc.expectedError != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Fatalf("expected error: %s, got: %s", tc.expectedError.Error(), err.Error())
				}
				return
			}

			// error not expected
			if err != nil {
				t.Fatal(err)
			}
			if tc.expectedScore != resultItem.Score {
				t.Fatalf("expected score: %d, got: %d", tc.expectedScore, resultItem.Score)
			}
			if tc.expectedMark != resultItem.LastMark {
				t.Fatalf("expected mark: %d, got: %d", tc.expectedMark, resultItem.LastMark)
			}
		})
	}
}

func TestKnowledgeItemService_SetLatestMark_NotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expectedItemID int64 = 51
	item := &models.KnowledgeItem{
		ID:     expectedItemID,
		Title:  "Test Item",
		Anchor: "Test Anchor",
		Data:   "Test Data and Something more",
		Tags:   []string{"tag1", "tag2"},
	}
	expectedError := errors.New("item not found")

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)
	repo.EXPECT().FindByID(expectedItemID).Return(item, expectedError)

	s := services.NewKnowledgeItemService(repo)
	_, err := s.SetLatestMark(expectedItemID, 5)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != expectedError.Error() {
		t.Fatalf("expected error: %s, got: %s", expectedError.Error(), err.Error())
	}
}

func TestKnowledgeItemService_SetLatestMark_SaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expectedItemID int64 = 51
	var expectedMark int64 = 5
	item := &models.KnowledgeItem{
		ID:     expectedItemID,
		Title:  "Test Item",
		Anchor: "Test Anchor",
		Data:   "Test Data and Something more",
		Tags:   []string{"tag1", "tag2"},
	}
	expectedError := errors.New("item not saved")

	repo := mock.NewMockKnowledgeItemsRepo(ctrl)
	repo.EXPECT().FindByID(expectedItemID).Return(item, nil)
	repo.EXPECT().Save(item).DoAndReturn(func(i *models.KnowledgeItem) error {
		if i.ID != expectedItemID {
			t.Fatalf("expected ID: %d, got: %d", expectedItemID, i.ID)
		}
		if i.LastMark != expectedMark {
			t.Fatalf("expected LastMark: %d, got: %d", expectedMark, i.LastMark)
		}

		return expectedError
	})

	s := services.NewKnowledgeItemService(repo)
	_, err := s.SetLatestMark(expectedItemID, expectedMark)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != expectedError.Error() {
		t.Fatalf("expected error: %s, got: %s", expectedError.Error(), err.Error())
	}
}
