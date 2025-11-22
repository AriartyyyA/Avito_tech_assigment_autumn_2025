package models

type UserAssignmentsStat struct {
	UserID                 string `json:"user_id"`
	Username               string `json:"username"`
	ReviewAssignmentsCount int    `json:"review_assignments_count"`
}
