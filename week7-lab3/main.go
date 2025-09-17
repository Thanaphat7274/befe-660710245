package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue

}

var db *sql.DB

func initDB() {
	var err error
	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")
	conST := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	// fmt.Println(conST)
	db, err = sql.Open("postgres", conST)
	if err != nil {
		log.Fatal("failed to open database")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connet to database", err)
	}
	log.Println("successfully connet to database")
}
func main() {
	initDB()
}
