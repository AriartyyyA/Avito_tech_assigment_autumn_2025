package repository

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PullRequestRepository struct {
	db *pgxpool.Pool
}

func NewPullRequestRepository(db *pgxpool.Pool) PullRequest {
	return &PullRequestRepository{
		db: db,
	}
}

func (p *PullRequestRepository) CreatePullRequest(pr models.PullRequest) (models.PullRequest, error) {
	panic("unimplemented")
}

func (p *PullRequestRepository) MergePullRequest(prID string) (models.PullRequest, error) {
	panic("unimplemented")
}

func (p *PullRequestRepository) ReassignPullRequest(prID string, oldReviewerID string) (models.PullRequest, error) {
	panic("unimplemented")
}
