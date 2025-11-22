package repository

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pullRequestRepository struct {
	db *pgxpool.Pool
}

func NewPullRequestRepository(db *pgxpool.Pool) PullRequestRepository {
	return &pullRequestRepository{
		db: db,
	}
}

func (r *pullRequestRepository) CreatePullRequest(ctx context.Context, pr *models.PullRequest) (*models.PullRequest, error) {
	if pr == nil {
		return nil, fmt.Errorf("pull request is nil")
	}
	if pr.PullRequestID == "" || pr.PullRequestName == "" || pr.AuthorID == "" {
		return nil, fmt.Errorf("id, name and author_id are required")
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx (create PR): %w", err)
	}
	defer tx.Rollback(ctx)

	const authorQuery = `SELECT team_name
	FROM users
	WHERE user_id = $1`

	var authorTeam string

	if err = tx.QueryRow(ctx, authorQuery, pr.AuthorID).Scan(&authorTeam); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrorCodeUserNotFound
		}

		return nil, fmt.Errorf("select author team: %w", err)
	}

	const candidatesQuery = `SELECT user_id
	FROM users
	WHERE team_name = $1
		AND is_active = TRUE
		AND user_id <> $2`

	candidatesRows, err := tx.Query(ctx, candidatesQuery, authorTeam, pr.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("select reviewer candidates: %w", err)
	}
	defer candidatesRows.Close()

	candidatesID := make([]string, 0)

	for candidatesRows.Next() {
		var uID string
		if err := candidatesRows.Scan(&uID); err != nil {
			return nil, fmt.Errorf("scan candidates: %w", err)
		}
		candidatesID = append(candidatesID, uID)
	}

	if err := candidatesRows.Err(); err != nil {
		return nil, fmt.Errorf("iterate candidates: %w", err)
	}

	selectedReviewers := chooseRandom(candidatesID, 2)

	const insertPR = `
INSERT INTO pull_requests (pull_request_id, pull_request_name, author_id, status)
VALUES ($1, $2, $3, $4)
`
	if _, err := tx.Exec(ctx, insertPR, pr.PullRequestID, pr.PullRequestName, pr.AuthorID, string(models.PullRequestStatusOpen)); err != nil {
		if isUnique(err) {
			return nil, models.ErrorCodePRExists
		}

		return nil, fmt.Errorf("insert pull_request: %w", err)
	}

	if len(selectedReviewers) > 0 {
		const insertReviewer = `
INSERT INTO pull_request_reviewers (pull_request_id, user_id)
VALUES ($1, $2)
`
		for _, reviewerID := range selectedReviewers {
			if _, err := tx.Exec(ctx, insertReviewer, pr.PullRequestID, reviewerID); err != nil {
				return nil, fmt.Errorf("insert pull_request_reviewer (%s): %w", reviewerID, err)
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit tx (created PR): %w", err)
	}

	createdPR := &models.PullRequest{
		PullRequestID:     pr.PullRequestID,
		PullRequestName:   pr.PullRequestName,
		AuthorID:          pr.AuthorID,
		Status:            models.PullRequestStatusOpen,
		AssignedReviewers: selectedReviewers,
	}

	return createdPR, nil
}

func (r *pullRequestRepository) ReassignPullRequest(ctx context.Context, prID string, oldReviewerID string) (*models.PullRequest, error) {
	if prID == "" || oldReviewerID == "" {
		return nil, fmt.Errorf("prID and oldReviewerID are Required")
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx (reassign PR): %w", err)
	}
	defer tx.Rollback(ctx)

	const selectPRQuery = `
SELECT pull_request_id, pull_request_name, author_id, status
FROM pull_requests
WHERE pull_request_id = $1
FOR UPDATE
`
	var (
		dbPRID   string
		name     string
		authorID string
		status   string
	)

	if err := tx.QueryRow(ctx, selectPRQuery, prID).
		Scan(&dbPRID, &name, &authorID, &status); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrorCodePRNotFound
		}

		return nil, fmt.Errorf("select pull_request for reassign: %w", err)
	}

	if status == string(models.PullRequestStatusMerged) {
		return nil, models.ErrorCodePRMerged
	}

	const checkReviewerQuery = `
SELECT user_id
FROM pull_request_reviewers
WHERE pull_request_id = $1 AND user_id = $2
`
	var assignedID string

	if err := tx.QueryRow(ctx, checkReviewerQuery, prID, oldReviewerID).
		Scan(&assignedID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrorCodeNotAssigned
		}

		return nil, fmt.Errorf("check reviewer assigned: %w", err)
	}

	const currentReviewersQuery = `
SELECT user_id
FROM pull_request_reviewers
WHERE pull_request_id = $1
`

	rows, err := tx.Query(ctx, currentReviewersQuery, prID)
	if err != nil {
		return nil, fmt.Errorf("select current reviewers: %w", err)
	}
	defer rows.Close()

	var (
		currentReviewers []string
		otherReviewersID string
	)

	for rows.Next() {
		var uID string
		if err := rows.Scan(&uID); err != nil {
			return nil, fmt.Errorf("scan current reviewer: %w", err)
		}

		currentReviewers = append(currentReviewers, uID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate current reviewers: %w", err)
	}

	for _, uID := range currentReviewers {
		if uID != oldReviewerID {
			otherReviewersID = uID
			break
		}
	}

	const reviewerTeamQuery = `
SELECT team_name
FROM users
WHERE user_id = $1
`

	var reviewerTeam string

	if err = tx.QueryRow(ctx, reviewerTeamQuery, oldReviewerID).
		Scan(&reviewerTeam); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("old reviewer %q not found", oldReviewerID)
		}

		return nil, fmt.Errorf("select reviewer team: %w", err)
	}

	const candidatesQuery = `
SELECT user_id
FROM users
WHERE team_name = $1
  AND is_active = TRUE
  AND user_id <> $2
  AND user_id <> $3
`

	candidatesRows, err := tx.Query(ctx, candidatesQuery, reviewerTeam, oldReviewerID, authorID)
	if err != nil {
		return nil, fmt.Errorf("select replacement candidates: %w", err)
	}
	defer candidatesRows.Close()

	candidatesID := make([]string, 0)

	for candidatesRows.Next() {
		var uID string
		if err := candidatesRows.Scan(&uID); err != nil {
			return nil, fmt.Errorf("scan replacement candidate: %w", err)
		}
		candidatesID = append(candidatesID, uID)
	}

	if err := candidatesRows.Err(); err != nil {
		return nil, fmt.Errorf("iterate replacement candidates: %w", err)
	}

	if len(candidatesID) == 0 {
		return nil, models.ErrorCodeNoCandidate
	}

	newReviewerID := chooseRandom(candidatesID, 1)[0]

	const deleteOld = `
DELETE FROM pull_request_reviewers
WHERE pull_request_id = $1 AND user_id = $2
`

	if _, err := tx.Exec(ctx, deleteOld, prID, oldReviewerID); err != nil {
		return nil, fmt.Errorf("delete old reviewer: %w", err)
	}

	if otherReviewersID == "" || newReviewerID != otherReviewersID {
		const insertNew = `
INSERT INTO pull_request_reviewers (pull_request_id, user_id)
VALUES ($1, $2)
`
		if _, err := tx.Exec(ctx, insertNew, prID, newReviewerID); err != nil {
			return nil, fmt.Errorf("insert new reviewer: %w", err)
		}
	}

	updatedPR, err := r.loadReviewersTx(ctx, tx, prID)
	if err != nil {
		return nil, err
	}

	updatedPR.NewReviewerID = newReviewerID

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit tx (reassign PR): %w", err)
	}

	return updatedPR, nil

}

