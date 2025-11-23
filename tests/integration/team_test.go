package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/gin-gonic/gin"
)

// TestAddTeam тестирует создание команды
func TestAddTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	expectedTeam := &models.Team{
		TeamName: "backend",
		Members: []models.TeamMember{
			{UserID: "u1", Username: "Alice", IsActive: true},
			{UserID: "u2", Username: "Bob", IsActive: true},
		},
	}

	mockSvc.Team.(*MockTeamService).AddTeamFunc = func(ctx context.Context, team *models.Team) (*models.Team, error) {
		return expectedTeam, nil
	}

	handler := setupTestHandler(mockSvc)
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

func TestAddTeamDuplicate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	mockSvc.Team.(*MockTeamService).AddTeamFunc = func(ctx context.Context, team *models.Team) (*models.Team, error) {
		return nil, models.ErrorCodeTeamExists
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"team_name": "backend",
		"members": []map[string]interface{}{
			{"user_id": "u1", "username": "Alice", "is_active": true},
		},
	}

	w := makeRequest(router, "POST", "/team/add", reqBody)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for duplicate team, got %d", w.Code)
	}
}

func TestGetTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	expectedTeam := &models.Team{
		TeamName: "frontend",
		Members: []models.TeamMember{
			{UserID: "u3", Username: "Charlie", IsActive: true},
		},
	}

	mockSvc.Team.(*MockTeamService).GetTeamFunc = func(ctx context.Context, teamName string) (*models.Team, error) {
		if teamName == "frontend" {
			return expectedTeam, nil
		}
		return nil, models.ErrorCodeTeamNotFound
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	w := makeRequest(router, "GET", "/team/get?team_name=frontend", nil)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var result models.Team
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if result.TeamName != "frontend" {
		t.Errorf("Expected team_name 'frontend', got %v", result.TeamName)
	}
}

func TestGetTeamNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	mockSvc.Team.(*MockTeamService).GetTeamFunc = func(ctx context.Context, teamName string) (*models.Team, error) {
		return nil, models.ErrorCodeTeamNotFound
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	w := makeRequest(router, "GET", "/team/get?team_name=nonexistent", nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestGetTeamPullRequests(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	expectedPRs := []models.PullRequestShort{
		{
			PullRequestID:   "pr-1",
			PullRequestName: "Test PR",
			AuthorID:        "u4",
			Status:          models.PullRequestStatusOpen,
		},
	}

	mockSvc.Team.(*MockTeamService).GetTeamFunc = func(ctx context.Context, teamName string) (*models.Team, error) {
		return &models.Team{TeamName: "devops"}, nil
	}

	mockSvc.Team.(*MockTeamService).GetTeamPullRequestsFunc = func(ctx context.Context, teamName string) ([]models.PullRequestShort, error) {
		return expectedPRs, nil
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	w := makeRequest(router, "GET", "/team/pullRequests?team_name=devops", nil)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestDeactivateTeamUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	expectedResult := &models.TeamDeactivate{
		TeamName:                "backend",
		DeactivatedUsers:        []string{"u22", "u23"},
		OpenPRCount:             0,
		SuccessfulReassignments: 0,
		FailedReassignments:     0,
	}

	mockSvc.Team.(*MockTeamService).DeactivateTeamFunc = func(ctx context.Context, teamName string) (*models.TeamDeactivate, error) {
		return expectedResult, nil
	}

	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	reqBody := map[string]interface{}{
		"team_name": "backend",
	}

	w := makeRequest(router, "POST", "/team/deactivateUsers", reqBody)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
