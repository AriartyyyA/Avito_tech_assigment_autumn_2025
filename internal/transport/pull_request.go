package transport

import (
	"errors"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPullRequest(c *gin.Context) {
	var req dto.CreatePRRequestDto

	if err := c.BindJSON(&req); err != nil {
		err := InvalidRequest("")
		c.JSON(400, err)
		return
	}

	if req.PullRequestID == "" {
		err := InvalidRequest("pull_request_id")
		c.JSON(400, err)
		return
	}

	if req.PullRequestName == "" {
		err := InvalidRequest("pull_request_name")
		c.JSON(400, err)
		return
	}

	if req.AuthorID == "" {
		err := InvalidRequest("author_id")
		c.JSON(400, err)
		return
	}

	pullRequest, err := h.services.PullRequest.CreatePullRequest(c.Request.Context(), req.PullRequestID, req.PullRequestName, req.AuthorID)
	switch {
	// case errors.Is(err, models.ErrorCodeTeamNotFound):
	// 	err := NotFound("Team")
	// 	c.JSON(404, err)
	// 	return
	case errors.Is(err, models.ErrorCodeUserNotFound):
		err := NotFound()
		c.JSON(404, err)
		return
	case errors.Is(err, models.ErrorCodePRExists):
		err := PRExists()
		c.JSON(409, err)
		return
	}

	resp := dto.PRResponseDto{
		PullRequest: *pullRequest,
	}

	c.JSON(201, resp)
}

func (h *Handler) mergePullRequest(c *gin.Context) {
	var req dto.MergePRRequestDto

	if err := c.BindJSON(&req); err != nil {
		err := InvalidRequest("")
		c.JSON(400, err)
		return
	}

	if req.PullRequestID == "" {
		err := InvalidRequest("pull_request_id")
		c.JSON(400, err)
		return
	}

	pullRequest, err := h.services.PullRequest.MergePullRequest(c.Request.Context(), req.PullRequestID)
	if err != nil {
		err := NotFound()
		c.JSON(404, err)
		return
	}

	resp := dto.MergePRResponseDto{
		PullRequest: *pullRequest,
		MergedAt:    *pullRequest.MergedAt,
	}

	c.JSON(200, resp)
}

func (h *Handler) reassignPullRequest(c *gin.Context) {
	var req dto.ReassignPRRequestDto

	if err := c.BindJSON(&req); err != nil {
		err := InvalidRequest("")
		c.JSON(400, err)
		return
	}

	if req.PullRequestID == "" {
		err := InvalidRequest("pull_request_id")
		c.JSON(400, err)
		return
	}

	if req.OldUserID == "" {
		err := InvalidRequest("old_user_id")
		c.JSON(400, err)
		return
	}

	finishPR, err := h.services.PullRequest.ReassignPullRequest(c.Request.Context(), req.PullRequestID, req.OldUserID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrorCodePRNotFound):
			err := NotFound()
			c.JSON(404, err)
			return
		case errors.Is(err, models.ErrorCodeUserNotFound):
			err := NotFound()
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
			err := InvalidRequest("")
			c.JSON(500, err)
		}
		return
	}

	resp := dto.ReassignPRResponseDto{
		PullRequest: finishPR,
		ReplacedBy:  finishPR.NewReviewerID,
	}

	c.JSON(200, resp)
}
