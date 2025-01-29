package main

import (
	"net/http"
	"testing/internal/api"
	"testing/internal/repository"
	"testing/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewUserRepository()
	useCase := usecase.NewUserUseCase(repo)

	r := api.SetupRouter(useCase)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})

	r.Run(":8080")
}
