package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/gin-gonic/gin"
)

func TestIntegrationSetIsActive(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	expectedUser := &models.User{
		ID:       "u6",
		Name:     "Frank",
		TeamName: "qa",
		IsActive: false,
	}

	mockRepo.UserRepository.(*MockUserRepository).SetIsActiveFunc = func(ctx context.Context, userID string, isActive bool) (*models.User, error) {
		if userID == "u6" && !isActive {
			return expectedUser, nil
		}
		return nil, models.ErrorCodeUserNotFound
	}

	handler := setupIntegrationTestHandler(mockRepo)
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

func TestIntegrationGetReview(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	expectedPRs := []models.PullRequestShort{
		{
			PullRequestID:   "pr-1",
			PullRequestName: "Feature A",
			AuthorID:        "u1",
			Status:          models.PullRequestStatusOpen,
		},
		{
			PullRequestID:   "pr-2",
			PullRequestName: "Feature B",
			AuthorID:        "u2",
			Status:          models.PullRequestStatusOpen,
		},
	}

	mockRepo.UserRepository.(*MockUserRepository).GetReviewFunc = func(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
		if userID == "u10" {
			return expectedPRs, nil
		}
		return nil, models.ErrorCodeUserNotFound
	}

	handler := setupIntegrationTestHandler(mockRepo)
	router := handler.InitRoutes()

	w := makeRequest(router, "GET", "/users/getReview?user_id=u10", nil)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// getReview возвращает GetUsersReviewResponse с полем "pull_request" (единственное число)
	if prs, ok := result["pull_request"].([]interface{}); ok {
		if len(prs) != 2 {
			t.Errorf("Expected 2 PRs, got %d", len(prs))
		}
	} else {
		t.Error("Response does not contain 'pull_request' field")
	}
}

func TestIntegrationGetUserAssignmentsStats(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	expectedStats := []models.UserAssignmentsStat{
		{
			UserID:                 "u1",
			Username:               "Alice",
			ReviewAssignmentsCount: 5,
		},
		{
			UserID:                 "u2",
			Username:               "Bob",
			ReviewAssignmentsCount: 3,
		},
	}

	mockRepo.UserRepository.(*MockUserRepository).GetAssignmentsStatsFunc = func(ctx context.Context) ([]models.UserAssignmentsStat, error) {
		return expectedStats, nil
	}

	handler := setupIntegrationTestHandler(mockRepo)
	router := handler.InitRoutes()

	w := makeRequest(router, "GET", "/users/userAssignments", nil)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if stats, ok := result["stats"].([]interface{}); ok {
		if len(stats) != 2 {
			t.Errorf("Expected 2 stats, got %d", len(stats))
		}
	} else {
		t.Error("Response does not contain 'stats' field")
	}
}
