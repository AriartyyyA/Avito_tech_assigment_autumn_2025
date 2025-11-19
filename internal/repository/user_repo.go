package repository

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
)

type UserRepository struct {
	//
}

func NewUserRepository() User {
	return &UserRepository{
		//
	}
}

func (u *UserRepository) GetReview(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
	panic("unimplemented")
}

func (u *UserRepository) SetIsActive(ctx context.Context, userID string, isActive bool) (models.User, error) {
	panic("unimplemented")
}
