package transport

import (
	"errors"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) setIsActive(c *gin.Context) {
	var req dto.SetUserIsActiveRequest

	if err := c.BindJSON(&req); err != nil {
		err := InvalidRequest("")
		c.JSON(400, err)
		return
	}

	if req.UserID == "" {
		err := InvalidRequest("user_id")
		c.JSON(400, err)
		return
	}

	user, err := h.services.User.SetIsActive(c.Request.Context(), req.UserID, req.IsActive)
	if err != nil {
		if errors.Is(err, models.ErrorCodeUserNotFound) {
			err := NotFound()
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

	if userID == "" {
		err := InvalidRequest("user_id")
		c.JSON(400, err)
		return
	}

	userPR, err := h.services.User.GetReview(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, models.ErrorCodeUserNotFound) {
			err := NotFound()
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
