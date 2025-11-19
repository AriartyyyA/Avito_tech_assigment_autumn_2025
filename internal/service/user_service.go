package service

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
)

type UserService struct {
	// ...
}

func NewUserService() User {
	return &UserService{}
}

func (s *UserService) SetIsActive(ctx context.Context, userID string, isActive bool) (models.User, error) {
	// ...
	return models.User{}, nil
}

func (s *UserService) GetReview(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
	// ...
	return nil, nil
}
