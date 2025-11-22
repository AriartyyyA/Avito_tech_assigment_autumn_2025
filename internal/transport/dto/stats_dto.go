package dto

import "github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"

type UserAssignmentsStatsResponse struct {
	Stats []models.UserAssignmentsStat `json:"stats"`
}
