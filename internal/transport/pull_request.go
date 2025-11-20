package transport

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPullRequest(c *gin.Context) {
	var req dto.CreatePRRequestDto

	if err := c.BindJSON(&req); err != nil {
		return
	}

	pullRequest, err := h.services.PullRequest.CreatePullRequest(req.PullRequestID, req.PullRequestName, req.AuthorID)
	if err != nil {
		return
	}

	resp := dto.PRResponseDto{
		PullRequest: pullRequest,
	}

	c.JSON(200, resp)
}

func (h *Handler) mergePullRequest(c *gin.Context) {
	var req dto.MergePRRequestDto

	if err := c.BindJSON(&req); err != nil {
		return
	}

	pullRequest, err := h.services.PullRequest.MergePullRequest(req.PullRequestID)
	if err != nil {
		return
	}

	resp := dto.PRResponseDto{
		PullRequest: pullRequest,
	}

	c.JSON(200, resp)
}

func (h *Handler) reassignPullRequest(c *gin.Context) {
	var req dto.ReassignPRRequestDto

	if err := c.BindJSON(&req); err != nil {
		return
	}

	pullRequest, err := h.services.PullRequest.ReassignPullRequest(req.PullRequestID, req.OldUserID)
	if err != nil {
		return
	}

	resp := dto.ReassignPRResponseDto{
		PullRequest: pullRequest,
		ReplacedBy:  pullRequest.AuthorID,
	}

	c.JSON(200, resp)
}
