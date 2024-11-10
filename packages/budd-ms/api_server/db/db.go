package db

import (
	"api_server/utils"
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func InitDB(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, 2000*time.Millisecond)
	conn, err := pgx.Connect(ctx, utils.AppConfig.POSTGRES_URL)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
		return err
	}
	log.Println("Connected to db")
	defer conn.Close(c)
	defer cancel()
	return nil
}
