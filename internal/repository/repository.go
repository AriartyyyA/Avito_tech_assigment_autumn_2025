package repository

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User interface {
	SetIsActive(userID string, isActive bool) (models.User, error)
	GetReview(userID string) ([]models.PullRequestShort, error)
}

type PullRequest interface {
	CreatePullRequest(pr models.PullRequest) (models.PullRequest, error)
	MergePullRequest(prID string) (models.PullRequest, error)
	ReassignPullRequest(prID string, oldReviewerID string) (models.PullRequest, error)
}

type Team interface {
	AddTeam(ctx context.Context, team *models.Team) (*models.Team, error)
	GetTeam(ctx context.Context, teamName string) (*models.Team, error)
}

type Repository struct {
	User
	PullRequest
	Team
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		User:        NewUserRepository(db),
		PullRequest: NewPullRequestRepository(db),
		Team:        NewTeamRepository(db),
	}
}
