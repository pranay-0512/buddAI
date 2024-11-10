package routes

import "github.com/gin-gonic/gin"

func SetupV1Routes(router *gin.RouterGroup) {
	authRoute := router.Group("/auth")
	{
		authRoute.POST("/signup")
		authRoute.POST("/signin")
	}
}
