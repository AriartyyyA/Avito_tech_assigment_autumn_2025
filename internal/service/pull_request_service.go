package service

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
)

type PullRequestService struct {
	// ...
}

func NewPullRequestService() PullRequest {
	return &PullRequestService{
		// ...
	}
}

// CreatePullRequest implements PullRequestInterface.
func (p *PullRequestService) CreatePullRequest(pullRequestID, pullRequestName, authorID string) (models.PullRequest, error) {
	panic("unimplemented")
}

// MergePullRequest implements PullRequestInterface.
func (p *PullRequestService) MergePullRequest(prID string) (models.PullRequest, error) {
	panic("unimplemented")
}

// ReassignPullRequest implements PullRequestInterface.
func (p *PullRequestService) ReassignPullRequest(prID string, oldReviewerID string) (models.PullRequest, error) {
	panic("unimplemented")
}
