package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// Open a new database connection
	db, err := sql.Open("mysql", "root:@/topup_game?parseTime=true")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Try to ping the database to check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	// Assign the database connection to the global variable
	DB = db

	log.Println("Database connected successfully")
}
