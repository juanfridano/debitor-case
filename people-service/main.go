package main

import (
	"people-service/database"
	"people-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.InitDB()

	r.GET("/people", handlers.GetPeople)
	r.GET("/people/:id", handlers.GetPerson)
	r.POST("/people", handlers.CreatePerson)
	r.PUT("/people/:id", handlers.UpdatePerson)
	r.DELETE("/people/:id", handlers.DeletePerson)

	r.Run(":8081") // People service running on port 8081
}
