package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handlers struct holds the storage reference
type Handlers struct {
	storage *Storage
}

// NewHandlers creates a new handlers instance
func NewHandlers(storage *Storage) *Handlers {
	return &Handlers{storage: storage}
}

// GetMenu returns all menu items
func (h *Handlers) GetMenu(c *gin.Context) {
	items := h.storage.GetAllMenuItems()
	c.JSON(http.StatusOK, items)
}

// GetMenuByCategory returns menu items by category
func (h *Handlers) GetMenuByCategory(c *gin.Context) {
	category := c.Param("category")
	items := h.storage.GetMenuItemsByCategory(category)
	c.JSON(http.StatusOK, items)
}

// CreateOrder creates a new order
func (h *Handlers) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid order data",
			"error":   err.Error(),
		})
		return
	}

	order := h.storage.CreateOrder(req)
	c.JSON(http.StatusCreated, order)
}

// GetOrders returns all orders
func (h *Handlers) GetOrders(c *gin.Context) {
	orders := h.storage.GetAllOrders()
	c.JSON(http.StatusOK, orders)
}

// GetOrder returns a specific order by ID
func (h *Handlers) GetOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid order ID",
		})
		return
	}

	order, exists := h.storage.GetOrder(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrderStatus updates the status of an order
func (h *Handlers) UpdateOrderStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid order ID",
		})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Status is required",
		})
		return
	}

	order, exists := h.storage.UpdateOrderStatus(id, req.Status)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

// CreateContactMessage creates a new contact message
func (h *Handlers) CreateContactMessage(c *gin.Context) {
	var req CreateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid contact message data",
			"error":   err.Error(),
		})
		return
	}

	message := h.storage.CreateContactMessage(req)
	c.JSON(http.StatusCreated, message)
}

// GetContactMessages returns all contact messages
func (h *Handlers) GetContactMessages(c *gin.Context) {
	messages := h.storage.GetAllContactMessages()
	c.JSON(http.StatusOK, messages)
}