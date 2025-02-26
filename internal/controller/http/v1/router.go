package v1

import (
	"github.com/DMAproject25/auth-service/internal/usecase"
	"github.com/DMAproject25/auth-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, logger logger.Interface, todoUseCase usecase.Todos, userUseCase usecase.Users) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	router := handler.Group("/api")

	newCommonRoutes(router)
	newTodosRoutes(router, logger, todoUseCase)
	newUserRoutes(router, logger, userUseCase)
}
