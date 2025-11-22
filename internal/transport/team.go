package transport

import (
	"errors"
	"log"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
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
		if errors.Is(err, models.ErrorCodeTeamExists) {
			err := TeamExists()
			c.JSON(400, err)
			return
		}

		err := InternalError()
		c.JSON(500, err)
		return
	}

	resp := dto.AddTeamDTO{
		Team: team,
	}

	c.JSON(201, resp)
}

func (h *Handler) getTeam(c *gin.Context) {
	teamName := c.Query("team_name")
	if teamName == "" {
		err := InvalidRequest("team_name")
		c.JSON(400, err)
		return
	}

	team, err := h.services.Team.GetTeam(c.Request.Context(), teamName)
	if err != nil {
		if errors.Is(err, models.ErrorCodeTeamNotFound) {
			err := NotFound()
			c.JSON(404, err)
			return
		}

		err := NotFound()
		c.JSON(500, err)
		return
	}

	c.JSON(200, team)
}

func (h *Handler) getTeamPullRequests(c *gin.Context) {
	teamName := c.Query("team_name")

	if teamName == "" {
		err := InvalidRequest("team_name")
		c.JSON(400, err)
		return
	}

	prs, err := h.services.Team.GetTeamPullRequests(c.Request.Context(), teamName)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrorCodeTeamNotFound):
			errResp := NotFound()
			c.JSON(404, errResp)
			return
		default:
			errResp := InternalError()
			c.JSON(500, errResp)
			return
		}
	}

	resp := dto.TeamPRsResponse{
		TeamName:    teamName,
		PullRequest: prs,
		PRcount:     len(prs),
	}

	c.JSON(200, resp)
}

func (h *Handler) deactivateTeamUsers(c *gin.Context) {
	var req dto.DeactivateTeamUsersRequest

	if err := c.BindJSON(&req); err != nil {
		err := InvalidRequest("")
		c.JSON(400, err)
		return
	}

	if req.TeamName == "" {
		err := InvalidRequest("team_name")
		c.JSON(400, err)
		return
	}

	result, err := h.services.Team.DeactivateTeam(c.Request.Context(), req.TeamName)
	if err != nil {
		if errors.Is(err, models.ErrorCodeTeamNotFound) {
			err := NotFound()
			c.JSON(404, err)
			return
		}

		err := InternalError()
		c.JSON(500, err)
		return
	}

	resp := dto.DeactivateTeamUsersResponse{
		TeamName:                result.TeamName,
		DeactivatedUsers:        result.DeactivatedUsers,
		OpenPRCount:             result.OpenPRCount,
		SuccessfulReassignments: result.SuccessfulReassignments,
		FailedReassignments:     result.FailedReassignments,
	}

	c.JSON(200, resp)

}
