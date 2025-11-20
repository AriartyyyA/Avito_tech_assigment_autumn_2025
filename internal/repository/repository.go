package repository

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
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
	AddTeam(team models.Team) (models.Team, error)
	GetTeam(teamName string) (models.Team, error)
}

type Repository struct {
	User
	PullRequest
	Team
}

func NewRepository() *Repository {
	return &Repository{
		User:        NewUserRepository(),
		PullRequest: NewPullRequestRepository(),
		Team:        NewTeamRepository(),
	}
}
