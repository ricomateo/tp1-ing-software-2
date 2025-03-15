package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	idCounter := uint64(0)
	courses := make([]Course, 0)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://editor.swagger.io"}

	r.Use(cors.New(config))
	r.POST("/courses", func(c *gin.Context) {
		course := Course{}
		if err := c.ShouldBind(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"title":    "Bad request",
				"type":     "about:blank",
				"status":   http.StatusBadRequest,
				"detail":   "Could not create the course",
				"instance": "/courses",
			})
			return
		}
		course.Id = idCounter

		// Store the course
		courses = append(courses, course)

		c.JSON(http.StatusCreated, gin.H{
			"data": course,
		})

		// Increment the counter
		idCounter += 1
	})

	r.GET("/courses", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"data": courses,
		})
	})

	r.Run()
}
