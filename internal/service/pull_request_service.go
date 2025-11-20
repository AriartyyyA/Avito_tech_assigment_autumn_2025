package service

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
)

type PullRequestService struct {
	repository *repository.Repository
}

func NewPullRequestService(repository *repository.Repository) PullRequest {
	return &PullRequestService{
		repository: repository,
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
