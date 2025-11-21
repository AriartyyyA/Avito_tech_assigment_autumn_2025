package dto

import "github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"

// /users/set_is_active
type SetUserIsActiveRequest struct {
	UserID   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}

type SetUserIsActiveResponse struct {
	User *models.User `json:"user"`
}

// /users/get_review
type GetUsersReviewResponse struct {
	UserID      string                    `json:"user_id"`
	PullRequest []models.PullRequestShort `json:"pull_request"`
}
