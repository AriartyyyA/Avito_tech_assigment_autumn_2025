package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
)

type teamService struct {
	repository *repository.Repository
}

func NewTeamService(repository *repository.Repository) TeamService {
	return &teamService{
		repository: repository,
	}
}

// CreateTeam implements TeamInterface.
func (s *teamService) AddTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	createdTeam, err := s.repository.TeamRepository.AddTeam(ctx, team)
	if err != nil {
		return nil, err
	}

	return createdTeam, nil
}

func (s *teamService) GetTeam(ctx context.Context, teamName string) (*models.Team, error) {
	team, err := s.repository.TeamRepository.GetTeam(ctx, teamName)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (s *teamService) GetTeamPullRequests(ctx context.Context, teamName string) ([]models.PullRequestShort, error) {
	if _, err := s.repository.TeamRepository.GetTeam(ctx, teamName); err != nil {
		return nil, err
	}

	return s.repository.TeamRepository.GetTeamPullRequests(ctx, teamName)
}

func (s *teamService) DeactivateTeam(ctx context.Context, teamName string) (*models.TeamDeactivate, error) {
	team, err := s.repository.TeamRepository.GetTeam(ctx, teamName)
	if err != nil {
		return nil, err
	}

	// Собираем активных пользователей
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

	prsWithReviewers, err := s.repository.PullRequestRepository.GetOpenPRsWithTeamReviewers(ctx, teamName, activeUsersID)
	if err != nil {
		return nil, fmt.Errorf("get open PRs with team reviewers: %w", err)
	}

	prSet := make(map[string]bool)
	for _, pr := range prsWithReviewers {
		prSet[pr.PullRequestID] = true
	}
	result.OpenPRCount = len(prSet)

	for _, pr := range prsWithReviewers {
		if _, err := s.repository.PullRequestRepository.ReassignPullRequest(ctx, pr.PullRequestID, pr.ReviewerID); err != nil {
			switch {
			case errors.Is(err, models.ErrorCodePRNotFound),
				errors.Is(err, models.ErrorCodePRMerged),
				errors.Is(err, models.ErrorCodeNotAssigned),
				errors.Is(err, models.ErrorCodeNoCandidate):
				result.FailedReassignments++
				continue
			default:
				return nil, fmt.Errorf("reassign PR %q for reviewer %q: %w", pr.PullRequestID, pr.ReviewerID, err)
			}
		}
		result.SuccessfulReassignments++
	}

	if err := s.repository.UserRepository.DeactivateUsers(ctx, activeUsersID); err != nil {
		return nil, fmt.Errorf("batch deactivate users: %w", err)
	}

	return result, nil
}
