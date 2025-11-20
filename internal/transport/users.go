package transport

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) setIsActive(c *gin.Context) {
	var req dto.SetUserIsActiveRequest

	if err := c.BindJSON(&req); err != nil {
		return
	}

	user, err := h.services.User.SetIsActive(req.UserID, req.IsActive)
	if err != nil {
		return
	}

	resp := dto.SetUserIsActiveResponse{
		User: user,
	}

	c.JSON(200, resp)
}

func (h *Handler) getReview(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		return
	}

	userPR, err := h.services.User.GetReview(userID)
	if err != nil {
		return
	}

	resp := dto.GetUsersReviewResponse{
		UserID:      userID,
		PullRequest: userPR,
	}

	c.JSON(200, resp)
}
