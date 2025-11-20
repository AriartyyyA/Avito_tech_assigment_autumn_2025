package service

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
)

type User interface {
	SetIsActive(userID string, isActive bool) (models.User, error)
	GetReview(userID string) ([]models.PullRequestShort, error)
}

type PullRequest interface {
	CreatePullRequest(pullRequestID, pullRequestName, authorID string) (models.PullRequest, error)
	MergePullRequest(prID string) (models.PullRequest, error)
	ReassignPullRequest(prID string, oldReviewerID string) (models.PullRequest, error)
}

type Team interface {
	AddTeam(ctx context.Context, team *models.Team) (*models.Team, error)
	GetTeam(ctx context.Context, teamName string) (*models.Team, error)
}

type Service struct {
	User
	PullRequest
	Team
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:        NewUserService(repo),
		PullRequest: NewPullRequestService(repo),
		Team:        NewTeamService(repo),
	}
}
