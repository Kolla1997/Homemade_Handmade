package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Your existing routes
	r.GET("/api/menu/:category", GetMenuByCategory)
	r.POST("/api/orders", CreateOrder)
	r.GET("/api/orders", GetOrders)
	r.GET("/api/orders/:id", GetOrder)
	r.PATCH("/api/orders/:id/status", UpdateOrderStatus)
	r.POST("/api/contact", CreateContactMessage)
	r.GET("/api/contact", GetContactMessages)

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