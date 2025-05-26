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
package main

import "time"

type MenuItem struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Available   bool    `json:"available"`
}

type OrderItem struct {
	ItemID   int     `json:"item_id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Order struct {
	ID                   int         `json:"id"`
	CustomerName         string      `json:"customer_name"`
	CustomerPhone        string      `json:"customer_phone"`
	CustomerEmail        string      `json:"customer_email"`
	CustomerAddress      string      `json:"customer_address"`
	DeliveryDate         string      `json:"delivery_date"`
	DeliveryTime         string      `json:"delivery_time"`
	Items                []OrderItem `json:"items"`
	SpecialInstructions  string      `json:"special_instructions"`
	Subtotal             float64     `json:"subtotal"`
	Tax                  float64     `json:"tax"`
	DeliveryFee          float64     `json:"delivery_fee"`
	Total                float64     `json:"total"`
	Status               string      `json:"status"`
	CreatedAt            time.Time   `json:"created_at"`
}

type ContactMessage struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateOrderRequest struct {
	CustomerName         string      `json:"customer_name" binding:"required"`
	CustomerPhone        string      `json:"customer_phone" binding:"required"`
	CustomerEmail        string      `json:"customer_email" binding:"required"`
	CustomerAddress      string      `json:"customer_address" binding:"required"`
	DeliveryDate         string      `json:"delivery_date" binding:"required"`
	DeliveryTime         string      `json:"delivery_time" binding:"required"`
	Items                []OrderItem `json:"items" binding:"required"`
	SpecialInstructions  string      `json:"special_instructions"`
	Subtotal             float64     `json:"subtotal" binding:"required"`
	Tax                  float64     `json:"tax" binding:"required"`
	DeliveryFee          float64     `json:"delivery_fee" binding:"required"`
	Total                float64     `json:"total" binding:"required"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type CreateContactMessageRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}
