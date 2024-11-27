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

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(fmt.Sprintf("Error creating table: %v", err))
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			date_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(fmt.Sprintf("Error creating table: %v", err))
	}

	fmt.Println("Tables created!")
}
