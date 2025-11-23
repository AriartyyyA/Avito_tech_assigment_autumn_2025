package transport

import (
	"errors"
	"log"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/models"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addTeam(c *gin.Context) {
	req, exists := c.Get("validated_request")
	if !exists {
		log.Printf("ERROR: validated_request not found in context for addTeam")
		err := InternalError()
		c.JSON(500, err)
		return

	}

	teamReq := req.(*models.Team)
	team, err := h.services.TeamService.AddTeam(c.Request.Context(), teamReq)

	if err != nil {
		if errors.Is(err, models.ErrorCodeTeamExists) {
			log.Printf("WARN: Team already exists: Team=%s", teamReq.TeamName)
			err := TeamExists()
			c.JSON(400, err)
			return

		}
		log.Printf("ERROR: Failed to add team: Team=%s, Error=%v", teamReq.TeamName, err)
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
	team, err := h.services.TeamService.GetTeam(c.Request.Context(), teamName)

	if err != nil {
		if errors.Is(err, models.ErrorCodeTeamNotFound) {
			log.Printf("WARN: Team not found: Team=%s", teamName)
			err := NotFound(models.ErrorCodeTeamNotFound)
			c.JSON(404, err)
			return
		}

		log.Printf("ERROR: Failed to get team: Team=%s, Error=%v", teamName, err)
		err := NotFound(models.ErrorCodeNotFound)
		c.JSON(500, err)
		return
	}

	c.JSON(200, team)

}

func (h *Handler) getTeamPullRequests(c *gin.Context) {
	teamName := c.Query("team_name")
	prs, err := h.services.TeamService.GetTeamPullRequests(c.Request.Context(), teamName)

	if err != nil {
		switch {
		case errors.Is(err, models.ErrorCodeTeamNotFound):
			log.Printf("WARN: Team not found for PRs: Team=%s", teamName)
			errResp := NotFound(models.ErrorCodeTeamNotFound)
			c.JSON(404, errResp)
			return

		default:
			log.Printf("ERROR: Failed to get team PRs: Team=%s, Error=%v", teamName, err)
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
		log.Printf("ERROR: validated_request not found in context for deactivateTeamUsers")
		err := InternalError()
		c.JSON(500, err)
		return
	}

	deactivateReq := req.(*dto.DeactivateTeamUsersRequest)
	result, err := h.services.TeamService.DeactivateTeam(c.Request.Context(), deactivateReq.TeamName)

	if err != nil {
		if errors.Is(err, models.ErrorCodeTeamNotFound) {
			log.Printf("WARN: Team not found for deactivation: Team=%s", deactivateReq.TeamName)
			err := NotFound(models.ErrorCodeTeamNotFound)
			c.JSON(404, err)
			return
		}

		log.Printf("ERROR: Failed to deactivate team: Team=%s, Error=%v", deactivateReq.TeamName, err)
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
