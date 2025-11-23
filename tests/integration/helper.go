package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/service"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport"
	"github.com/gin-gonic/gin"
)

func setupTestHandler(mockSvc *MockService) *transport.Handler {

	testService := &service.Service{
		User:        mockSvc.User,
		PullRequest: mockSvc.PullRequest,
		Team:        mockSvc.Team,
	}

	return transport.NewHandler(testService)
}

func makeRequest(router *gin.Engine, method, url string, body interface{}) *httptest.ResponseRecorder {
	var reqBody []byte
	if body != nil {
		reqBody, _ = json.Marshal(body)
	}

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
