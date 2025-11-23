package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

// invalidRequest создает ошибку валидации

func invalidRequest(item string) models.ErrorResponse {
	errStr := fmt.Sprintf("%s is required", item)

	if item == "" {
		errStr = "invalid request"
	}

	error := models.NewErrorDetail(
		models.ErrorCodeInvalidRequest,
		errStr,
	)

	return models.ErrorResponse{
		Error: error,
	}

}

func ValidateSetIsActiveRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.SetUserIsActiveRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("WARN: Invalid JSON in setIsActive request: Error=%v, Path=%s", err, c.Request.URL.Path)
			errResp := invalidRequest("")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return

		}

		if req.UserID == "" {
			log.Printf("WARN: Missing user_id in setIsActive request: Path=%s", c.Request.URL.Path)
			errResp := invalidRequest("user_id")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return

		}

		c.Set("validated_request", &req)
		c.Next()

	}

}

func ValidateAddTeamRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Team
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("WARN: Invalid JSON in addTeam request: Error=%v, Path=%s", err, c.Request.URL.Path)
			errResp := invalidRequest("")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return

		}

		if req.TeamName == "" {
			log.Printf("WARN: Missing team_name in addTeam request: Path=%s", c.Request.URL.Path)
			errResp := invalidRequest("team_name")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return

		}
		c.Set("validated_request", &req)
		c.Next()

	}

}

func ValidateDeactivateTeamUsersRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.DeactivateTeamUsersRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("WARN: Invalid JSON in deactivateTeamUsers request: Error=%v, Path=%s", err, c.Request.URL.Path)
			errResp := invalidRequest("")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.TeamName == "" {
			log.Printf("WARN: Missing team_name in deactivateTeamUsers request: Path=%s", c.Request.URL.Path)
			errResp := invalidRequest("team_name")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		c.Set("validated_request", &req)
		c.Next()

	}

}

func ValidateCreatePRRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreatePRRequestDto
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("WARN: Invalid JSON in createPR request: Error=%v, Path=%s", err, c.Request.URL.Path)
			errResp := invalidRequest("")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.PullRequestID == "" {
			log.Printf("WARN: Missing pull_request_id in createPR request: Path=%s", c.Request.URL.Path)
			errResp := invalidRequest("pull_request_id")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.PullRequestName == "" {
			log.Printf("WARN: Missing pull_request_name in createPR request: Path=%s", c.Request.URL.Path)
			errResp := invalidRequest("pull_request_name")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return

		}

		if req.AuthorID == "" {
			log.Printf("WARN: Missing author_id in createPR request: Path=%s", c.Request.URL.Path)
			errResp := invalidRequest("author_id")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return

		}

		c.Set("validated_request", &req)
		c.Next()

	}

}

func ValidateMergePRRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.MergePRRequestDto
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("WARN: Invalid JSON in mergePR request: Error=%v, Path=%s", err, c.Request.URL.Path)
			errResp := invalidRequest("")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.PullRequestID == "" {
			log.Printf("WARN: Missing pull_request_id in mergePR request: Path=%s", c.Request.URL.Path)
			errResp := invalidRequest("pull_request_id")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return

		}

		c.Set("validated_request", &req)
		c.Next()

	}

}

func ValidateReassignPRRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.ReassignPRRequestDto
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("WARN: Invalid JSON in reassignPR request: Error=%v, Path=%s", err, c.Request.URL.Path)
			errResp := invalidRequest("")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.PullRequestID == "" {
			log.Printf("WARN: Missing pull_request_id in reassignPR request: Path=%s", c.Request.URL.Path)
			errResp := invalidRequest("pull_request_id")

			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.OldUserID == "" {
			log.Printf("WARN: Missing old_user_id in reassignPR request: Path=%s", c.Request.URL.Path)
			errResp := invalidRequest("old_user_id")
			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return

		}
		c.Set("validated_request", &req)

		c.Next()

	}

}

func ValidateQueryParams(requiredParams ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, param := range requiredParams {
			value := c.Query(param)
			if value == "" {
				log.Printf("WARN: Missing query parameter in request: Param=%s, Path=%s", param, c.Request.URL.Path)
				errResp := invalidRequest(param)
				c.JSON(http.StatusBadRequest, errResp)
				c.Abort()
				return

			}

		}
		c.Next()

	}

}
