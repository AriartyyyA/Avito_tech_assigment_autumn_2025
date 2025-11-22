package service

import (
	"context"

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
func (s *PullRequestService) CreatePullRequest(ctx context.Context, pullRequestID, pullRequestName, authorID string) (*models.PullRequest, error) {
	pullRequest := models.NewPullRequest(pullRequestID, pullRequestName, authorID)

	return s.repository.PullRequestRepository.CreatePullRequest(ctx, pullRequest)
}

// MergePullRequest implements PullRequestInterface.
func (s *PullRequestService) MergePullRequest(ctx context.Context, prID string) (*models.PullRequest, error) {
	return s.repository.PullRequestRepository.MergePullRequest(ctx, prID)
}

// ReassignPullRequest implements PullRequestInterface.
func (s *PullRequestService) ReassignPullRequest(ctx context.Context, prID string, oldReviewerID string) (*models.PullRequest, error) {
	return s.repository.PullRequestRepository.ReassignPullRequest(ctx, prID, oldReviewerID)
}
