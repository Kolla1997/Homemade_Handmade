package main

import "time"

// MenuItem represents a menu item
type MenuItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Category    string `json:"category"`
	Available   bool   `json:"available"`
}

// Order represents a customer order
type Order struct {
	ID                  int       `json:"id"`
	CustomerName        string    `json:"customerName"`
	CustomerPhone       string    `json:"customerPhone"`
	CustomerEmail       string    `json:"customerEmail"`
	CustomerAddress     string    `json:"customerAddress"`
	DeliveryDate        string    `json:"deliveryDate"`
	DeliveryTime        string    `json:"deliveryTime"`
	Items               []int     `json:"items"`
	SpecialInstructions *string   `json:"specialInstructions,omitempty"`
	Subtotal            string    `json:"subtotal"`
	Tax                 string    `json:"tax"`
	DeliveryFee         string    `json:"deliveryFee"`
	Total               string    `json:"total"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"createdAt"`
}

// ContactMessage represents a contact form submission
type ContactMessage struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateOrderRequest for creating an order
type CreateOrderRequest struct {
	CustomerName        string  `json:"customerName" binding:"required"`
	CustomerPhone       string  `json:"customerPhone" binding:"required"`
	CustomerEmail       string  `json:"customerEmail" binding:"required,email"`
	CustomerAddress     string  `json:"customerAddress" binding:"required"`
	DeliveryDate        string  `json:"deliveryDate" binding:"required"`
	DeliveryTime        string  `json:"deliveryTime" binding:"required"`
	Items               []int   `json:"items" binding:"required"`
	SpecialInstructions *string `json:"specialInstructions,omitempty"`
	Subtotal            string  `json:"subtotal" binding:"required"`
	Tax                 string  `json:"tax" binding:"required"`
	DeliveryFee         string  `json:"deliveryFee" binding:"required"`
	Total               string  `json:"total" binding:"required"`
	Status              string  `json:"status"`
}

// CreateContactRequest for creating a contact message
type CreateContactRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}
package main

import "time"

type MenuItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Category    string `json:"category"`
	Image       string `json:"image"`
}

type Order struct {
	ID           int       `json:"id"`
	CustomerName string    `json:"customerName"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Address      string    `json:"address"`
	Date         string    `json:"date"`
	Time         string    `json:"time"`
	Items        string    `json:"items"`
	Subtotal     string    `json:"subtotal"`
	Tax          string    `json:"tax"`
	DeliveryFee  string    `json:"deliveryFee"`
	Total        string    `json:"total"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
}

type ContactMessage struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}
