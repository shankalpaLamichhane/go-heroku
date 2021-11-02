package main

import (
	"go-heroku/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	version := r.Group("/v1")
	dbUrl := os.Getenv("DATABASE_URL") + " sslmode=require"
	log.Printf("DB [%s]", dbUrl)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	version.GET("/tasks", func(c *gin.Context) {
		var tasks []models.Task
		db.First(&tasks)

		c.JSON(200, gin.H{
			"tasks": tasks,
		})
	})
	r.Run(":8080")

}
