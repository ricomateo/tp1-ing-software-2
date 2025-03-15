package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Id = uint64

func main() {
	r := gin.Default()

	idCounter := Id(0)
	coursesById := make(map[Id]Course)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://editor.swagger.io"}

	r.Use(cors.New(config))
	r.POST("/courses", func(c *gin.Context) {
		id := idCounter
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

		// Store the course
		coursesById[id] = course

		data := make(map[string]interface{})
		data["description"] = course.Description
		data["title"] = course.Title
		data["id"] = id

		c.JSON(http.StatusCreated, gin.H{
			"data": data,
		})

		// Increment the counter
		idCounter += 1
	})

	r.Run()
}
