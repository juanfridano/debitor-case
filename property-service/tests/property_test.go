package tests

import (
	"net/http"
	"net/http/httptest"
	"property-service/database"
	"property-service/handlers"
	"property-service/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	database.InitDB() // Initialize the database connection

	r.GET("/properties", handlers.GetProperties)
	r.GET("/properties/:id", handlers.GetProperty)
	r.POST("/properties", handlers.CreateProperty)
	r.PUT("/properties/:id", handlers.UpdateProperty)
	r.DELETE("/properties/:id", handlers.DeleteProperty)

	return r
}

func TestCreateProperty(t *testing.T) {
	r := setupRouter()

	// Mock request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/properties", nil)

	r.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusBadRequest, w.Code) // Expecting Bad Request since the body is nil
}

func TestGetProperties(t *testing.T) {
	r := setupRouter()

	// Mock request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/properties", nil)

	r.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateProperty(t *testing.T) {
	r := setupRouter()

	// Create a test property in the database
	property := models.Property{ContractID: 1, Type: "House", Location: "123 St", Specs: map[string]string{"bedrooms": "3", "bathrooms": "2"}}
	database.DB.Create(&property)

	// Mock request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/properties/1", nil)

	r.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
