package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Db is exported variable
var Db *sql.DB

// InitDB is exported function
func InitDB() {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/grep")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	Db = db
}
