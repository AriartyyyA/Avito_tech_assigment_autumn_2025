package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/gin-gonic/gin"
)

func TestIntegrationCreatePullRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	expectedPR := &models.PullRequest{
		PullRequestID:     "pr-3",
		PullRequestName:   "New Feature",
		AuthorID:          "u11",
		Status:            models.PullRequestStatusOpen,
		AssignedReviewers: []string{"u12", "u13"},
	}

	mockRepo.PullRequestRepository.(*MockPullRequestRepository).CreatePullRequestFunc = func(ctx context.Context, pr *models.PullRequest) (*models.PullRequest, error) {
		if pr.PullRequestID == "pr-3" {
			return expectedPR, nil
		}
		return nil, models.ErrorCodeUserNotFound
	}

	handler := setupIntegrationTestHandler(mockRepo)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"pull_request_id":   "pr-3",
		"pull_request_name": "New Feature",
		"author_id":         "u11",
	}

	w := makeRequest(router, "POST", "/pullRequest/create", reqBody)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if prData, ok := result["pr"].(map[string]interface{}); ok {
		if prData["status"] != "OPEN" {
			t.Errorf("Expected status 'OPEN', got %v", prData["status"])
		}
		if reviewers, ok := prData["assigned_reviewers"].([]interface{}); ok {
			if len(reviewers) == 0 || len(reviewers) > 2 {
				t.Errorf("Expected 1-2 reviewers, got %d", len(reviewers))
			}
		}
	}
}

func TestIntegrationMergePullRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	mergedAt := time.Now()
	expectedPR := &models.PullRequest{
		PullRequestID:     "pr-5",
		PullRequestName:   "Feature",
		AuthorID:          "u1",
		Status:            models.PullRequestStatusMerged,
		AssignedReviewers: []string{"u2", "u3"},
		MergedAt:          &mergedAt,
	}

	mockRepo.PullRequestRepository.(*MockPullRequestRepository).MergePullRequestFunc = func(ctx context.Context, prID string) (*models.PullRequest, error) {
		if prID == "pr-5" {
			return expectedPR, nil
		}
		return nil, models.ErrorCodePRNotFound
	}

	handler := setupIntegrationTestHandler(mockRepo)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"pull_request_id": "pr-5",
	}

	w := makeRequest(router, "POST", "/pullRequest/merge", reqBody)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if prData, ok := result["pr"].(map[string]interface{}); ok {
		if prData["status"] != "MERGED" {
			t.Errorf("Expected status 'MERGED', got %v", prData["status"])
		}
	}
}

func TestIntegrationReassignPullRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	expectedPR := &models.PullRequest{
		PullRequestID:     "pr-6",
		PullRequestName:   "Feature",
		AuthorID:          "u1",
		Status:            models.PullRequestStatusOpen,
		AssignedReviewers: []string{"u4", "u5"},
	}

	mockRepo.PullRequestRepository.(*MockPullRequestRepository).ReassignPullRequestFunc = func(ctx context.Context, prID string, oldUserID string) (*models.PullRequest, error) {
		if prID == "pr-6" && oldUserID == "u2" {
			return expectedPR, nil
		}
		return nil, models.ErrorCodePRNotFound
	}

	handler := setupIntegrationTestHandler(mockRepo)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"pull_request_id": "pr-6",
		"old_user_id":     "u2",
	}

	w := makeRequest(router, "POST", "/pullRequest/reassign", reqBody)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d. Body: %s", w.Code, w.Body.String())
	}

	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if prData, ok := result["pr"].(map[string]interface{}); ok {
		if prData["status"] != "OPEN" {
			t.Errorf("Expected status 'OPEN', got %v", prData["status"])
		}
		if reviewers, ok := prData["assigned_reviewers"].([]interface{}); ok {
			if len(reviewers) != 2 {
				t.Errorf("Expected 2 reviewers, got %d", len(reviewers))
			}
		}
	}
}

func TestIntegrationReassignMergedPullRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()

	mockRepo.PullRequestRepository.(*MockPullRequestRepository).ReassignPullRequestFunc = func(ctx context.Context, prID string, oldUserID string) (*models.PullRequest, error) {
		return nil, models.ErrorCodePRMerged
	}

	handler := setupIntegrationTestHandler(mockRepo)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"pull_request_id": "pr-7",
		"old_user_id":     "u2",
	}

	w := makeRequest(router, "POST", "/pullRequest/reassign", reqBody)

	if w.Code != http.StatusConflict {
		t.Errorf("Expected status 409 for merged PR, got %d", w.Code)
	}
}
