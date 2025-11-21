package models

import "time"

type PullRequestStatus string

const (
	PullRequestStatusOpen   PullRequestStatus = "OPEN"
	PullRequestStatusMerged PullRequestStatus = "MERGED"
)

type PullRequest struct {
	PullRequestID     string            `json:"pull_request_id"`
	PullRequestName   string            `json:"pull_request_name"`
	AuthorID          string            `json:"author_id"`
	Status            PullRequestStatus `json:"status"`
	AssignedReviewers []string          `json:"assigned_reviewers"`

	MergedAt      *time.Time `json:"-"`
	NewReviewerID string     `json:"-"`
}

type PullRequestShort struct {
	PullRequestID   string            `json:"pull_request_id"`
	PullRequestName string            `json:"pull_request_name"`
	AuthorID        string            `json:"author_id"`
	Status          PullRequestStatus `json:"status"`
}

func NewPullRequestShort(
	id string,
	name string,
	authorID string,
	status string,
) *PullRequestShort {
	return &PullRequestShort{
		PullRequestID:   id,
		PullRequestName: name,
		AuthorID:        authorID,
		Status:          PullRequestStatus(status),
	}
}

func NewPullRequest(
	id string,
	name string,
	authorID string,
) *PullRequest {
	return &PullRequest{
		PullRequestID:   id,
		PullRequestName: name,
		AuthorID:        authorID,
		Status:          PullRequestStatusOpen,
	}
}
