package service

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
)

type pullRequestService struct {
	repository *repository.Repository
}

func NewPullRequestService(repository *repository.Repository) PullRequestService {
	return &pullRequestService{
		repository: repository,
	}
}

// CreatePullRequest implements PullRequestInterface.

func (s *pullRequestService) CreatePullRequest(ctx context.Context, pullRequestID, pullRequestName, authorID string) (*models.PullRequest, error) {
	pullRequest := models.NewPullRequest(pullRequestID, pullRequestName, authorID)
	return s.repository.PullRequestRepository.CreatePullRequest(ctx, pullRequest)
}

// MergePullRequest implements PullRequestInterface.

func (s *pullRequestService) MergePullRequest(ctx context.Context, prID string) (*models.PullRequest, error) {
	return s.repository.PullRequestRepository.MergePullRequest(ctx, prID)
}

// ReassignPullRequest implements PullRequestInterface.

func (s *pullRequestService) ReassignPullRequest(ctx context.Context, prID string, OldUserID string) (*models.PullRequest, error) {
	return s.repository.PullRequestRepository.ReassignPullRequest(ctx, prID, OldUserID)
}
