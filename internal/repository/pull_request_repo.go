package repository

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
)

type PullRequestRepository struct {
	//
}

func NewPullRequestRepository() PullRequest {
	return &PullRequestRepository{
		//
	}
}

func (p *PullRequestRepository) CreatePullRequest(ctx context.Context, pr models.PullRequest) (models.PullRequest, error) {
	panic("unimplemented")
}

func (p *PullRequestRepository) MergePullRequest(ctx context.Context, prID string) (models.PullRequest, error) {
	panic("unimplemented")
}

func (p *PullRequestRepository) ReassignPullRequest(ctx context.Context, prID string, oldReviewerID string) (models.PullRequest, error) {
	panic("unimplemented")
}
