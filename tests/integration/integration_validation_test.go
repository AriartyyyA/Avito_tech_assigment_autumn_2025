package integration

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIntegrationAddTeamInvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	handler := setupIntegrationTestHandler(mockRepo)
	router := handler.InitRoutes()

	// Пустое тело запроса
	w := makeRequest(router, "POST", "/team/add", nil)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for invalid request, got %d", w.Code)
	}
}

func TestIntegrationCreatePullRequestInvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	handler := setupIntegrationTestHandler(mockRepo)
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

func TestIntegrationGetTeamInvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := NewMockRepository()
	handler := setupIntegrationTestHandler(mockRepo)
	router := handler.InitRoutes()

	w := makeRequest(router, "GET", "/team/get", nil)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for invalid request, got %d", w.Code)
	}
}
