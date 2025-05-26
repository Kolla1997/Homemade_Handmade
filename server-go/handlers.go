package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	storage *Storage
}

func NewHandlers(storage *Storage) *Handlers {
	return &Handlers{storage: storage}
}

func (h *Handlers) GetMenuByCategory(c *gin.Context) {
	category := c.Param("category")
	items := h.storage.GetMenuByCategory(category)
	c.JSON(http.StatusOK, items)
}

func (h *Handlers) CreateOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdOrder := h.storage.CreateOrder(order)
	c.JSON(http.StatusCreated, createdOrder)
}

func (h *Handlers) GetOrders(c *gin.Context) {
	orders := h.storage.GetOrders()
	c.JSON(http.StatusOK, orders)
}

func (h *Handlers) GetOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order := h.storage.GetOrder(id)
	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handlers) UpdateOrderStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var statusUpdate struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := h.storage.UpdateOrderStatus(id, statusUpdate.Status)
	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handlers) CreateContactMessage(c *gin.Context) {
	var message ContactMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMessage := h.storage.CreateContactMessage(message)
	c.JSON(http.StatusCreated, createdMessage)
}

func (h *Handlers) GetContactMessages(c *gin.Context) {
	messages := h.storage.GetContactMessages()
	c.JSON(http.StatusOK, messages)
}
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	storage *Storage
}

func NewHandlers(storage *Storage) *Handlers {
	return &Handlers{storage: storage}
}

func (h *Handlers) GetMenuByCategory(c *gin.Context) {
	category := c.Param("category")
	
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
		return
	}
	
	items := h.storage.GetMenuByCategory(category)
	c.JSON(http.StatusOK, items)
}

func (h *Handlers) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	order := h.storage.CreateOrder(req)
	c.JSON(http.StatusCreated, order)
}

func (h *Handlers) GetOrders(c *gin.Context) {
	orders := h.storage.GetOrders()
	c.JSON(http.StatusOK, orders)
}

func (h *Handlers) GetOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}
	
	order, found := h.storage.GetOrder(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	
	c.JSON(http.StatusOK, order)
}

func (h *Handlers) UpdateOrderStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}
	
	var req UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	order, found := h.storage.UpdateOrderStatus(id, req.Status)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	
	c.JSON(http.StatusOK, order)
}

func (h *Handlers) CreateContactMessage(c *gin.Context) {
	var req CreateContactMessageRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	message := h.storage.CreateContactMessage(req)
	c.JSON(http.StatusCreated, message)
}

func (h *Handlers) GetContactMessages(c *gin.Context) {
	messages := h.storage.GetContactMessages()
	c.JSON(http.StatusOK, messages)
}
