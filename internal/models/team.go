package models

type Team struct {
	TeamName string       `json:"team_name"`
	Members  []TeamMember `json:"members"`
}

type TeamMember struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type TeamDeactivate struct {
	TeamName                string   `json:"team_name"`
	DeactivatedUsers        []string `json:"deactivated_users"`
	OpenPRCount             int      `json:"open_pr_count"`
	SuccessfulReassignments int      `json:"successful_reassignments"`
	FailedReassignments     int      `json:"failed_reassignments"`
}

func NewTeam(name string, members []TeamMember) *Team {
	return &Team{
		TeamName: name,
		Members:  members,
	}
}
