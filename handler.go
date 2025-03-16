package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourseHandler(c *gin.Context, state *State) {
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
	course.Id = state.IdCounter

	// Store the course
	state.Courses = append(state.Courses, course)

	c.JSON(http.StatusCreated, gin.H{
		"data": course,
	})

	// Increment the counter
	state.IdCounter += 1
}

func GetCoursesHandler(c *gin.Context, state *State) {
	c.JSON(http.StatusCreated, gin.H{
		"data": ReverseCourses(state.Courses),
	})
}
