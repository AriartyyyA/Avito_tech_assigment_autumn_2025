package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active"`
}

func NewUser(id, name, teamName string, isActive bool) *User {
	return &User{
		ID:       id,
		Name:     name,
		TeamName: teamName,
		IsActive: isActive,
	}
}
