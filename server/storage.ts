import { 
  users, menuItems, orders, contactMessages,
  type User, type InsertUser,
  type MenuItem, type InsertMenuItem,
  type Order, type InsertOrder,
  type ContactMessage, type InsertContactMessage
} from "@shared/schema";

export interface IStorage {
  // Users
  getUser(id: number): Promise<User | undefined>;
  getUserByUsername(username: string): Promise<User | undefined>;
  createUser(user: InsertUser): Promise<User>;
  
  // Menu Items
  getAllMenuItems(): Promise<MenuItem[]>;
  getMenuItemsByCategory(category: string): Promise<MenuItem[]>;
  createMenuItem(item: InsertMenuItem): Promise<MenuItem>;
  
  // Orders
  createOrder(order: InsertOrder): Promise<Order>;
  getOrder(id: number): Promise<Order | undefined>;
  getAllOrders(): Promise<Order[]>;
  updateOrderStatus(id: number, status: string): Promise<Order | undefined>;
  
  // Contact Messages
  createContactMessage(message: InsertContactMessage): Promise<ContactMessage>;
  getAllContactMessages(): Promise<ContactMessage[]>;
}

export class MemStorage implements IStorage {
  private users: Map<number, User>;
  private menuItems: Map<number, MenuItem>;
  private orders: Map<number, Order>;
  private contactMessages: Map<number, ContactMessage>;
  private currentUserId: number;
  private currentMenuItemId: number;
  private currentOrderId: number;
  private currentContactMessageId: number;

  constructor() {
    this.users = new Map();
    this.menuItems = new Map();
    this.orders = new Map();
    this.contactMessages = new Map();
    this.currentUserId = 1;
    this.currentMenuItemId = 1;
    this.currentOrderId = 1;
    this.currentContactMessageId = 1;
    
    // Initialize with menu items
    this.initializeMenuItems();
  }

  private initializeMenuItems() {
    const menuData: InsertMenuItem[] = [
      // Starters
      { name: "Samosas (2 pieces)", description: "Crispy pastries filled with spiced potatoes and peas", price: "6.99", category: "Starters", available: true },
      { name: "Pakoras", description: "Mixed vegetable fritters with mint chutney", price: "8.99", category: "Starters", available: true },
      { name: "Aloo Tikki", description: "Pan-fried potato patties with tamarind sauce", price: "7.99", category: "Starters", available: true },
      { name: "Chana Chaat", description: "Spiced chickpea salad with yogurt and chutneys", price: "9.99", category: "Starters", available: true },
      
      // Main Courses
      { name: "Butter Chicken", description: "Tender chicken in rich tomato and cream sauce", price: "16.99", category: "Main Courses", available: true },
      { name: "Dal Makhani", description: "Slow-cooked black lentils with butter and cream", price: "14.99", category: "Main Courses", available: true },
      { name: "Palak Paneer", description: "Fresh spinach curry with cottage cheese", price: "15.99", category: "Main Courses", available: true },
      { name: "Chicken Biryani", description: "Aromatic basmati rice with spices and chicken", price: "18.99", category: "Main Courses", available: true },
      { name: "Vegetable Biryani", description: "Aromatic basmati rice with spices and vegetables", price: "16.99", category: "Main Courses", available: true },
      { name: "Rajma", description: "Kidney beans in spiced tomato gravy", price: "13.99", category: "Main Courses", available: true },
      { name: "Chole Bhature", description: "Spiced chickpeas with fluffy fried bread", price: "15.99", category: "Main Courses", available: true },
      
      // Desserts
      { name: "Gulab Jamun (2 pieces)", description: "Soft milk dumplings in rose-scented syrup", price: "5.99", category: "Desserts", available: true },
      { name: "Kheer", description: "Traditional rice pudding with cardamom and nuts", price: "6.99", category: "Desserts", available: true },
      { name: "Kulfi", description: "Dense, creamy Indian ice cream with pistachios", price: "4.99", category: "Desserts", available: true },
      { name: "Rasmalai", description: "Soft cheese patties in sweetened milk", price: "7.99", category: "Desserts", available: true },
    ];

    menuData.forEach(item => {
      const id = this.currentMenuItemId++;
      const menuItem: MenuItem = { ...item, id };
      this.menuItems.set(id, menuItem);
    });
  }

  // Users
  async getUser(id: number): Promise<User | undefined> {
    return this.users.get(id);
  }

  async getUserByUsername(username: string): Promise<User | undefined> {
    return Array.from(this.users.values()).find(user => user.username === username);
  }

  async createUser(insertUser: InsertUser): Promise<User> {
    const id = this.currentUserId++;
    const user: User = { ...insertUser, id };
    this.users.set(id, user);
    return user;
  }

  // Menu Items
  async getAllMenuItems(): Promise<MenuItem[]> {
    return Array.from(this.menuItems.values());
  }

  async getMenuItemsByCategory(category: string): Promise<MenuItem[]> {
    return Array.from(this.menuItems.values()).filter(item => item.category === category);
  }

  async createMenuItem(insertItem: InsertMenuItem): Promise<MenuItem> {
    const id = this.currentMenuItemId++;
    const item: MenuItem = { ...insertItem, id };
    this.menuItems.set(id, item);
    return item;
  }

  // Orders
  async createOrder(insertOrder: InsertOrder): Promise<Order> {
    const id = this.currentOrderId++;
    const order: Order = { 
      ...insertOrder, 
      id, 
      createdAt: new Date()
    };
    this.orders.set(id, order);
    return order;
  }

  async getOrder(id: number): Promise<Order | undefined> {
    return this.orders.get(id);
  }

  async getAllOrders(): Promise<Order[]> {
    return Array.from(this.orders.values());
  }

  async updateOrderStatus(id: number, status: string): Promise<Order | undefined> {
    const order = this.orders.get(id);
    if (order) {
      order.status = status;
      this.orders.set(id, order);
      return order;
    }
    return undefined;
  }

  // Contact Messages
  async createContactMessage(insertMessage: InsertContactMessage): Promise<ContactMessage> {
    const id = this.currentContactMessageId++;
    const message: ContactMessage = { 
      ...insertMessage, 
      id, 
      createdAt: new Date()
    };
    this.contactMessages.set(id, message);
    return message;
  }

  async getAllContactMessages(): Promise<ContactMessage[]> {
    return Array.from(this.contactMessages.values());
  }
}

export const storage = new MemStorage();
