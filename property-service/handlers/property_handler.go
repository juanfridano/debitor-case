package handlers

import (
	"net/http"
	"property-service/database"
	"property-service/models"

	"github.com/gin-gonic/gin"
)

func GetProperties(c *gin.Context) {
	var properties []models.Property
	database.DB.Find(&properties)
	c.JSON(http.StatusOK, properties)
}

func GetProperty(c *gin.Context) {
	id := c.Param("id")
	var property models.Property
	if err := database.DB.First(&property, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}
	c.JSON(http.StatusOK, property)
}

func CreateProperty(c *gin.Context) {
	var property models.Property
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&property)
	c.JSON(http.StatusOK, property)
}

func UpdateProperty(c *gin.Context) {
	id := c.Param("id")
	var property models.Property
	if err := database.DB.First(&property, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&property)
	c.JSON(http.StatusOK, property)
}

func DeleteProperty(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Property{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Property deleted"})
}
