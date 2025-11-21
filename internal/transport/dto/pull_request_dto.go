package dto

import (
	"time"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
)

// /pull_request/create
type CreatePRRequestDto struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
}

// Эта DTO у нас используется и для create и для merge
type PRResponseDto struct {
	PullRequest models.PullRequest `json:"pr"`
}

// /pull_request/merge
type MergePRRequestDto struct {
	PullRequestID string `json:"pull_request_id"`
}

type MergePRResponseDto struct {
	PullRequest models.PullRequest `json:"pr"`
	MergedAt    time.Time          `json:"merged_at"`
}

// /pull_request/reassign
type ReassignPRRequestDto struct {
	PullRequestID string `json:"pull_request_id"`
	OldUserID     string `json:"old_reviewer_id"`
}

type ReassignPRResponseDto struct {
	PullRequest *models.PullRequest `json:"pull_request"`
	ReplacedBy  string              `json:"replaced_by"`
}
