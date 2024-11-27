package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to PostgreSQL: %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(3)

	// Test the connection
	if err := DB.Ping(); err != nil {
		panic(fmt.Sprintf("Error pinging PostgreSQL: %v", err))
	}

	fmt.Println("Connected to PostgreSQL!")
	createTable()
}

func createTable() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			date_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			user_id INTEGER NOT NULL
		)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic(fmt.Sprintf("Error creating table: %v", err))
	}

	fmt.Println("Events table created or already exists.")
}
