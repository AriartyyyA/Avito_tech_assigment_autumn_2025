package service

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
)

type UserService struct {
	// ...
}

func NewUserService() User {
	return &UserService{}
}

func (s *UserService) SetIsActive(userID string, isActive bool) (models.User, error) {
	// ...
	return models.User{}, nil
}

func (s *UserService) GetReview(userID string) ([]models.PullRequestShort, error) {
	// ...
	return nil, nil
}
