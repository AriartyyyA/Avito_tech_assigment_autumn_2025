package service

import (
	"context"
	"log"

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
	pr, err := s.repository.PullRequestRepository.CreatePullRequest(ctx, pullRequest)
	if err != nil {
		log.Printf("ERROR: Failed to create PR in repository: PR=%s, Author=%s, Error=%v", pullRequestID, authorID, err)
		return nil, err
	}
	return pr, nil
}

// MergePullRequest implements PullRequestInterface.

func (s *pullRequestService) MergePullRequest(ctx context.Context, prID string) (*models.PullRequest, error) {
	pr, err := s.repository.PullRequestRepository.MergePullRequest(ctx, prID)
	if err != nil {
		log.Printf("ERROR: Failed to merge PR in repository: PR=%s, Error=%v", prID, err)
		return nil, err
	}
	return pr, nil
}

// ReassignPullRequest implements PullRequestInterface.

func (s *pullRequestService) ReassignPullRequest(ctx context.Context, prID string, OldUserID string) (*models.PullRequest, error) {
	pr, err := s.repository.PullRequestRepository.ReassignPullRequest(ctx, prID, OldUserID)
	if err != nil {
		log.Printf("ERROR: Failed to reassign PR in repository: PR=%s, OldReviewer=%s, Error=%v", prID, OldUserID, err)
		return nil, err
	}
	return pr, nil
}
