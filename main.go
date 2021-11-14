package main

import (
	"go-heroku/models"
	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	version := r.Group("/v1")
	dbUrl := "postgres://rvgxldtdvoulcj:9fc2b3c5a5c47d4380b970789d13bdbd2af85333786669302b587687e0c8f87a@ec2-52-204-72-14.compute-1.amazonaws.com:5432/d27iah2ij2g50j"
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
	r.Run(":5000")
	http.ListenAndServe(":5001",r)

}
