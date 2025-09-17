package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		err := db.Ping()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "unhealty", "error": err})
		}
		c.JSON(200, gin.H{"message": "healthy"})
	})

	// api := r.Group("/api/v1")
	// {
	// 	api.GET("/books", getBooks)
	// 	api.GET("/books/:id", getBook)
	// 	api.POST("/books", createBook)
	// 	api.PUT("/books/:id", updateBook)
	// 	api.DELETE("/books/:id", deleteBook)
	// }

	r.Run(":8080")
}
