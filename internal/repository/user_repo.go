package repository

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) User {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetReview(userID string) ([]models.PullRequestShort, error) {
	panic("unimplemented")
}

func (u *UserRepository) SetIsActive(userID string, isActive bool) (models.User, error) {
	panic("unimplemented")
}
