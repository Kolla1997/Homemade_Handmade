package main

import (
	"sync"
	"time"
)

type Storage struct {
	mutex           sync.RWMutex
	menuItems       map[int]MenuItem
	orders          []Order
	contactMessages []ContactMessage
	orderCounter    int
	contactCounter  int
}

func NewStorage() *Storage {
	storage := &Storage{
		menuItems:       make(map[int]MenuItem),
		orders:          []Order{},
		contactMessages: []ContactMessage{},
		orderCounter:    1,
		contactCounter:  1,
	}

	// Initialize with sample menu data
	storage.initMenuData()

	return storage
}

func (s *Storage) initMenuData() {
	menuItems := []MenuItem{
		{
			ID:          1,
			Name:        "Butter Chicken",
			Description: "Tender chicken in a rich, creamy tomato-based sauce with aromatic spices",
			Price:       16.99,
			Category:    "Main Courses",
			Image:       "https://images.unsplash.com/photo-1603894584373-5ac82b2ae398?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&h=300",
			Spicy:       true,
			Vegetarian:  false,
			Popular:     true,
		},
		{
			ID:          2,
			Name:        "Palak Paneer",
			Description: "Fresh spinach curry with soft paneer cubes, seasoned with traditional spices",
			Price:       14.99,
			Category:    "Main Courses",
			Image:       "https://images.unsplash.com/photo-1631452180539-96aca7d48617?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&h=300",
			Spicy:       false,
			Vegetarian:  true,
			Popular:     true,
		},
		{
			ID:          3,
			Name:        "Samosas",
			Description: "Crispy pastry triangles filled with spiced potatoes and peas",
			Price:       8.99,
			Category:    "Starters",
			Image:       "https://images.unsplash.com/photo-1601050690597-df0568f70950?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&h=300",
			Spicy:       false,
			Vegetarian:  true,
			Popular:     false,
		},
		{
			ID:          4,
			Name:        "Gulab Jamun",
			Description: "Soft milk dumplings in sweet rose-flavored syrup",
			Price:       6.99,
			Category:    "Desserts",
			Image:       "https://images.unsplash.com/photo-1571019613454-1cb2f99b2d8b?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&h=300",
			Spicy:       false,
			Vegetarian:  true,
			Popular:     true,
		},
		{
			ID:          5,
			Name:        "Biryani",
			Description: "Fragrant basmati rice layered with marinated chicken and aromatic spices",
			Price:       18.99,
			Category:    "Main Courses",
			Image:       "https://images.unsplash.com/photo-1563379091339-03246963d888?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&h=300",
			Spicy:       true,
			Vegetarian:  false,
			Popular:     true,
		},
	}

	for _, item := range menuItems {
		s.menuItems[item.ID] = item
	}
}

func (s *Storage) GetMenuByCategory(category string) []MenuItem {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var items []MenuItem
	for _, item := range s.menuItems {
		if category == "" || item.Category == category {
			items = append(items, item)
		}
	}
	return items
}

func (s *Storage) GetAllMenuItems() []MenuItem {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var items []MenuItem
	for _, item := range s.menuItems {
		items = append(items, item)
	}
	return items
}

func (s *Storage) CreateOrder(order Order) Order {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	order.ID = s.orderCounter
	order.CreatedAt = time.Now()
	order.Status = "Pending"
	s.orders = append(s.orders, order)
	s.orderCounter++

	return order
}

func (s *Storage) GetOrders() []Order {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.orders
}

func (s *Storage) GetOrder(id int) *Order {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, order := range s.orders {
		if order.ID == id {
			return &order
		}
	}
	return nil
}

func (s *Storage) UpdateOrderStatus(id int, status string) *Order {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, order := range s.orders {
		if order.ID == id {
			s.orders[i].Status = status
			return &s.orders[i]
		}
	}
	return nil
}

func (s *Storage) CreateContactMessage(message ContactMessage) ContactMessage {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	message.ID = s.contactCounter
	message.CreatedAt = time.Now()
	s.contactMessages = append(s.contactMessages, message)
	s.contactCounter++

	return message
}

func (s *Storage) GetContactMessages() []ContactMessage {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.contactMessages
}
package main

import (
	"encoding/json"
	"sync"
	"time"
)

type Storage struct {
	menuItems       []MenuItem
	orders          []Order
	contactMessages []ContactMessage
	orderCounter    int
	contactCounter  int
	mutex           sync.RWMutex
}

func NewStorage() *Storage {
	storage := &Storage{
		menuItems:       make([]MenuItem, 0),
		orders:          make([]Order, 0),
		contactMessages: make([]ContactMessage, 0),
		orderCounter:    1,
		contactCounter:  1,
	}
	
	// Initialize with sample menu items
	storage.initializeMenuItems()
	
	return storage
}

func (s *Storage) initializeMenuItems() {
	sampleItems := []MenuItem{
		{ID: 1, Name: "Butter Chicken", Description: "Creamy tomato-based curry with tender chicken", Price: 18.99, Category: "mains", Available: true},
		{ID: 2, Name: "Biryani", Description: "Fragrant basmati rice with spices and meat", Price: 16.99, Category: "mains", Available: true},
		{ID: 3, Name: "Samosas", Description: "Crispy pastries filled with spiced vegetables", Price: 8.99, Category: "appetizers", Available: true},
		{ID: 4, Name: "Naan Bread", Description: "Fresh baked bread from tandoor oven", Price: 4.99, Category: "sides", Available: true},
		{ID: 5, Name: "Gulab Jamun", Description: "Sweet milk dumplings in rose syrup", Price: 6.99, Category: "desserts", Available: true},
		{ID: 6, Name: "Mango Lassi", Description: "Refreshing yogurt drink with mango", Price: 5.99, Category: "beverages", Available: true},
	}
	
	s.menuItems = sampleItems
}

func (s *Storage) GetMenuByCategory(category string) []MenuItem {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	var items []MenuItem
	for _, item := range s.menuItems {
		if item.Category == category && item.Available {
			items = append(items, item)
		}
	}
	return items
}

func (s *Storage) CreateOrder(req CreateOrderRequest) Order {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	order := Order{
		ID:                  s.orderCounter,
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
		Status:              "pending",
		CreatedAt:           time.Now(),
	}
	
	s.orders = append(s.orders, order)
	s.orderCounter++
	
	return order
}

func (s *Storage) GetOrders() []Order {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	return s.orders
}

func (s *Storage) GetOrder(id int) (*Order, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	for _, order := range s.orders {
		if order.ID == id {
			return &order, true
		}
	}
	return nil, false
}

func (s *Storage) UpdateOrderStatus(id int, status string) (*Order, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	for i, order := range s.orders {
		if order.ID == id {
			s.orders[i].Status = status
			return &s.orders[i], true
		}
	}
	return nil, false
}

func (s *Storage) CreateContactMessage(req CreateContactMessageRequest) ContactMessage {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	message := ContactMessage{
		ID:        s.contactCounter,
		Name:      req.Name,
		Email:     req.Email,
		Subject:   req.Subject,
		Message:   req.Message,
		CreatedAt: time.Now(),
	}
	
	s.contactMessages = append(s.contactMessages, message)
	s.contactCounter++
	
	return message
}

func (s *Storage) GetContactMessages() []ContactMessage {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	return s.contactMessages
}
