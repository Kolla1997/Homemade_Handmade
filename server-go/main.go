package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize storage and handlers
	storage := NewStorage()
	handlers := NewHandlers(storage)

	r := gin.Default()

	// Your existing routes
	r.GET("/api/menu/:category", handlers.GetMenuByCategory)
	r.POST("/api/orders", handlers.CreateOrder)
	r.GET("/api/orders", handlers.GetOrders)
	r.GET("/api/orders/:id", handlers.GetOrder)
	r.PATCH("/api/orders/:id/status", handlers.UpdateOrderStatus)
	r.POST("/api/contact", handlers.CreateContactMessage)
	r.GET("/api/contact", handlers.GetContactMessages)

	// Add this root handler
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API!",
		})
	})

	// Static assets route if needed
	r.Static("/assets", "./assets")

	// Start the server on port 5000
	r.Run("0.0.0.0:5000")
}