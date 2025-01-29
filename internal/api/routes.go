package api

import (
	"testing/internal/delivery"
	"testing/internal/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter(useCase usecase.UserUseCase) *gin.Engine {
	r := gin.Default()

	handler := delivery.NewUserHandler(useCase)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", handler.AddUser)
		userRoutes.GET("/:id", handler.GetUser)
		userRoutes.PUT("/:id", handler.EditUser)
		userRoutes.DELETE("/:id", handler.DeleteUser)
	}

	return r
}
