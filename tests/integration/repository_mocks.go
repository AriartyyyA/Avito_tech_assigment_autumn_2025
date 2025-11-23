package integration

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
)

type MockUserRepository struct {
	SetIsActiveFunc         func(ctx context.Context, userID string, isActive bool) (*models.User, error)
	GetReviewFunc           func(ctx context.Context, userID string) ([]models.PullRequestShort, error)
	GetAssignmentsStatsFunc func(ctx context.Context) ([]models.UserAssignmentsStat, error)
	DeactivateUsersFunc     func(ctx context.Context, userIDs []string) error
}

func (m *MockUserRepository) SetIsActive(ctx context.Context, userID string, isActive bool) (*models.User, error) {
	if m.SetIsActiveFunc != nil {
		return m.SetIsActiveFunc(ctx, userID, isActive)
	}
	return nil, nil
}

func (m *MockUserRepository) GetReview(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
	if m.GetReviewFunc != nil {
		return m.GetReviewFunc(ctx, userID)
	}
	return nil, nil
}

func (m *MockUserRepository) GetAssignmentsStats(ctx context.Context) ([]models.UserAssignmentsStat, error) {
	if m.GetAssignmentsStatsFunc != nil {
		return m.GetAssignmentsStatsFunc(ctx)
	}
	return nil, nil
}

func (m *MockUserRepository) DeactivateUsers(ctx context.Context, userIDs []string) error {
	if m.DeactivateUsersFunc != nil {
		return m.DeactivateUsersFunc(ctx, userIDs)
	}
	return nil
}

// MockPullRequestRepository мок для PullRequestRepository
type MockPullRequestRepository struct {
	CreatePullRequestFunc           func(ctx context.Context, pr *models.PullRequest) (*models.PullRequest, error)
	MergePullRequestFunc            func(ctx context.Context, prID string) (*models.PullRequest, error)
	ReassignPullRequestFunc         func(ctx context.Context, prID string, OldUserID string) (*models.PullRequest, error)
	GetOpenPRsWithTeamReviewersFunc func(ctx context.Context, teamName string, userIDs []string) ([]repository.PRWithReviewer, error)
}

func (m *MockPullRequestRepository) CreatePullRequest(ctx context.Context, pr *models.PullRequest) (*models.PullRequest, error) {
	if m.CreatePullRequestFunc != nil {
		return m.CreatePullRequestFunc(ctx, pr)
	}
	return nil, nil
}

func (m *MockPullRequestRepository) MergePullRequest(ctx context.Context, prID string) (*models.PullRequest, error) {
	if m.MergePullRequestFunc != nil {
		return m.MergePullRequestFunc(ctx, prID)
	}
	return nil, nil
}

func (m *MockPullRequestRepository) ReassignPullRequest(ctx context.Context, prID string, OldUserID string) (*models.PullRequest, error) {
	if m.ReassignPullRequestFunc != nil {
		return m.ReassignPullRequestFunc(ctx, prID, OldUserID)
	}
	return nil, nil
}

func (m *MockPullRequestRepository) GetOpenPRsWithTeamReviewers(ctx context.Context, teamName string, userIDs []string) ([]repository.PRWithReviewer, error) {
	if m.GetOpenPRsWithTeamReviewersFunc != nil {
		return m.GetOpenPRsWithTeamReviewersFunc(ctx, teamName, userIDs)
	}
	return nil, nil
}

type MockTeamRepository struct {
	AddTeamFunc             func(ctx context.Context, team *models.Team) (*models.Team, error)
	GetTeamFunc             func(ctx context.Context, teamName string) (*models.Team, error)
	GetTeamPullRequestsFunc func(ctx context.Context, teamName string) ([]models.PullRequestShort, error)
}

func (m *MockTeamRepository) AddTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	if m.AddTeamFunc != nil {
		return m.AddTeamFunc(ctx, team)
	}
	return nil, nil
}

func (m *MockTeamRepository) GetTeam(ctx context.Context, teamName string) (*models.Team, error) {
	if m.GetTeamFunc != nil {
		return m.GetTeamFunc(ctx, teamName)
	}
	return nil, nil
}

func (m *MockTeamRepository) GetTeamPullRequests(ctx context.Context, teamName string) ([]models.PullRequestShort, error) {
	if m.GetTeamPullRequestsFunc != nil {
		return m.GetTeamPullRequestsFunc(ctx, teamName)
	}
	return nil, nil
}

func NewMockRepository() *repository.Repository {
	return &repository.Repository{
		UserRepository:        &MockUserRepository{},
		PullRequestRepository: &MockPullRequestRepository{},
		TeamRepository:        &MockTeamRepository{},
	}
}
