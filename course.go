package main

type Course struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
