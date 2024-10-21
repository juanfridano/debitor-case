package tests

import (
	"contract-service/database"
	"contract-service/handlers"
	"contract-service/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	database.InitDB() // Initialize the database connection

	r.GET("/contracts", handlers.GetContracts)
	r.GET("/contracts/:id", handlers.GetContract)
	r.POST("/contracts", handlers.CreateContract)
	r.PUT("/contracts/:id", handlers.UpdateContract)
	r.DELETE("/contracts/:id", handlers.DeleteContract)

	return r
}

func TestCreateContract(t *testing.T) {
	r := setupRouter()

	// Mock request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/contracts", nil)

	r.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusBadRequest, w.Code) // Expecting Bad Request since the body is nil
}

func TestGetContracts(t *testing.T) {
	r := setupRouter()

	// Mock request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/contracts", nil)

	r.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateContract(t *testing.T) {
	r := setupRouter()

	// Create a test contract in the database
	contract := models.Contract{HolderID: 1, PropertyID: 1, Amount: 1000.00}
	database.DB.Create(&contract)

	// Mock request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/contracts/1", nil)

	r.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
