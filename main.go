package main

import (
	"fmt"
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
			c.String(http.StatusBadRequest, "Bad request error")
			return
		}
		// Store the course
		coursesById[id] = course

		// Increment the counter
		idCounter += 1

		c.String(http.StatusCreated, "Course created successfully")

		// Debugging stuff
		for id, course := range coursesById {
			fmt.Println("id: ", id, "course.title: ", course.Title, "course.descr: ", course.Description)
		}
	})

	r.Run()
}
