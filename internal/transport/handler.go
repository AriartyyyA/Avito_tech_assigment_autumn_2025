package transport

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/service"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	users := router.Group("/users")
	{
		users.POST("/setIsActive", middleware.ValidateSetIsActiveRequest(), h.setIsActive)
		users.GET("/getReview", middleware.ValidateQueryParams("user_id"), h.getReview)
		users.GET("/userAssignments", h.getUserAssignmentsStats)
	}

	team := router.Group("/team")
	{
		team.POST("/add", middleware.ValidateAddTeamRequest(), h.addTeam)
		team.GET("/get", middleware.ValidateQueryParams("team_name"), h.getTeam)
		team.GET("/pullRequests", middleware.ValidateQueryParams("team_name"), h.getTeamPullRequests)
		team.POST("/deactivateUsers", middleware.ValidateDeactivateTeamUsersRequest(), h.deactivateTeamUsers)
	}

	pullRequest := router.Group("/pullRequest")
	{
		pullRequest.POST("/create", middleware.ValidateCreatePRRequest(), h.createPullRequest)
		pullRequest.POST("/merge", middleware.ValidateMergePRRequest(), h.mergePullRequest)
		pullRequest.POST("/reassign", middleware.ValidateReassignPRRequest(), h.reassignPullRequest)
	}

	return router
}
