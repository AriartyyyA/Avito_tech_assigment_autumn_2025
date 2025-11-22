package integration

import (
	"context"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/service"
)

// MockUserService - мок для User сервиса
type MockUserService struct {
	SetIsActiveFunc          func(ctx context.Context, userID string, isActive bool) (*models.User, error)
	GetReviewFunc            func(ctx context.Context, userID string) ([]models.PullRequestShort, error)
	GetUserAssignmentsStatsFunc func(ctx context.Context) ([]models.UserAssignmentsStat, error)
}

func (m *MockUserService) SetIsActive(ctx context.Context, userID string, isActive bool) (*models.User, error) {
	if m.SetIsActiveFunc != nil {
		return m.SetIsActiveFunc(ctx, userID, isActive)
	}
	return nil, nil
}

func (m *MockUserService) GetReview(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
	if m.GetReviewFunc != nil {
		return m.GetReviewFunc(ctx, userID)
	}
	return nil, nil
}

func (m *MockUserService) GetUserAssignmentsStats(ctx context.Context) ([]models.UserAssignmentsStat, error) {
	if m.GetUserAssignmentsStatsFunc != nil {
		return m.GetUserAssignmentsStatsFunc(ctx)
	}
	return nil, nil
}

// MockPullRequestService - мок для PullRequest сервиса
type MockPullRequestService struct {
	CreatePullRequestFunc   func(ctx context.Context, pullRequestID, pullRequestName, authorID string) (*models.PullRequest, error)
	MergePullRequestFunc    func(ctx context.Context, prID string) (*models.PullRequest, error)
	ReassignPullRequestFunc func(ctx context.Context, prID string, oldReviewerID string) (*models.PullRequest, error)
}

func (m *MockPullRequestService) CreatePullRequest(ctx context.Context, pullRequestID, pullRequestName, authorID string) (*models.PullRequest, error) {
	if m.CreatePullRequestFunc != nil {
		return m.CreatePullRequestFunc(ctx, pullRequestID, pullRequestName, authorID)
	}
	return nil, nil
}

func (m *MockPullRequestService) MergePullRequest(ctx context.Context, prID string) (*models.PullRequest, error) {
	if m.MergePullRequestFunc != nil {
		return m.MergePullRequestFunc(ctx, prID)
	}
	return nil, nil
}

func (m *MockPullRequestService) ReassignPullRequest(ctx context.Context, prID string, oldReviewerID string) (*models.PullRequest, error) {
	if m.ReassignPullRequestFunc != nil {
		return m.ReassignPullRequestFunc(ctx, prID, oldReviewerID)
	}
	return nil, nil
}

// MockTeamService - мок для Team сервиса
type MockTeamService struct {
	AddTeamFunc            func(ctx context.Context, team *models.Team) (*models.Team, error)
	GetTeamFunc            func(ctx context.Context, teamName string) (*models.Team, error)
	GetTeamPullRequestsFunc func(ctx context.Context, teamName string) ([]models.PullRequestShort, error)
	DeactivateTeamFunc     func(ctx context.Context, teamName string) (*models.TeamDeactivate, error)
}

func (m *MockTeamService) AddTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	if m.AddTeamFunc != nil {
		return m.AddTeamFunc(ctx, team)
	}
	return nil, nil
}

func (m *MockTeamService) GetTeam(ctx context.Context, teamName string) (*models.Team, error) {
	if m.GetTeamFunc != nil {
		return m.GetTeamFunc(ctx, teamName)
	}
	return nil, nil
}

func (m *MockTeamService) GetTeamPullRequests(ctx context.Context, teamName string) ([]models.PullRequestShort, error) {
	if m.GetTeamPullRequestsFunc != nil {
		return m.GetTeamPullRequestsFunc(ctx, teamName)
	}
	return nil, nil
}

func (m *MockTeamService) DeactivateTeam(ctx context.Context, teamName string) (*models.TeamDeactivate, error) {
	if m.DeactivateTeamFunc != nil {
		return m.DeactivateTeamFunc(ctx, teamName)
	}
	return nil, nil
}

// MockService - мок для Service
type MockService struct {
	User        service.User
	PullRequest service.PullRequest
	Team        service.Team
}

func NewMockService() *MockService {
	return &MockService{
		User:        &MockUserService{},
		PullRequest: &MockPullRequestService{},
		Team:        &MockTeamService{},
	}
}

