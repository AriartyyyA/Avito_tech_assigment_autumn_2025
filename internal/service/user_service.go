package service

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
)

type UserService struct {
	repository *repository.Repository
}

func NewUserService(repository *repository.Repository) User {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) SetIsActive(ctx context.Context, userID string, isActive bool) (*models.User, error) {
	user, err := s.repository.UserRepository.SetIsActive(ctx, userID, isActive)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetReview(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
	userPR, err := s.repository.UserRepository.GetReview(ctx, userID)
	if err != nil {
		return nil, err
	}

	return userPR, nil
}

func (s *UserService) GetUserAssignmentsStats(ctx context.Context) ([]models.UserAssignmentsStat, error) {
	stats, err := s.repository.UserRepository.GetAssignmentsStats(ctx)
	if err != nil {
		return nil, err
	}

	return stats, nil
}
