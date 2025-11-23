package transport

import (
	"errors"
	"log"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPullRequest(c *gin.Context) {
	req, exists := c.Get("validated_request")

	if !exists {
		log.Printf("ERROR: validated_request not found in context for createPullRequest")
		err := InternalError()
		c.JSON(500, err)
		return
	}

	createReq := req.(*dto.CreatePRRequestDto)
	pullRequest, err := h.services.PullRequestService.CreatePullRequest(c.Request.Context(), createReq.PullRequestID, createReq.PullRequestName, createReq.AuthorID)

	if err != nil {
		switch {
		case errors.Is(err, models.ErrorCodeUserNotFound):
			log.Printf("WARN: User not found for PR creation: PR=%s, Author=%s", createReq.PullRequestID, createReq.AuthorID)
			err := NotFound(models.ErrorCodeUserNotFound)
			c.JSON(404, err)
			return

		case errors.Is(err, models.ErrorCodePRExists):
			log.Printf("WARN: PR already exists: PR=%s", createReq.PullRequestID)
			err := PRExists()
			c.JSON(409, err)
			return

		default:
			log.Printf("ERROR: Failed to create PR: PR=%s, Author=%s, Error=%v", createReq.PullRequestID, createReq.AuthorID, err)
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
		log.Printf("ERROR: validated_request not found in context for mergePullRequest")
		err := InternalError()
		c.JSON(500, err)
		return
	}

	mergeReq := req.(*dto.MergePRRequestDto)
	pullRequest, err := h.services.PullRequestService.MergePullRequest(c.Request.Context(), mergeReq.PullRequestID)

	if err != nil {
		if errors.Is(err, models.ErrorCodePRNotFound) {
			log.Printf("WARN: PR not found for merge: PR=%s", mergeReq.PullRequestID)
			err := NotFound(models.ErrorCodePRNotFound)
			c.JSON(404, err)
			return

		}

		log.Printf("ERROR: Failed to merge PR: PR=%s, Error=%v", mergeReq.PullRequestID, err)
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
		log.Printf("ERROR: validated_request not found in context for reassignPullRequest")
		err := InternalError()
		c.JSON(500, err)
		return
	}

	reassignReq := req.(*dto.ReassignPRRequestDto)

	finishPR, err := h.services.PullRequestService.ReassignPullRequest(c.Request.Context(), reassignReq.PullRequestID, reassignReq.OldUserID)

	if err != nil {
		switch {
		case errors.Is(err, models.ErrorCodePRNotFound):
			log.Printf("WARN: PR not found for reassign: PR=%s, OldReviewer=%s", reassignReq.PullRequestID, reassignReq.OldUserID)
			err := NotFound(models.ErrorCodePRNotFound)
			c.JSON(404, err)
			return

		case errors.Is(err, models.ErrorCodeUserNotFound):
			log.Printf("WARN: User not found for reassign: PR=%s, OldReviewer=%s", reassignReq.PullRequestID, reassignReq.OldUserID)
			err := NotFound(models.ErrorCodeUserNotFound)
			c.JSON(404, err)
			return

		case errors.Is(err, models.ErrorCodePRMerged):
			log.Printf("WARN: Attempt to reassign merged PR: PR=%s, OldReviewer=%s", reassignReq.PullRequestID, reassignReq.OldUserID)
			err := PRMerged()
			c.JSON(409, err)
			return

		case errors.Is(err, models.ErrorCodeNotAssigned):
			log.Printf("WARN: Reviewer not assigned to PR: PR=%s, OldReviewer=%s", reassignReq.PullRequestID, reassignReq.OldUserID)
			err := NotAssigned()
			c.JSON(409, err)
			return

		case errors.Is(err, models.ErrorCodeNoCandidate):
			log.Printf("WARN: No candidate found for reassign: PR=%s, OldReviewer=%s", reassignReq.PullRequestID, reassignReq.OldUserID)
			err := NoCandidate()
			c.JSON(409, err)
			return

		default:
			log.Printf("ERROR: Failed to reassign PR: PR=%s, OldReviewer=%s, Error=%v", reassignReq.PullRequestID, reassignReq.OldUserID, err)
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
