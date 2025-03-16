package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const HOST string = "0.0.0.0"
const PORT string = "8080"
const URL string = "http://" + HOST + ":" + PORT

func TestCreateCourse(t *testing.T) {
	go StartService(HOST, PORT)

	time.Sleep(1 * time.Second)
	url := URL + "/courses"

	t.Run("Create course", func(t *testing.T) {
		course := `{"title":"test title","description":"test description"}`
		resp, err := http.Post(url, "application/json", strings.NewReader(course))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		expectedResponse := `{"data":{"title":"test title","description":"test description","id":0}}`
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, expectedResponse, string(body))
		resp.Body.Close()
	})

	t.Run("Get courses", func(t *testing.T) {
		resp, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		expectedResponse := `{"data":[{"title":"test title","description":"test description","id":0}]}`
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, expectedResponse, string(body))
		resp.Body.Close()
	})

	t.Run("Get course with id 0", func(t *testing.T) {
		id := "0"
		resp, err := http.Get(url + "/" + id)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		expectedResponse := `{"data":{"title":"test title","description":"test description","id":0}}`
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, expectedResponse, string(body))
		resp.Body.Close()
	})

	t.Run("Check that an error is returned when requesting a non existing course", func(t *testing.T) {
		id := "1"
		resp, err := http.Get(url + "/" + id)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)

		expectedResponse := `{"detail":"The course with ID 1 was not found","instance":"/courses/1","status":404,"title":"Course not found","type":"about:blank"}`
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, expectedResponse, string(body))
		resp.Body.Close()
	})

	t.Run("Check that courses are created with incremental ids", func(t *testing.T) {
		course := `{"title":"test title 2","description":"test description 2"}`
		resp, err := http.Post(url, "application/json", strings.NewReader(course))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		expectedResponse := `{"data":{"title":"test title 2","description":"test description 2","id":1}}`
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, expectedResponse, string(body))
		resp.Body.Close()
	})

	t.Run("Check that courses are returned in an inversed chronological order", func(t *testing.T) {
		resp, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		expectedResponse := `{"data":[{"title":"test title 2","description":"test description 2","id":1},{"title":"test title","description":"test description","id":0}]}`
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, expectedResponse, string(body))
		resp.Body.Close()
	})
}
