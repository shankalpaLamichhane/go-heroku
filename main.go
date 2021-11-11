package main


import (
	"github.com/gin-gonic/gin"
	"net/http"

	"os"
)

func main() {
	router := gin.Default()
	router.GET("/v1/tasks", func(c *gin.Context) {
		c.JSON(200,"This is mycodesmells/golang-examples server from Heroku!")
	})
    // http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
    //     fmt.Fprint(rw, "This is mycodesmells/golang-examples server from Heroku!")
    // })
		http.ListenAndServe(":"+os.Getenv("PORT"),router)
}