package main

import "time"

type MenuItem struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Spicy       bool    `json:"spicy"`
	Vegetarian  bool    `json:"vegetarian"`
	Popular     bool    `json:"popular"`
}

type OrderItem struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

type Order struct {
	ID                  int         `json:"id"`
	CustomerName        string      `json:"customerName"`
	CustomerPhone       string      `json:"customerPhone"`
	CustomerEmail       string      `json:"customerEmail"`
	CustomerAddress     string      `json:"customerAddress"`
	DeliveryDate        string      `json:"deliveryDate"`
	DeliveryTime        string      `json:"deliveryTime"`
	Items               []OrderItem `json:"items"`
	SpecialInstructions string      `json:"specialInstructions"`
	Subtotal            float64     `json:"subtotal"`
	Tax                 float64     `json:"tax"`
	DeliveryFee         float64     `json:"deliveryFee"`
	Total               float64     `json:"total"`
	Status              string      `json:"status"`
	CreatedAt           time.Time   `json:"createdAt"`
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