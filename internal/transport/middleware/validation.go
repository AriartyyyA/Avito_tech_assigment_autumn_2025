package middleware

import (
	"fmt"
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
			errResp := invalidRequest("")
			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.UserID == "" {
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
			errResp := invalidRequest("")
			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.TeamName == "" {
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
			errResp := invalidRequest("")
			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.TeamName == "" {
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
			errResp := invalidRequest("")
			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.PullRequestID == "" {
			errResp := invalidRequest("pull_request_id")
			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.PullRequestName == "" {
			errResp := invalidRequest("pull_request_name")
			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.AuthorID == "" {
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
			errResp := invalidRequest("")
			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.PullRequestID == "" {
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
			errResp := invalidRequest("")
			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.PullRequestID == "" {
			errResp := invalidRequest("pull_request_id")
			c.JSON(http.StatusBadRequest, errResp)
			c.Abort()
			return
		}

		if req.OldUserID == "" {
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
				errResp := invalidRequest(param)
				c.JSON(http.StatusBadRequest, errResp)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
