package main

import (
	"property-service/database"
	"property-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.InitDB()

	r.GET("/properties", handlers.GetProperties)
	r.GET("/properties/:id", handlers.GetProperty)
	r.POST("/properties", handlers.CreateProperty)
	r.PUT("/properties/:id", handlers.UpdateProperty)
	r.DELETE("/properties/:id", handlers.DeleteProperty)

	r.Run(":8083") // Property service running on port 8083
}
