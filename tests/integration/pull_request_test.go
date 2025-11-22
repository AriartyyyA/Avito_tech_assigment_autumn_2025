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

// TestCreatePullRequest тестирует создание PR
func TestCreatePullRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	expectedPR := &models.PullRequest{
		PullRequestID:     "pr-3",
		PullRequestName:   "New Feature",
		AuthorID:          "u11",
		Status:            models.PullRequestStatusOpen,
		AssignedReviewers: []string{"u12", "u13"},
	}

	mockSvc.PullRequest.(*MockPullRequestService).CreatePullRequestFunc = func(ctx context.Context, pullRequestID, pullRequestName, authorID string) (*models.PullRequest, error) {
		if pullRequestID == "pr-3" {
			return expectedPR, nil
		}
		return nil, models.ErrorCodeUserNotFound
	}

	handler := setupTestHandler(mockSvc)
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

// TestCreatePullRequestDuplicate тестирует создание дублирующегося PR
func TestCreatePullRequestDuplicate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	mockSvc.PullRequest.(*MockPullRequestService).CreatePullRequestFunc = func(ctx context.Context, pullRequestID, pullRequestName, authorID string) (*models.PullRequest, error) {
		return nil, models.ErrorCodePRExists
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"pull_request_id":   "pr-4",
		"pull_request_name": "Feature",
		"author_id":         "u14",
	}

	w := makeRequest(router, "POST", "/pullRequest/create", reqBody)

	if w.Code != http.StatusConflict {
		t.Errorf("Expected status 409 for duplicate PR, got %d", w.Code)
	}
}

// TestCreatePullRequestAuthorNotFound тестирует создание PR с несуществующим автором
func TestCreatePullRequestAuthorNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	mockSvc.PullRequest.(*MockPullRequestService).CreatePullRequestFunc = func(ctx context.Context, pullRequestID, pullRequestName, authorID string) (*models.PullRequest, error) {
		return nil, models.ErrorCodeUserNotFound
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"pull_request_id":   "pr-5",
		"pull_request_name": "Feature",
		"author_id":         "nonexistent",
	}

	w := makeRequest(router, "POST", "/pullRequest/create", reqBody)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

// TestMergePullRequest тестирует слияние PR
func TestMergePullRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	mergedAt := time.Now()
	expectedPR := &models.PullRequest{
		PullRequestID:     "pr-5",
		PullRequestName:   "Bug Fix",
		AuthorID:          "u15",
		Status:            models.PullRequestStatusMerged,
		AssignedReviewers: []string{"u16"},
		MergedAt:          &mergedAt,
	}

	mockSvc.PullRequest.(*MockPullRequestService).MergePullRequestFunc = func(ctx context.Context, prID string) (*models.PullRequest, error) {
		if prID == "pr-5" {
			return expectedPR, nil
		}
		return nil, models.ErrorCodePRNotFound
	}

	handler := setupTestHandler(mockSvc)
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

// TestMergePullRequestNotFound тестирует слияние несуществующего PR
func TestMergePullRequestNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	mockSvc.PullRequest.(*MockPullRequestService).MergePullRequestFunc = func(ctx context.Context, prID string) (*models.PullRequest, error) {
		return nil, models.ErrorCodePRNotFound
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"pull_request_id": "nonexistent",
	}

	w := makeRequest(router, "POST", "/pullRequest/merge", reqBody)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

// TestReassignPullRequest тестирует переназначение ревьювера
func TestReassignPullRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	expectedPR := &models.PullRequest{
		PullRequestID:     "pr-6",
		PullRequestName:   "Refactor",
		AuthorID:          "u17",
		Status:            models.PullRequestStatusOpen,
		AssignedReviewers: []string{"u19"},
		NewReviewerID:     "u19",
	}

	mockSvc.PullRequest.(*MockPullRequestService).ReassignPullRequestFunc = func(ctx context.Context, prID string, oldReviewerID string) (*models.PullRequest, error) {
		if prID == "pr-6" && oldReviewerID == "u18" {
			return expectedPR, nil
		}
		return nil, models.ErrorCodeNotAssigned
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"pull_request_id": "pr-6",
		"old_reviewer_id": "u18",
	}

	w := makeRequest(router, "POST", "/pullRequest/reassign", reqBody)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// TestReassignMergedPullRequest тестирует попытку переназначения слитого PR
func TestReassignMergedPullRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	mockSvc.PullRequest.(*MockPullRequestService).ReassignPullRequestFunc = func(ctx context.Context, prID string, oldReviewerID string) (*models.PullRequest, error) {
		return nil, models.ErrorCodePRMerged
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"pull_request_id": "pr-7",
		"old_reviewer_id": "u21",
	}

	w := makeRequest(router, "POST", "/pullRequest/reassign", reqBody)

	if w.Code != http.StatusConflict {
		t.Errorf("Expected status 409 for merged PR, got %d", w.Code)
	}
}

// TestReassignNotAssigned тестирует переназначение не назначенного ревьювера
func TestReassignNotAssigned(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	mockSvc.PullRequest.(*MockPullRequestService).ReassignPullRequestFunc = func(ctx context.Context, prID string, oldReviewerID string) (*models.PullRequest, error) {
		return nil, models.ErrorCodeNotAssigned
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"pull_request_id": "pr-8",
		"old_reviewer_id": "u25",
	}

	w := makeRequest(router, "POST", "/pullRequest/reassign", reqBody)

	if w.Code != http.StatusConflict {
		t.Errorf("Expected status 409 for not assigned reviewer, got %d", w.Code)
	}
}

// TestReassignNoCandidate тестирует переназначение когда нет кандидатов
func TestReassignNoCandidate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	mockSvc.PullRequest.(*MockPullRequestService).ReassignPullRequestFunc = func(ctx context.Context, prID string, oldReviewerID string) (*models.PullRequest, error) {
		return nil, models.ErrorCodeNoCandidate
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"pull_request_id": "pr-9",
		"old_reviewer_id": "u26",
	}

	w := makeRequest(router, "POST", "/pullRequest/reassign", reqBody)

	if w.Code != http.StatusConflict {
		t.Errorf("Expected status 409 for no candidate, got %d", w.Code)
	}
}

