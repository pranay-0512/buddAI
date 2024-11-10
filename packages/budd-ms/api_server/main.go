package main

import (
	"api_server/db"
	"api_server/routes"
	"api_server/utils"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	utils.LoadConfig()
	if err := db.InitDB(ctx); err != nil {
		log.Fatal("error initiating db")
		return
	}
	r := gin.Default()
	// setup v1 routes
	v1Route := r.Group("/api/v1")
	routes.SetupV1Routes(v1Route)

	// run on port
	r.Run(utils.AppConfig.PORT)
}
