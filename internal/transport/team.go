package transport

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addTeam(c *gin.Context) {
	var req dto.AddTeamDTO

	if err := c.BindJSON(&req); err != nil {
		return
	}

	team, err := h.services.Team.AddTeam(req.Team)
	if err != nil {
		return
	}

	c.JSON(200, team)

}

func (h *Handler) getTeam(c *gin.Context) {
	teamName := c.Query("team_name")
	if teamName == "" {
		return
	}

	team, err := h.services.Team.GetTeam(teamName)
	if err != nil {
		return
	}

	c.JSON(200, team)
}
