package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const SWAGGER_URL string = "https://editor.swagger.io"

func StartService(host, port string) {
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

	r.GET("/courses/:id", func(c *gin.Context) {
		GetCourseHandler(c, &state)
	})
	address := host + ":" + port
	r.Run(address)
}
