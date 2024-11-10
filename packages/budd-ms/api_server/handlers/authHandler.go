package handlers

import (
	"api_server/config"
	"api_server/models"
	"api_server/utils"
	"log"

	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(config.AppConfig.JWT_SECRET)

func SignUpAuth(c *gin.Context) {
	var req models.SignupReq
	var users string
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("error binding json", err)
		return
	}
	utils.InsertIntoTable(ctx, req, users)

}

func SignInAuth(c *gin.Context) {

}