func (r *pullRequestRepository) MergePullRequest(ctx context.Context, prID string) (*models.PullRequest, error) {
	if prID == "" {
		return nil, fmt.Errorf("prID is required")
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx (merge PR): %w", err)
	}
	defer tx.Rollback(ctx)

	const selectStatusQuery = `
SELECT status
FROM pull_requests
WHERE pull_request_id = $1
FOR UPDATE
`

	var status string

	if err = tx.QueryRow(ctx, selectStatusQuery, prID).Scan(&status); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrorCodePRNotFound
		}

		return nil, fmt.Errorf("selected pull_request status: %w", err)
	}

	if status == string(models.PullRequestStatusOpen) {
		const updateStatus = `
UPDATE pull_requests
SET status = $2,
    merged_at = NOW()
WHERE pull_request_id = $1
`
		if _, err := tx.Exec(ctx, updateStatus, prID, string(models.PullRequestStatusMerged)); err != nil {
			return nil, fmt.Errorf("update pull_request status to MERGED: %w", err)
		}

		status = string(models.PullRequestStatusMerged)
	}

	pr, err := r.loadReviewersTx(ctx, tx, prID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit tx (merge PR): %w", err)
	}

	return pr, nil
}

func (r *pullRequestRepository) loadReviewersTx(ctx context.Context, tx pgx.Tx, prID string) (*models.PullRequest, error) {
	const prQuery = `
SELECT pull_request_id, pull_request_name, author_id, status, merged_at
FROM pull_requests
WHERE pull_request_id = $1
`

	var (
		id       string
		name     string
		authorID string
		status   string
		mergedAt *time.Time
	)

	err := tx.QueryRow(ctx, prQuery, prID).
		Scan(&id, &name, &authorID, &status, &mergedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrorCodePRNotFound
		}

		return nil, fmt.Errorf("select pull_request: %w", err)
	}

	const reviewersQuery = `
SELECT user_id
FROM pull_request_reviewers
WHERE pull_request_id = $1
ORDER BY user_id
`

	rows, err := tx.Query(ctx, reviewersQuery, prID)
	if err != nil {
		return nil, fmt.Errorf("select pull_request_reviewers: %w", err)
	}
	defer rows.Close()

	reviewers := make([]string, 0)

	for rows.Next() {
		var uid string
		if scanErr := rows.Scan(&uid); scanErr != nil {
			return nil, fmt.Errorf("scan reviewer: %w", scanErr)
		}
		reviewers = append(reviewers, uid)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, fmt.Errorf("iterate reviewers: %w", rowsErr)
	}

	pr := &models.PullRequest{
		PullRequestID:     id,
		PullRequestName:   name,
		AuthorID:          authorID,
		Status:            models.PullRequestStatus(status),
		AssignedReviewers: reviewers,
		MergedAt:          mergedAt,
	}

	return pr, nil
}

func isUnique(err error) bool {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}

	return false
}

func chooseRandom(ids []string, maxCount int) []string {
	if len(ids) == 0 || maxCount <= 0 {
		return nil
	}

	if len(ids) <= maxCount {
		out := make([]string, len(ids))
		copy(out, ids)
		return out
	}

	idx := rand.Perm(len(ids))
	out := make([]string, 0, maxCount)

	for i := 0; i < maxCount; i++ {
		out = append(out, ids[idx[i]])
	}

	return out
}

// func (r *pullRequestRepository) loadPullRequestWithReviewers(ctx context.Context, prID string) (*models.PullRequest, error) {
// 	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
// 	if err != nil {
// 		return nil, fmt.Errorf("begin tx (load PR): %w", err)
// 	}
// 	defer tx.Rollback(ctx)

// 	pr, err := r.loadReviewersTx(ctx, tx, prID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := tx.Commit(ctx); err != nil {
// 		return nil, fmt.Errorf("commit tx (load PR): %w", err)
// 	}

// 	return pr, nil
// }
