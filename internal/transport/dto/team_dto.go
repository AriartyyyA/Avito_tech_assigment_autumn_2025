package dto

import "github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"

// /team/add
type AddTeamDTO struct {
	Team *models.Team `json:"team"`
}

type TeamPRsResponse struct {
	TeamName    string                    `json:"team_name"`
	PullRequest []models.PullRequestShort `json:"pull_request"`
	PRcount     int                       `json:"pr_count"`
}

type DeactivateTeamUsersRequest struct {
	TeamName string `json:"team_name"`
}

type DeactivateTeamUsersResponse struct {
	TeamName                string   `json:"team_name"`
	DeactivatedUsers        []string `json:"deactivated_users"`
	OpenPRCount             int      `json:"open_pr_count"`
	SuccessfulReassignments int      `json:"successful_reassignments"`
	FailedReassignments     int      `json:"failed_reassignments"`
}
