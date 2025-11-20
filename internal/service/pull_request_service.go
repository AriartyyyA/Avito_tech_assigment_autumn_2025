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
func (s *PullRequestService) CreatePullRequest(pullRequestID, pullRequestName, authorID string) (models.PullRequest, error) {
	pullRequest := models.NewPullRequest(pullRequestID, pullRequestName, authorID)

	pullRequest, err := s.repository.PullRequest.CreatePullRequest(pullRequest)
	if err != nil {
		return models.PullRequest{}, err
	}

	return pullRequest, nil
}

// MergePullRequest implements PullRequestInterface.
func (s *PullRequestService) MergePullRequest(prID string) (models.PullRequest, error) {
	pullRequest, err := s.repository.PullRequest.MergePullRequest(prID)
	if err != nil {
		return models.PullRequest{}, err
	}

	return pullRequest, nil

}

// ReassignPullRequest implements PullRequestInterface.
func (s *PullRequestService) ReassignPullRequest(prID string, oldReviewerID string) (models.PullRequest, error) {
	pullRequest, err := s.repository.PullRequest.ReassignPullRequest(prID, oldReviewerID)
	if err != nil {
		return models.PullRequest{}, err
	}

	return pullRequest, nil
}
