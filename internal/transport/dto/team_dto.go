package dto

import "github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"

// /team/add
type AddTeamDTO struct {
	Team *models.Team `json:"team"`
}

type TeamPRsResponse struct {
	TeamName    string                    `json:"team_name"`
	PullRequest []models.PullRequestShort `json:"pull_request"`
}
