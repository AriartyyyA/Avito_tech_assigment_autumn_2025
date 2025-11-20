package service

import (
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
func (t *TeamService) AddTeam(team models.Team) (models.Team, error) {
	team, err := t.repository.Team.AddTeam(team)
	if err != nil {
		return models.Team{}, err
	}

	return team, nil
}

// GetTeam implements TeamInterface.
func (t *TeamService) GetTeam(teamName string) (models.Team, error) {
	team, err := t.repository.Team.GetTeam(teamName)
	if err != nil {
		return models.Team{}, err
	}

	return team, nil
}
