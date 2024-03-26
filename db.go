package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}

// Initialize the database and seed data
func initAndSeedDB() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Create users table if not exists
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255),
		avatar VARCHAR(255),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating users table: %s", err)
	}

	// Seed the users table
	// Adjust seeding logic as per your requirement
	seedUsersSQL := `INSERT INTO users (username, password, avatar) VALUES
		('JohnDoe', 'password123', 'avatar1.jpg'),
		('JaneDoe', 'password123', 'avatar2.jpg')
		ON CONFLICT (username) DO NOTHING;`
	_, err = db.Exec(seedUsersSQL)
	if err != nil {
		log.Fatalf("Error seeding users table: %s", err)
	}
}
