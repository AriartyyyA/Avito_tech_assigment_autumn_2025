package transport

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addTeam(c *gin.Context) {
	var req dto.AddTeamDTO

	if req.Team.TeamName == "" {
		err := InvalidRequest("team_name")
		c.JSON(400, err)
		return
	}

	if err := c.BindJSON(&req); err != nil {
		err := InvalidRequest("")
		c.JSON(400, err)
		return
	}

	team, err := h.services.Team.AddTeam(req.Team)
	if err != nil {
		err := TeamExists()
		c.JSON(400, err)
		return
	}

	c.JSON(200, team)
}

func (h *Handler) getTeam(c *gin.Context) {
	teamName := c.Query("team_name")
	if teamName == "" {
		err := InvalidRequest("team_name")
		c.JSON(400, err)
		return
	}

	team, err := h.services.Team.GetTeam(teamName)
	if err != nil {
		err := NotFound("Team")
		c.JSON(404, err)
		return
	}

	c.JSON(200, team)
}
