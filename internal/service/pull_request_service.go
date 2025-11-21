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

	pullRequest, err := s.repository.PullRequestRepository.CreatePullRequest(ctx, pullRequest)
	if err != nil {
		return nil, err
	}

	return pullRequest, nil
}

// MergePullRequest implements PullRequestInterface.
func (s *PullRequestService) MergePullRequest(ctx context.Context, prID string) (*models.PullRequest, error) {
	pullRequest, err := s.repository.PullRequestRepository.MergePullRequest(ctx, prID)
	if err != nil {
		return nil, err
	}

	return pullRequest, nil

}

// ReassignPullRequest implements PullRequestInterface.
func (s *PullRequestService) ReassignPullRequest(ctx context.Context, prID string, oldReviewerID string) (*models.PullRequest, error) {
	pullRequest, err := s.repository.PullRequestRepository.ReassignPullRequest(ctx, prID, oldReviewerID)
	if err != nil {
		return nil, err
	}

	return pullRequest, nil
}
