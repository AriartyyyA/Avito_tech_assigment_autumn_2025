package transport

import (
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/service"
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
		users.POST("/set_is_active", h.setIsActive)
		users.GET("/get_review/:id", h.getReview)
	}

	team := router.Group("/team")
	{
		team.GET("/create_team", h.createTeam)
		team.POST("/get_team", h.getTeam)
	}

	pullRequest := router.Group("/pullRequest")
	{
		pullRequest.POST("/create", h.createPullRequest)
		pullRequest.POST("/merge", h.mergePullRequest)
		pullRequest.POST("/reassign", h.reassignPullRequest)
	}

	return router
}
