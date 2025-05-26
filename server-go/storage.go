package main

import (
	"sync"
	"time"
)

// Storage represents the in-memory storage
type Storage struct {
	mu                       sync.RWMutex
	menuItems                map[int]MenuItem
	orders                   map[int]Order
	contactMessages          map[int]ContactMessage
	currentMenuItemID        int
	currentOrderID           int
	currentContactMessageID  int
}

// NewStorage creates a new storage instance with initialized data
func NewStorage() *Storage {
	storage := &Storage{
		menuItems:               make(map[int]MenuItem),
		orders:                  make(map[int]Order),
		contactMessages:         make(map[int]ContactMessage),
		currentMenuItemID:       1,
		currentOrderID:          1,
		currentContactMessageID: 1,
	}
	
	storage.initializeMenuItems()
	return storage
}

// initializeMenuItems populates the storage with sample menu items
func (s *Storage) initializeMenuItems() {
	menuData := []MenuItem{
		// Starters
		{Name: "Samosas (2 pieces)", Description: "Crispy pastries filled with spiced potatoes and peas", Price: "6.99", Category: "Starters", Available: true},
		{Name: "Pakoras", Description: "Mixed vegetable fritters with mint chutney", Price: "8.99", Category: "Starters", Available: true},
		{Name: "Aloo Tikki", Description: "Pan-fried potato patties with tamarind sauce", Price: "7.99", Category: "Starters", Available: true},
		{Name: "Chana Chaat", Description: "Spiced chickpea salad with yogurt and chutneys", Price: "9.99", Category: "Starters", Available: true},
		
		// Main Courses
		{Name: "Butter Chicken", Description: "Tender chicken in rich tomato and cream sauce", Price: "16.99", Category: "Main Courses", Available: true},
		{Name: "Dal Makhani", Description: "Slow-cooked black lentils with butter and cream", Price: "14.99", Category: "Main Courses", Available: true},
		{Name: "Palak Paneer", Description: "Fresh spinach curry with cottage cheese", Price: "15.99", Category: "Main Courses", Available: true},
		{Name: "Chicken Biryani", Description: "Aromatic basmati rice with spices and chicken", Price: "18.99", Category: "Main Courses", Available: true},
		{Name: "Vegetable Biryani", Description: "Aromatic basmati rice with spices and vegetables", Price: "16.99", Category: "Main Courses", Available: true},
		{Name: "Rajma", Description: "Kidney beans in spiced tomato gravy", Price: "13.99", Category: "Main Courses", Available: true},
		{Name: "Chole Bhature", Description: "Spiced chickpeas with fluffy fried bread", Price: "15.99", Category: "Main Courses", Available: true},
		
		// Desserts
		{Name: "Gulab Jamun (2 pieces)", Description: "Soft milk dumplings in rose-scented syrup", Price: "5.99", Category: "Desserts", Available: true},
		{Name: "Kheer", Description: "Traditional rice pudding with cardamom and nuts", Price: "6.99", Category: "Desserts", Available: true},
		{Name: "Kulfi", Description: "Dense, creamy Indian ice cream with pistachios", Price: "4.99", Category: "Desserts", Available: true},
		{Name: "Rasmalai", Description: "Soft cheese patties in sweetened milk", Price: "7.99", Category: "Desserts", Available: true},
	}

	for _, item := range menuData {
		id := s.currentMenuItemID
		s.currentMenuItemID++
		item.ID = id
		s.menuItems[id] = item
	}
}

// GetAllMenuItems returns all menu items
func (s *Storage) GetAllMenuItems() []MenuItem {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	items := make([]MenuItem, 0, len(s.menuItems))
	for _, item := range s.menuItems {
		items = append(items, item)
	}
	return items
}

// GetMenuItemsByCategory returns menu items filtered by category
func (s *Storage) GetMenuItemsByCategory(category string) []MenuItem {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	var items []MenuItem
	for _, item := range s.menuItems {
		if item.Category == category {
			items = append(items, item)
		}
	}
	return items
}

// CreateOrder creates a new order
func (s *Storage) CreateOrder(req CreateOrderRequest) Order {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	id := s.currentOrderID
	s.currentOrderID++
	
	status := req.Status
	if status == "" {
		status = "pending"
	}
	
	order := Order{
		ID:                  id,
		CustomerName:        req.CustomerName,
		CustomerPhone:       req.CustomerPhone,
		CustomerEmail:       req.CustomerEmail,
		CustomerAddress:     req.CustomerAddress,
		DeliveryDate:        req.DeliveryDate,
		DeliveryTime:        req.DeliveryTime,
		Items:               req.Items,
		SpecialInstructions: req.SpecialInstructions,
		Subtotal:            req.Subtotal,
		Tax:                 req.Tax,
		DeliveryFee:         req.DeliveryFee,
		Total:               req.Total,
		Status:              status,
		CreatedAt:           time.Now(),
	}
	
	s.orders[id] = order
	return order
}

// GetOrder returns an order by ID
func (s *Storage) GetOrder(id int) (Order, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	order, exists := s.orders[id]
	return order, exists
}

// GetAllOrders returns all orders
func (s *Storage) GetAllOrders() []Order {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	orders := make([]Order, 0, len(s.orders))
	for _, order := range s.orders {
		orders = append(orders, order)
	}
	return orders
}

// UpdateOrderStatus updates the status of an order
func (s *Storage) UpdateOrderStatus(id int, status string) (Order, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	order, exists := s.orders[id]
	if !exists {
		return Order{}, false
	}
	
	order.Status = status
	s.orders[id] = order
	return order, true
}

// CreateContactMessage creates a new contact message
func (s *Storage) CreateContactMessage(req CreateContactRequest) ContactMessage {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	id := s.currentContactMessageID
	s.currentContactMessageID++
	
	message := ContactMessage{
		ID:        id,
		Name:      req.Name,
		Email:     req.Email,
		Subject:   req.Subject,
		Message:   req.Message,
		CreatedAt: time.Now(),
	}
	
	s.contactMessages[id] = message
	return message
}

// GetAllContactMessages returns all contact messages
func (s *Storage) GetAllContactMessages() []ContactMessage {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	messages := make([]ContactMessage, 0, len(s.contactMessages))
	for _, message := range s.contactMessages {
		messages = append(messages, message)
	}
	return messages
}