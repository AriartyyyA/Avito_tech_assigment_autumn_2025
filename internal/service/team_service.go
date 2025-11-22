package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
)

type TeamService struct {
	repository *repository.Repository
}

func NewTeamService(repository *repository.Repository) Team {
	return &TeamService{
		repository: repository,
	}
}

// CreateTeam implements TeamInterface.
func (s *TeamService) AddTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	createdTeam, err := s.repository.Team.AddTeam(ctx, team)
	if err != nil {
		return nil, err
	}

	return createdTeam, nil
}

// GetTeam implements TeamInterface.
func (s *TeamService) GetTeam(ctx context.Context, teamName string) (*models.Team, error) {
	team, err := s.repository.Team.GetTeam(ctx, teamName)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (s *TeamService) GetTeamPullRequests(ctx context.Context, teamName string) ([]models.PullRequestShort, error) {
	if _, err := s.repository.Team.GetTeam(ctx, teamName); err != nil {
		return nil, err
	}

	return s.repository.Team.GetTeamPullRequests(ctx, teamName)
}

func (s *TeamService) DeactivateTeam(ctx context.Context, teamName string) (*models.TeamDeactivate, error) {
	team, err := s.repository.Team.GetTeam(ctx, teamName)
	if err != nil {
		return nil, err
	}

	activeUsersID := make([]string, 0, len(team.Members))
	for _, member := range team.Members {
		if member.IsActive {
			activeUsersID = append(activeUsersID, member.UserID)
		}
	}

	result := &models.TeamDeactivate{
		TeamName:         teamName,
		DeactivatedUsers: activeUsersID,
	}

	if len(activeUsersID) == 0 {
		return result, nil
	}

	for _, userID := range activeUsersID {
		if _, err := s.repository.UserRepository.SetIsActive(ctx, userID, false); err != nil {
			if errors.Is(err, models.ErrorCodeUserNotFound) {
				continue
			}
			return nil, fmt.Errorf("deactivate user: %w", err)
		}
	}

	for _, userID := range activeUsersID {
		prs, err := s.repository.UserRepository.GetReview(ctx, userID)
		if err != nil {
			if errors.Is(err, models.ErrorCodeUserNotFound) {
				continue
			}
			return nil, fmt.Errorf("get user review: %w", err)
		}

		for _, pr := range prs {
			if pr.Status != models.PullRequestStatusOpen {
				continue
			}

			result.OpenPRCount++

			if _, err := s.repository.PullRequestRepository.ReassignPullRequest(ctx, pr.PullRequestID, userID); err != nil {
				switch {
				case errors.Is(err, models.ErrorCodePRNotFound),
					errors.Is(err, models.ErrorCodePRMerged),
					errors.Is(err, models.ErrorCodeNotAssigned),
					errors.Is(err, models.ErrorCodeNoCandidate):
					result.FailedReassignments++
					continue
				default:
					return nil, fmt.Errorf("reassign PR %q for user %q: %w", pr.PullRequestID, userID, err)
				}
			}

			result.SuccessfulReassignments++
		}
	}

	return result, nil
}
