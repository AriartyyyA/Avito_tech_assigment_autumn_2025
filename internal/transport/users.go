package transport

import (
	"errors"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) setIsActive(c *gin.Context) {
	req, exists := c.Get("validated_request")
	if !exists {
		err := InternalError()
		c.JSON(500, err)
		return
	}

	setIsActiveReq := req.(*dto.SetUserIsActiveRequest)
	user, err := h.services.User.SetIsActive(c.Request.Context(), setIsActiveReq.UserID, setIsActiveReq.IsActive)
	if err != nil {
		if errors.Is(err, models.ErrorCodeUserNotFound) {
			err := NotFound(models.ErrorCodeUserNotFound)
			c.JSON(404, err)
			return
		}

		err := InternalError()
		c.JSON(500, err)
		return
	}

	resp := dto.SetUserIsActiveResponse{
		User: user,
	}

	c.JSON(200, resp)
}

func (h *Handler) getReview(c *gin.Context) {
	userID := c.Query("user_id")

	userPR, err := h.services.User.GetReview(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, models.ErrorCodeUserNotFound) {
			err := NotFound(models.ErrorCodeUserNotFound)
			c.JSON(404, err)
			return
		}

		err := InternalError()
		c.JSON(500, err)
		return
	}

	resp := dto.GetUsersReviewResponse{
		UserID:      userID,
		PullRequest: userPR,
	}

	c.JSON(200, resp)
}

func (h *Handler) getUserAssignmentsStats(c *gin.Context) {
	stats, err := h.services.User.GetUserAssignmentsStats(c.Request.Context())
	if err != nil {
		resp := InternalError()
		c.JSON(500, resp)
		return
	}

	resp := dto.UserAssignmentsStatsResponse{
		Stats: stats,
	}

	c.JSON(200, resp)
}
