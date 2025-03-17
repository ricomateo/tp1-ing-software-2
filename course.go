package main

type Course struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Id          uint64 `json:"id"`
}

// Receives an array of courses and returns an array of courses in the inverse order.
func ReverseCourses(courses []Course) []Course {
	lenCourses := len(courses)
	reversedCourses := make([]Course, lenCourses)
	for i := 0; i < len(courses); i++ {
		reversedCourses[i] = courses[lenCourses-1-i]
	}
	return reversedCourses
}

// Receives an array of courses and removes the course at the given index.
func RemoveCourseWithIndex(courses []Course, index int) []Course {
	return append(courses[:index], courses[index+1:]...)
}
