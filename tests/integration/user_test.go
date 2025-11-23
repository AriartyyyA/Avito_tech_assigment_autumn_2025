package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/gin-gonic/gin"
)

func TestSetIsActive(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	expectedUser := &models.User{
		ID:       "u6",
		Name:     "Frank",
		TeamName: "qa",
		IsActive: false,
	}

	mockSvc.User.(*MockUserService).SetIsActiveFunc = func(ctx context.Context, userID string, isActive bool) (*models.User, error) {
		if userID == "u6" && !isActive {
			return expectedUser, nil
		}
		return nil, models.ErrorCodeUserNotFound
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"user_id":   "u6",
		"is_active": false,
	}

	w := makeRequest(router, "POST", "/users/setIsActive", reqBody)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if user, ok := result["user"].(map[string]interface{}); ok {
		if user["is_active"] != false {
			t.Errorf("Expected is_active false, got %v", user["is_active"])
		}
	}
}

func TestSetIsActiveNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	mockSvc.User.(*MockUserService).SetIsActiveFunc = func(ctx context.Context, userID string, isActive bool) (*models.User, error) {
		return nil, models.ErrorCodeUserNotFound
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"user_id":   "nonexistent",
		"is_active": false,
	}

	w := makeRequest(router, "POST", "/users/setIsActive", reqBody)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestGetReview(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	expectedPRs := []models.PullRequestShort{
		{
			PullRequestID:   "pr-2",
			PullRequestName: "Feature PR",
			AuthorID:        "u7",
			Status:          models.PullRequestStatusOpen,
		},
	}

	mockSvc.User.(*MockUserService).GetReviewFunc = func(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
		if userID == "u8" {
			return expectedPRs, nil
		}
		return nil, models.ErrorCodeUserNotFound
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	w := makeRequest(router, "GET", "/users/getReview?user_id=u8", nil)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetUserAssignmentsStats(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	expectedStats := []models.UserAssignmentsStat{
		{
			UserID:                 "u9",
			Username:               "Ivan",
			ReviewAssignmentsCount: 5,
		},
		{
			UserID:                 "u10",
			Username:               "John",
			ReviewAssignmentsCount: 3,
		},
	}

	mockSvc.User.(*MockUserService).GetUserAssignmentsStatsFunc = func(ctx context.Context) ([]models.UserAssignmentsStat, error) {
		return expectedStats, nil
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	w := makeRequest(router, "GET", "/users/userAssignments", nil)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
