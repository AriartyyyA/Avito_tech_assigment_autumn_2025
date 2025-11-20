package service

import (
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

func (s *UserService) SetIsActive(userID string, isActive bool) (models.User, error) {
	user, err := s.repository.User.SetIsActive(userID, isActive)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *UserService) GetReview(userID string) ([]models.PullRequestShort, error) {
	userPR, err := s.repository.User.GetReview(userID)
	if err != nil {
		return nil, err
	}

	return userPR, nil
}
