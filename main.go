package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const SWAGGER_URL string = "https://editor.swagger.io"

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		log.Fatal("Missing required environment variable HOST")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Missing required environment variable PORT")
	}
	r := gin.Default()
	state := State{
		Courses:   make([]Course, 0),
		IdCounter: uint64(0),
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{SWAGGER_URL}

	r.Use(cors.New(config))

	r.POST("/courses", func(c *gin.Context) {
		CreateCourseHandler(c, &state)
	})
	r.GET("/courses", func(c *gin.Context) {
		GetCoursesHandler(c, &state)
	})
	address := host + ":" + port
	r.Run(address)
}
