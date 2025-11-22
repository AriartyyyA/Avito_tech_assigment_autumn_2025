package transport

import (
	"errors"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addTeam(c *gin.Context) {
	req, exists := c.Get("validated_request")
	if !exists {
		err := InternalError()
		c.JSON(500, err)
		return
	}

	teamReq := req.(*models.Team)
	team, err := h.services.Team.AddTeam(c.Request.Context(), teamReq)
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

	team, err := h.services.Team.GetTeam(c.Request.Context(), teamName)
	if err != nil {
		if errors.Is(err, models.ErrorCodeTeamNotFound) {
			err := NotFound(models.ErrorCodeTeamNotFound)
			c.JSON(404, err)
			return
		}

		err := NotFound(models.ErrorCodeNotFound)
		c.JSON(500, err)
		return
	}

	c.JSON(200, team)
}

func (h *Handler) getTeamPullRequests(c *gin.Context) {
	teamName := c.Query("team_name")

	prs, err := h.services.Team.GetTeamPullRequests(c.Request.Context(), teamName)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrorCodeTeamNotFound):
			errResp := NotFound(models.ErrorCodeTeamNotFound)
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
	req, exists := c.Get("validated_request")
	if !exists {
		err := InternalError()
		c.JSON(500, err)
		return
	}

	deactivateReq := req.(*dto.DeactivateTeamUsersRequest)
	result, err := h.services.Team.DeactivateTeam(c.Request.Context(), deactivateReq.TeamName)
	if err != nil {
		if errors.Is(err, models.ErrorCodeTeamNotFound) {
			err := NotFound(models.ErrorCodeTeamNotFound)
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
