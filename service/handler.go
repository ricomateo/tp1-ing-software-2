package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCourseHandler is the handler for the 'POST /courses' endpoint.
// It creates a new course with the given title and description.
// If the request is wrong, it responds a 400 error.
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

// GetCoursesHandler is the handler for the 'GET /courses' endpoint.
// It responds an array containing all the courses.
func GetCoursesHandler(c *gin.Context, state *State) {
	c.JSON(http.StatusCreated, gin.H{
		"data": ReverseCourses(state.Courses),
	})
}

// GetCourseHandler is the handler for the 'GET /courses/{id}' endpoint.
// On success, it returns the requested course.
// If the course does not exist, it responds with a 404 error.
func GetCourseHandler(c *gin.Context, state *State) {
	id := c.Param("id")

	// Search for the course
	for _, course := range state.Courses {
		courseId := strconv.Itoa(int(course.Id))
		if courseId == id {
			c.JSON(http.StatusCreated, gin.H{
				"data": course,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"title":    "Course not found",
		"type":     "about:blank",
		"status":   http.StatusNotFound,
		"detail":   "The course with ID " + id + " was not found",
		"instance": "/courses/" + id,
	})
}

// DeleteCourseHandler is the handler for the 'DELETE /courses/{id}' endpoint.
// It deletes the given course. If the course does not exist, it responds with a 404 error.
func DeleteCourseHandler(c *gin.Context, state *State) {
	id := c.Param("id")

	// Search for the course
	for index, course := range state.Courses {
		courseId := strconv.Itoa(int(course.Id))
		if courseId == id {
			// Remove the course from the array
			state.Courses = RemoveCourseWithIndex(state.Courses, index)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"title":    "Course not found",
		"type":     "about:blank",
		"status":   http.StatusNotFound,
		"detail":   "The course with ID " + id + " was not found",
		"instance": "/courses/" + id,
	})
}
