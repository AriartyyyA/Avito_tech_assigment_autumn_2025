package repository

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
)

type TeamRepository struct {
	//
}

func NewTeamRepository() Team {
	return &TeamRepository{
		//
	}
}

func (t *TeamRepository) AddTeam(team models.Team) (models.Team, error) {
	panic("unimplemented")
}

func (t *TeamRepository) GetTeam(teamName string) (models.Team, error) {
	panic("unimplemented")
}
