package dbconnection

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
  _ "github.com/lib/pq"
)
 

var db *sql.DB

func InitDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env variables")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	 dbHost, dbUsername, dbPassword, dbName, dbPort))

	 if err != nil {
		panic(err.Error())
	 }

	 err = db.Ping()

	 if err != nil {
		panic(err.Error())
	 }

	 fmt.Println("Successfully connected to the database")

	 _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email	VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	 )`)

	 if err != nil {
		panic(err.Error())
	 }

	 _, err = db.Exec(`CREATE TABLE IF NOT EXISTS measurements (
		id SERIAL PRIMARY KEY,
		userId INT NOT NULL,
		weight DOUBLE PRECISION NOT NULL,
    height DOUBLE PRECISION NOT NULL,
    body_fat DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	 )`)

	 if err != nil {
		panic(err.Error())
	 }

}
