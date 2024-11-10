package main

import (
	"api_server/config"
	"api_server/db"
	"api_server/routes"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	config.LoadConfig()
	if err := db.InitDB(ctx); err != nil {
		log.Fatal("error initiating db")
		return
	}
	r := gin.Default()
	// setup v1 routes
	v1Route := r.Group("/api/v1")
	routes.SetupV1Routes(v1Route)

	// run on port
	r.Run(config.AppConfig.PORT)
}
