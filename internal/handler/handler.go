package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"visitor/internal/service"
)

type Handler struct {
	services *service.Service
	log      *slog.Logger
}

func NewHandler(services *service.Service, log *slog.Logger) *Handler {
	return &Handler{
		services: services,
		log:      log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			users.Use(debugStartMiddleware)
			{
				users.POST("/", h.addUsersHandler, debugEndMiddleware)
				users.GET("/", h.getUserHandler, debugEndMiddleware)
				users.PUT("/:id", h.updateUserHandler, debugEndMiddleware)
				users.DELETE("/:id", h.deleteUserHandler, debugEndMiddleware)
			}
		}
	}

	return router
}
