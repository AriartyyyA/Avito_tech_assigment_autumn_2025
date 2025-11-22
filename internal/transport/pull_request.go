package transport

import (
	"errors"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPullRequest(c *gin.Context) {
	req, exists := c.Get("validated_request")
	if !exists {
		err := InternalError()
		c.JSON(500, err)
		return
	}

	createReq := req.(*dto.CreatePRRequestDto)
	pullRequest, err := h.services.PullRequest.CreatePullRequest(c.Request.Context(), createReq.PullRequestID, createReq.PullRequestName, createReq.AuthorID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrorCodeUserNotFound):
			err := NotFound(models.ErrorCodeUserNotFound)
			c.JSON(404, err)
			return
		case errors.Is(err, models.ErrorCodePRExists):
			err := PRExists()
			c.JSON(409, err)
			return
		default:
			err := InternalError()
			c.JSON(500, err)
			return
		}
	}

	resp := dto.PRResponseDto{
		PullRequest: *pullRequest,
	}

	c.JSON(201, resp)
}

func (h *Handler) mergePullRequest(c *gin.Context) {
	req, exists := c.Get("validated_request")
	if !exists {
		err := InternalError()
		c.JSON(500, err)
		return
	}

	mergeReq := req.(*dto.MergePRRequestDto)
	pullRequest, err := h.services.PullRequest.MergePullRequest(c.Request.Context(), mergeReq.PullRequestID)
	if err != nil {
		if errors.Is(err, models.ErrorCodePRNotFound) {
			err := NotFound(models.ErrorCodePRNotFound)
			c.JSON(404, err)
			return
		}

		err := InternalError()
		c.JSON(500, err)
		return
	}

	resp := dto.MergePRResponseDto{
		PullRequest: *pullRequest,
		MergedAt:    *pullRequest.MergedAt,
	}

	c.JSON(200, resp)
}

func (h *Handler) reassignPullRequest(c *gin.Context) {
	req, exists := c.Get("validated_request")
	if !exists {
		err := InternalError()
		c.JSON(500, err)
		return
	}

	reassignReq := req.(*dto.ReassignPRRequestDto)
	finishPR, err := h.services.PullRequest.ReassignPullRequest(c.Request.Context(), reassignReq.PullRequestID, reassignReq.OldUserID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrorCodePRNotFound):
			err := NotFound(models.ErrorCodePRNotFound)
			c.JSON(404, err)
			return
		case errors.Is(err, models.ErrorCodeUserNotFound):
			err := NotFound(models.ErrorCodeUserNotFound)
			c.JSON(404, err)
			return
		case errors.Is(err, models.ErrorCodePRMerged):
			err := PRMerged()
			c.JSON(409, err)
			return
		case errors.Is(err, models.ErrorCodeNotAssigned):
			err := NotAssigned()
			c.JSON(409, err)
			return
		case errors.Is(err, models.ErrorCodeNoCandidate):
			err := NoCandidate()
			c.JSON(409, err)
			return
		default:
			err := InternalError()
			c.JSON(500, err)
			return
		}
	}

	resp := dto.ReassignPRResponseDto{
		PullRequest: finishPR,
		ReplacedBy:  finishPR.NewReviewerID,
	}

	c.JSON(200, resp)
}
