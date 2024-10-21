package tests

import (
	"net/http"
	"net/http/httptest"
	"people-service/database"
	"people-service/handlers"
	"people-service/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	database.InitDB() // Initialize the database connection

	r.GET("/people", handlers.GetPeople)
	r.GET("/people/:id", handlers.GetPerson)
	r.POST("/people", handlers.CreatePerson)
	r.PUT("/people/:id", handlers.UpdatePerson)
	r.DELETE("/people/:id", handlers.DeletePerson)

	return r
}

func TestCreatePerson(t *testing.T) {
	r := setupRouter()

	// Mock request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/people", nil)

	r.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusBadRequest, w.Code) // Expecting Bad Request since the body is nil
}

func TestGetPeople(t *testing.T) {
	r := setupRouter()

	// Mock request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/people", nil)

	r.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdatePerson(t *testing.T) {
	r := setupRouter()

	// Create a test person in the database
	person := models.Person{FirstName: "John", LastName: "Doe", Address: "123 Main St"}
	database.DB.Create(&person)

	// Mock request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/people/1", nil)

	r.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
