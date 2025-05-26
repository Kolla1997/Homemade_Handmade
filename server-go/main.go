package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize storage
	storage := NewStorage()
	handlers := NewHandlers(storage)

	// Create Gin router
	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// API routes
	api := r.Group("/api")
	{
		// Menu routes
		api.GET("/menu", handlers.GetMenu)
		api.GET("/menu/:category", handlers.GetMenuByCategory)

		// Order routes
		api.POST("/orders", handlers.CreateOrder)
		api.GET("/orders", handlers.GetOrders)
		api.GET("/orders/:id", handlers.GetOrder)
		api.PATCH("/orders/:id/status", handlers.UpdateOrderStatus)

		// Contact routes
		api.POST("/contact", handlers.CreateContactMessage)
		api.GET("/contact", handlers.GetContactMessages)
	}

	// Serve static files (for production)
	r.Static("/assets", "./dist/assets")
	r.StaticFile("/", "./dist/index.html")
	
	// Fallback to index.html for SPA routing
	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	// Get port from environment or default to 5000
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	fmt.Printf("🚀 Server starting on port %s\n", port)
	fmt.Printf("📱 Menu API: http://localhost:%s/api/menu\n", port)
	fmt.Printf("🛒 Orders API: http://localhost:%s/api/orders\n", port)
	fmt.Printf("📧 Contact API: http://localhost:%s/api/contact\n", port)

	// Start server
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}