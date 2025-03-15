package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	state := State{
		Courses:   make([]Course, 0),
		IdCounter: uint64(0),
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://editor.swagger.io"}

	r.Use(cors.New(config))

	r.POST("/courses", func(c *gin.Context) {
		CreateCourseHandler(c, &state)
	})
	r.GET("/courses", func(c *gin.Context) {
		GetCoursesHandler(c, &state)
	})

	r.Run()
}
