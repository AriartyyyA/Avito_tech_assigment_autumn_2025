package dto

import "github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"

// /pull_request/create
type CreatePRRequestDto struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
}

// Эта DTO у нас используется и для create и для merge
type PRResponseDto struct {
	PullRequest models.PullRequest `json:"pull_request"`
}

// /pull_request/merge
type MergePRRequestDto struct {
	PullRequestID string `json:"pull_request_id"`
}

// /pull_request/reassign
type ReassignPRRequestDto struct {
	PullRequestID string `json:"pull_request_id"`
	OldUserID     string `json:"old_user_id"`
}

type ReassignPRResponseDto struct {
	PullRequest models.PullRequest `json:"pull_request"`
	ReplacedBy  string             `json:"replaced_by"`
}
