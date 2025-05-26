
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	storage := NewStorage()
	handlers := NewHandlers(storage)

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		api.GET("/menu", handlers.GetMenu)
		api.GET("/menu/:category", handlers.GetMenuByCategory)
		api.POST("/orders", handlers.CreateOrder)
		api.GET("/orders", handlers.GetOrders)
		api.GET("/orders/:id", handlers.GetOrder)
		api.PATCH("/orders/:id/status", handlers.UpdateOrderStatus)
		api.POST("/contact", handlers.CreateContactMessage)
		api.GET("/contact", handlers.GetContactMessages)
	}

	// Serve static files for SPA
	r.Static("/assets", "./dist/assets")
	r.StaticFile("/", "./dist/index.html")
	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	fmt.Printf("🚀 Server starting on port %s\n", port)
	fmt.Printf("📱 Menu API: http://0.0.0.0:%s/api/menu\n", port)
	fmt.Printf("🛒 Orders API: http://0.0.0.0:%s/api/orders\n", port)
	fmt.Printf("📧 Contact API: http://0.0.0.0:%s/api/contact\n", port)

	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
