package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
	"github.com/gin-gonic/gin"
)

func TestIntegrationAddTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	expectedTeam := &models.Team{
		TeamName: "backend",
		Members: []models.TeamMember{
			{UserID: "u1", Username: "Alice", IsActive: true},
			{UserID: "u2", Username: "Bob", IsActive: true},
		},
	}

	mockRepo.TeamRepository.(*MockTeamRepository).AddTeamFunc = func(ctx context.Context, team *models.Team) (*models.Team, error) {
		return expectedTeam, nil
	}

	handler := setupIntegrationTestHandler(mockRepo)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"team_name": "backend",
		"members": []map[string]interface{}{
			{"user_id": "u1", "username": "Alice", "is_active": true},
			{"user_id": "u2", "username": "Bob", "is_active": true},
		},
	}

	w := makeRequest(router, "POST", "/team/add", reqBody)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if teamData, ok := result["team"].(map[string]interface{}); ok {
		if teamData["team_name"] != "backend" {
			t.Errorf("Expected team_name 'backend', got %v", teamData["team_name"])
		}
	} else {
		t.Error("Response does not contain 'team' field")
	}
}

func TestIntegrationGetTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	expectedTeam := &models.Team{
		TeamName: "frontend",
		Members: []models.TeamMember{
			{UserID: "u3", Username: "Charlie", IsActive: true},
		},
	}

	mockRepo.TeamRepository.(*MockTeamRepository).GetTeamFunc = func(ctx context.Context, teamName string) (*models.Team, error) {
		if teamName == "frontend" {
			return expectedTeam, nil
		}
		return nil, models.ErrorCodeTeamNotFound
	}

	handler := setupIntegrationTestHandler(mockRepo)
	router := handler.InitRoutes()

	w := makeRequest(router, "GET", "/team/get?team_name=frontend", nil)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// getTeam возвращает Team напрямую, без обертки
	if result["team_name"] != "frontend" {
		t.Errorf("Expected team_name 'frontend', got %v", result["team_name"])
	}
}

func TestIntegrationDeactivateTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	teamName := "backend"

	team := &models.Team{
		TeamName: teamName,
		Members: []models.TeamMember{
			{UserID: "u1", Username: "Alice", IsActive: true},
			{UserID: "u2", Username: "Bob", IsActive: true},
			{UserID: "u3", Username: "Charlie", IsActive: false},
		},
	}

	mockRepo.TeamRepository.(*MockTeamRepository).GetTeamFunc = func(ctx context.Context, name string) (*models.Team, error) {
		if name == teamName {
			return team, nil
		}
		return nil, models.ErrorCodeTeamNotFound
	}

	prsWithReviewers := []repository.PRWithReviewer{
		{PullRequestID: "pr-1", ReviewerID: "u1", AuthorID: "u4"},
		{PullRequestID: "pr-2", ReviewerID: "u2", AuthorID: "u5"},
	}

	mockRepo.PullRequestRepository.(*MockPullRequestRepository).GetOpenPRsWithTeamReviewersFunc = func(ctx context.Context, name string, userIDs []string) ([]repository.PRWithReviewer, error) {
		if name == teamName {
			return prsWithReviewers, nil
		}
		return nil, nil
	}

	reassignedPR := &models.PullRequest{
		PullRequestID:     "pr-1",
		PullRequestName:   "Feature",
		AuthorID:          "u4",
		Status:            models.PullRequestStatusOpen,
		AssignedReviewers: []string{"u5"},
	}

	mockRepo.PullRequestRepository.(*MockPullRequestRepository).ReassignPullRequestFunc = func(ctx context.Context, prID string, oldUserID string) (*models.PullRequest, error) {
		return reassignedPR, nil
	}

	mockRepo.UserRepository.(*MockUserRepository).DeactivateUsersFunc = func(ctx context.Context, userIDs []string) error {
		if len(userIDs) != 2 {
			t.Errorf("Expected 2 user IDs, got %d", len(userIDs))
		}
		return nil
	}

	handler := setupIntegrationTestHandler(mockRepo)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"team_name": teamName,
	}

	w := makeRequest(router, "POST", "/team/deactivateUsers", reqBody)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d. Body: %s", w.Code, w.Body.String())
	}

	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// deactivateTeamUsers возвращает DeactivateTeamUsersResponse напрямую
	if result["team_name"] != teamName {
		t.Errorf("Expected team_name '%s', got %v", teamName, result["team_name"])
	}
	if deactivatedUsers, ok := result["deactivated_users"].([]interface{}); ok {
		if len(deactivatedUsers) != 2 {
			t.Errorf("Expected 2 deactivated users, got %d", len(deactivatedUsers))
		}
	} else {
		t.Error("Response does not contain 'deactivated_users' field")
	}
}
