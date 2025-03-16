package main

type Course struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Id          uint64 `json:"id"`
}

func ReverseCourses(courses []Course) []Course {
	lenCourses := len(courses)
	reversedCourses := make([]Course, lenCourses)
	for i := 0; i < len(courses); i++ {
		reversedCourses[i] = courses[lenCourses-1-i]
	}
	return reversedCourses
}

func RemoveCourseWithIndex(courses []Course, index int) []Course {
	return append(courses[:index], courses[index+1:]...)
}
