package routes

import (
	"api_server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupV1Routes(router *gin.RouterGroup) {
	authRoute := router.Group("/auth")
	{
		authRoute.POST("/signup", handlers.SignUpAuth)
		authRoute.POST("/signin", handlers.SignInAuth)
	}
}
