package service

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
)

type User interface {
	SetIsActive(ctx context.Context, userID string, isActive bool) (models.User, error)
	GetReview(ctx context.Context, userID string) ([]models.PullRequestShort, error)
}

type PullRequest interface {
	CreatePullRequest(ctx context.Context, pr models.PullRequest) (models.PullRequest, error)
	MergePullRequest(ctx context.Context, prID string) (models.PullRequest, error)
	ReassignPullRequest(ctx context.Context, prID string, oldReviewerID string) (models.PullRequest, error)
}

type Team interface {
	CreateTeam(ctx context.Context, team models.Team) (models.Team, error)
	GetTeam(ctx context.Context, teamName string) (models.Team, error)
}

type Service struct {
	User
	PullRequest
	Team
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:        NewUserService(),
		PullRequest: NewPullRequestService(),
		Team:        NewTeamService(),
	}
}
