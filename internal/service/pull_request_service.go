package service

import (
	"context"

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
func (p *PullRequestService) CreatePullRequest(ctx context.Context, pr models.PullRequest) (models.PullRequest, error) {
	panic("unimplemented")
}

// MergePullRequest implements PullRequestInterface.
func (p *PullRequestService) MergePullRequest(ctx context.Context, prID string) (models.PullRequest, error) {
	panic("unimplemented")
}

// ReassignPullRequest implements PullRequestInterface.
func (p *PullRequestService) ReassignPullRequest(ctx context.Context, prID string, oldReviewerID string) (models.PullRequest, error) {
	panic("unimplemented")
}
