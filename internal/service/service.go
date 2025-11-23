package service

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
)

type UserService interface {
	SetIsActive(ctx context.Context, userID string, isActive bool) (*models.User, error)
	GetReview(ctx context.Context, userID string) ([]models.PullRequestShort, error)
	GetUserAssignmentsStats(ctx context.Context) ([]models.UserAssignmentsStat, error)
}

type PullRequestService interface {
	CreatePullRequest(ctx context.Context, pullRequestID, pullRequestName, authorID string) (*models.PullRequest, error)
	MergePullRequest(ctx context.Context, prID string) (*models.PullRequest, error)
	ReassignPullRequest(ctx context.Context, prID string, OldUserID string) (*models.PullRequest, error)
}

type TeamService interface {
	AddTeam(ctx context.Context, team *models.Team) (*models.Team, error)
	GetTeam(ctx context.Context, teamName string) (*models.Team, error)
	GetTeamPullRequests(ctx context.Context, teamName string) ([]models.PullRequestShort, error)
	DeactivateTeam(ctx context.Context, teamName string) (*models.TeamDeactivate, error)
}

type Service struct {
	UserService
	PullRequestService
	TeamService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService:        NewUserService(repo),
		PullRequestService: NewPullRequestService(repo),
		TeamService:        NewTeamService(repo),
	}
}
