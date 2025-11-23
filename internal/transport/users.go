package transport

import (
	"errors"
	"log"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) setIsActive(c *gin.Context) {
	req, exists := c.Get("validated_request")

	if !exists {
		log.Printf("ERROR: validated_request not found in context for setIsActive")
		err := InternalError()
		c.JSON(500, err)
		return

	}

	setIsActiveReq := req.(*dto.SetUserIsActiveRequest)
	user, err := h.services.UserService.SetIsActive(c.Request.Context(), setIsActiveReq.UserID, setIsActiveReq.IsActive)
	if err != nil {
		if errors.Is(err, models.ErrorCodeUserNotFound) {
			log.Printf("WARN: User not found for setIsActive: User=%s, IsActive=%v", setIsActiveReq.UserID, setIsActiveReq.IsActive)
			err := NotFound(models.ErrorCodeUserNotFound)
			c.JSON(404, err)
			return
		}

		log.Printf("ERROR: Failed to set user active status: User=%s, IsActive=%v, Error=%v", setIsActiveReq.UserID, setIsActiveReq.IsActive, err)
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
	userPR, err := h.services.UserService.GetReview(c.Request.Context(), userID)

	if err != nil {
		if errors.Is(err, models.ErrorCodeUserNotFound) {
			log.Printf("WARN: User not found for getReview: User=%s", userID)
			err := NotFound(models.ErrorCodeUserNotFound)
			c.JSON(404, err)
			return
		}

		log.Printf("ERROR: Failed to get user review: User=%s, Error=%v", userID, err)
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
	stats, err := h.services.UserService.GetUserAssignmentsStats(c.Request.Context())

	if err != nil {
		log.Printf("ERROR: Failed to get user assignments stats: Error=%v", err)
		resp := InternalError()
		c.JSON(500, resp)
		return

	}
	resp := dto.UserAssignmentsStatsResponse{
		Stats: stats,
	}
	c.JSON(200, resp)

}
