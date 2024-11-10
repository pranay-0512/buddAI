package utils

import (
	"api_server/db"
	"api_server/utils/queries"
	"context"
	"log"
)

func DoesTableExist(ctx context.Context, table_name string) bool {
	var exists bool
	err := db.DB.QueryRow(ctx, queries.TABLE_EXISTS).Scan(&exists)
	if err != nil {
		log.Println("error checking for table", err)
		return false
	}
	return exists
}

func DoesRowExist(ctx context.Context, req interface{}, table_name string) bool {
	var exists bool
	err := db.DB.QueryRow(ctx, queries.ROW_EXISTS).Scan(&exists)
	if err != nil {
		log.Println("error checking for row", err)
		return false
	}
	return exists
}

func InsertIntoTable(ctx context.Context, req interface{}, table_name string) {

}
