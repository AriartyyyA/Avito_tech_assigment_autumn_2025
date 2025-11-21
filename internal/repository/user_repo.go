package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetReview(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
	const query = `SELECT
    pr.pull_request_id,
    pr.pull_request_name,
    pr.author_id,
    pr.status
FROM pull_requests AS pr
JOIN pull_request_reviewers AS prr
    ON pr.pull_request_id = prr.pull_request_id
WHERE prr.user_id = $1
ORDER BY pr.created_at DESC, pr.pull_request_id`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("get review: %w", err)
	}
	defer rows.Close()

	result := make([]models.PullRequestShort, 0)

	for rows.Next() {
		var pr models.PullRequestShort
		var status string

		if err := rows.Scan(&pr.PullRequestID, &pr.PullRequestName, &pr.AuthorID, &status); err != nil {
			return nil, fmt.Errorf("scan review: %w", err)
		}

		pr.Status = models.PullRequestStatus(status)
		result = append(result, pr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows review: %w", err)
	}

	return result, nil
}

func (r *userRepository) SetIsActive(ctx context.Context, userID string, isActive bool) (*models.User, error) {
	const query = `UPDATE users
		SET is_active = $2
		WHERE user_id = $1
		RETURNING user_id, username, team_name, is_active`

	var user models.User

	if err := r.db.QueryRow(ctx, query, userID, isActive).
		Scan(&user.ID, &user.Name, &user.TeamName, &user.IsActive); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrorCodeUserNotFound
		}

		return nil, fmt.Errorf("set user is_active: %w", err)
	}

	return &user, nil
}
