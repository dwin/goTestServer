package model

import (
	"database/sql"

	u "github.com/dwin/goTestServer/app/utils"
	_ "github.com/lib/pq"
)

func getDBconn() *sql.DB {
	db, err := sql.Open("postgres", "user="+u.DBuser+" password="+u.DBpass+" dbname="+u.DBname+" sslmode=disable")
	if err != nil {
		u.Log.Error().Err(err)
	}
	return db
}
