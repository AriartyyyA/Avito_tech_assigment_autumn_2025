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
		users.POST("/setIsActive", h.setIsActive)
		users.GET("/getReview", h.getReview)
	}

	team := router.Group("/team")
	{
		team.POST("/add", h.addTeam)
		team.GET("/get", h.getTeam)
	}

	pullRequest := router.Group("/pullRequest")
	{
		pullRequest.POST("/create", h.createPullRequest)
		pullRequest.POST("/merge", h.mergePullRequest)
		pullRequest.POST("/reassign", h.reassignPullRequest)
	}

	return router
}
