package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDb() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbName := "db6"
	dbUser := ""
	dbPass := ""
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}	

