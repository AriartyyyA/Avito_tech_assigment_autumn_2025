package repository

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	SetIsActive(ctx context.Context, userID string, isActive bool) (*models.User, error)
	GetReview(ctx context.Context, userID string) ([]models.PullRequestShort, error)
	GetAssignmentsStats(ctx context.Context) ([]models.UserAssignmentsStat, error)
}

type PullRequestRepository interface {
	CreatePullRequest(ctx context.Context, pr *models.PullRequest) (*models.PullRequest, error)
	MergePullRequest(ctx context.Context, prID string) (*models.PullRequest, error)
	ReassignPullRequest(ctx context.Context, prID string, oldReviewerID string) (*models.PullRequest, error)
}

type Team interface {
	AddTeam(ctx context.Context, team *models.Team) (*models.Team, error)
	GetTeam(ctx context.Context, teamName string) (*models.Team, error)
	GetTeamPullRequests(ctx context.Context, teamName string) ([]models.PullRequestShort, error)
}

type Repository struct {
	UserRepository
	PullRequestRepository
	Team
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		UserRepository:        NewUserRepository(db),
		PullRequestRepository: NewPullRequestRepository(db),
		Team:                  NewTeamRepository(db),
	}
}
