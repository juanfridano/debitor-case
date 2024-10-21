package handlers

import (
	"contract-service/database"
	"contract-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetContracts(c *gin.Context) {
	var contracts []models.Contract
	database.DB.Preload("Comments").Find(&contracts)
	c.JSON(http.StatusOK, contracts)
}

func GetContract(c *gin.Context) {
	id := c.Param("id")
	var contract models.Contract
	if err := database.DB.Preload("Comments").First(&contract, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract not found"})
		return
	}
	c.JSON(http.StatusOK, contract)
}

func CreateContract(c *gin.Context) {
	var contract models.Contract
	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&contract)
	c.JSON(http.StatusOK, contract)
}

func UpdateContract(c *gin.Context) {
	id := c.Param("id")
	var contract models.Contract
	if err := database.DB.First(&contract, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract not found"})
		return
	}
	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&contract)
	c.JSON(http.StatusOK, contract)
}

func DeleteContract(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Contract{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Contract deleted"})
}
