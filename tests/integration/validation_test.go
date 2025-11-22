package integration

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestAddTeamInvalidRequest тестирует валидацию запроса создания команды
func TestAddTeamInvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	// Пустое тело запроса
	w := makeRequest(router, "POST", "/team/add", nil)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for invalid request, got %d", w.Code)
	}
}

// TestCreatePullRequestInvalidRequest тестирует валидацию запроса создания PR
func TestCreatePullRequestInvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	// Запрос без обязательных полей
	reqBody := map[string]interface{}{
		"pull_request_id": "pr-10",
		// отсутствует pull_request_name и author_id
	}

	w := makeRequest(router, "POST", "/pullRequest/create", reqBody)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for invalid request, got %d", w.Code)
	}
}

// TestGetTeamInvalidRequest тестирует валидацию запроса получения команды
func TestGetTeamInvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := NewMockService()
	handler := setupTestHandler(mockSvc)
	router := handler.InitRoutes()

	// Запрос без параметра team_name
	w := makeRequest(router, "GET", "/team/get", nil)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for invalid request, got %d", w.Code)
	}
}

