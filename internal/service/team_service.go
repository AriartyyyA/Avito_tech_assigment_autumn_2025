package service

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
)

type TeamService struct {
	repository *repository.Repository
}

func NewTeamService(repository *repository.Repository) Team {
	return &TeamService{
		repository: repository,
	}
}

// CreateTeam implements TeamInterface.
func (s *TeamService) AddTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	createdTeam, err := s.repository.Team.AddTeam(ctx, team)
	if err != nil {
		return nil, err
	}

	return createdTeam, nil
}

// GetTeam implements TeamInterface.
func (s *TeamService) GetTeam(ctx context.Context, teamName string) (*models.Team, error) {
	team, err := s.repository.Team.GetTeam(ctx, teamName)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (s *TeamService) GetTeamPullRequests(ctx context.Context, teamName string) ([]models.PullRequestShort, error) {
	if _, err := s.repository.Team.GetTeam(ctx, teamName); err != nil {
		return nil, err
	}

	return s.repository.Team.GetTeamPullRequests(ctx, teamName)
}
