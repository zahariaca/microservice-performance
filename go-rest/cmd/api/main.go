package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go-rest/database"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting service")
	database.ConnectDatabase()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	log.Println("Gin initialized")

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})
	r.GET("/users", database.GetUsers)
	r.POST("/add", database.AddUser)

	err := r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}
}
