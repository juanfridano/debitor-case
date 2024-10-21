package handlers

import (
	"net/http"
	"people-service/database"
	"people-service/models"

	"github.com/gin-gonic/gin"
)

func GetPeople(c *gin.Context) {
	var people []models.Person
	database.DB.Find(&people)
	c.JSON(http.StatusOK, people)
}

func GetPerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := database.DB.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func CreatePerson(c *gin.Context) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&person)
	c.JSON(http.StatusOK, person)
}

func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := database.DB.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&person)
	c.JSON(http.StatusOK, person)
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Person{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
}
