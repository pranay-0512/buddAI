package main

import (
	"api_server/routes"
	"api_server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadConfig()

	r := gin.Default()

	// setup v1 routes
	v1Route := r.Group("/api/v1")
	routes.SetupV1Routes(v1Route)

	// run on port
	r.Run(utils.AppConfig.PORT)
}
