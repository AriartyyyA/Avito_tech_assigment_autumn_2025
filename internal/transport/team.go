package transport

import (
	"log"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addTeam(c *gin.Context) {
	var req models.Team

	if err := c.BindJSON(&req); err != nil {
		log.Println("Error in handler")
		err := InvalidRequest("")
		c.JSON(400, err)
		return
	}

	if req.TeamName == "" {
		err := InvalidRequest("team_name")
		c.JSON(400, err)
		return
	}

	team, err := h.services.Team.AddTeam(c.Request.Context(), &req)
	if err != nil {
		err := TeamExists()
		c.JSON(400, err)
		return
	}

	c.JSON(200, team)
}

func (h *Handler) getTeam(c *gin.Context) {
	teamName := c.Param("team_name")
	if teamName == "" {
		err := InvalidRequest("team_name")
		c.JSON(400, err)
		return
	}

	team, err := h.services.Team.GetTeam(c.Request.Context(), teamName)
	if err != nil {
		err := NotFound("Team")
		c.JSON(404, err)
		return
	}

	c.JSON(200, team)
}
