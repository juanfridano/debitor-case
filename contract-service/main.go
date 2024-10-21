package main

import (
	"contract-service/database"
	"contract-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.InitDB()

	r.GET("/contracts", handlers.GetContracts)
	r.GET("/contracts/:id", handlers.GetContract)
	r.POST("/contracts", handlers.CreateContract)
	r.PUT("/contracts/:id", handlers.UpdateContract)
	r.DELETE("/contracts/:id", handlers.DeleteContract)

	r.Run(":8082") // Contract service running on port 8082
}
