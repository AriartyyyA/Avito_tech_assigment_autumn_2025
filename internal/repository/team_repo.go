package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TeamRepository struct {
	db *pgxpool.Pool
}

func NewTeamRepository(db *pgxpool.Pool) Team {
	return &TeamRepository{
		db: db,
	}
}

func (t *TeamRepository) AddTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx,
		`INSERT INTO teams (team_name) VALUES ($1)`,
		team.TeamName,
	); err != nil {
		var pgErr *pgconn.PgError
		// 23505 -- код ошибки, означающий, что запись уже существует
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, models.ErrorCodeTeamExists
		}

		return nil, fmt.Errorf("insert team: %w", err)
	}

	for _, member := range team.Members {
		if member.UserID == "" || member.Username == "" {
			return nil, models.ErrorCodeInvalidRequest
		}

		if _, err := tx.Exec(ctx,
			`INSERT INTO users (user_id, username, team_name, is_active)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (user_id) DO UPDATE
			SET username = EXCLUDED.username,
				team_name = EXCLUDED.team_name,
				is_active = EXCLUDED.is_active`,
			member.UserID,
			member.Username,
			team.TeamName,
			member.IsActive,
		); err != nil {
			return nil, fmt.Errorf("upsert user: %s: %w", member.UserID, err)
		}
	}

	rows, err := tx.Query(ctx,
		`SELECT user_id, username, is_active
		FROM users
		WHERE team_name = $1
		ORDER BY user_id`,
		team.TeamName,
	)
	if err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}
	defer rows.Close()

	members := make([]models.TeamMember, 0)

	for rows.Next() {
		var m models.TeamMember

		if err := rows.Scan(&m.UserID, &m.Username, &m.IsActive); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}

		members = append(members, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	createdTeam := &models.Team{
		TeamName: team.TeamName,
		Members:  members,
	}

	return createdTeam, nil
}

func (t *TeamRepository) GetTeam(ctx context.Context, teamName string) (*models.Team, error) {
	var name string

	if err := t.db.QueryRow(ctx,
		`SELECT team_name
		 FROM teams
		 WHERE team_name = $1`,
		teamName,
	).Scan(&name); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrorCodeTeamNotFound
		} else {
			return nil, fmt.Errorf("select team: %w", err)
		}
	}

	rows, err := t.db.Query(ctx,
		`SELECT user_id, username, is_active
		 FROM users
		 WHERE team_name = $1
		 ORDER BY user_id`,
		teamName,
	)
	if err != nil {
		return nil, fmt.Errorf("select team members: %w", err)
	}
	defer rows.Close()

	members := make([]models.TeamMember, 0)

	for rows.Next() {
		var m models.TeamMember

		if err := rows.Scan(&m.UserID, &m.Username, &m.IsActive); err != nil {
			return nil, fmt.Errorf("scan team member: %w", err)
		}

		members = append(members, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	team := &models.Team{
		TeamName: name,
		Members:  members,
	}

	return team, nil
}
