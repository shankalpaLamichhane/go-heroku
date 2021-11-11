package main


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-heroku/models"
	"gorm.io/driver/postgres"
	"log"
	"gorm.io/gorm"
	"fmt"

	"os"
)

var dbConfig struct {
	db *gorm.DB
}

func main() {
	dns := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR IN CONNECTION ::::::: ", err.Error())
		panic(err)
	}
	router := gin.Default()
	router.GET("/v1/tasks", func(c *gin.Context) {
		var tasks []models.Task
		fmt.Print(tasks)
		db.First(&tasks)
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "success",
			"data":    tasks,
		})
	})
		http.ListenAndServe(":"+os.Getenv("PORT"),router)
}