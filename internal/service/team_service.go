package service

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
)

type TeamService struct {
	//
}

func NewTeamService() Team {
	return &TeamService{
		//
	}
}

// CreateTeam implements TeamInterface.
func (t *TeamService) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	panic("unimplemented")
}

// GetTeam implements TeamInterface.
func (t *TeamService) GetTeam(ctx context.Context, teamName string) (models.Team, error) {
	panic("unimplemented")
}
