import { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { z } from "zod";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { useToast } from "@/hooks/use-toast";
import { apiRequest } from "@/lib/queryClient";
import type { MenuItem } from "@shared/schema";
import { MenuSection } from "./menu-section";

const orderFormSchema = z.object({
  customerName: z.string().min(2, "Name must be at least 2 characters"),
  customerPhone: z.string().min(10, "Phone number is required"),
  customerEmail: z.string().email("Valid email is required"),
  customerAddress: z.string().min(10, "Address is required"),
  deliveryDate: z.string().min(1, "Delivery date is required"),
  deliveryTime: z.string().min(1, "Delivery time is required"),
  specialInstructions: z.string().optional(),
});

type OrderFormData = z.infer<typeof orderFormSchema>;

export function OrderForm() {
  const [selectedItems, setSelectedItems] = useState<{ [key: number]: number }>({});
  const { toast } = useToast();
  const queryClient = useQueryClient();

  const { data: menuItems } = useQuery<MenuItem[]>({
    queryKey: ["/api/menu"],
  });

  const form = useForm<OrderFormData>({
    resolver: zodResolver(orderFormSchema),
    defaultValues: {
      customerName: "",
      customerPhone: "",
      customerEmail: "",
      customerAddress: "",
      deliveryDate: "",
      deliveryTime: "",
      specialInstructions: "",
    },
  });

  const createOrderMutation = useMutation({
    mutationFn: async (orderData: any) => {
      const response = await apiRequest("POST", "/api/orders", orderData);
      return response.json();
    },
    onSuccess: () => {
      toast({
        title: "Order Confirmed!",
        description: "Your order has been successfully placed. You will receive a confirmation email shortly.",
      });
      form.reset();
      setSelectedItems({});
      queryClient.invalidateQueries({ queryKey: ["/api/orders"] });
    },
    onError: () => {
      toast({
        title: "Order Failed",
        description: "There was an error placing your order. Please try again.",
        variant: "destructive",
      });
    },
  });

  const handleItemToggle = (itemId: number, checked: boolean) => {
    setSelectedItems(prev => {
      if (checked) {
        return { ...prev, [itemId]: 1 };
      } else {
        const newItems = { ...prev };
        delete newItems[itemId];
        return newItems;
      }
    });
  };

  const handleQuantityChange = (itemId: number, quantity: number) => {
    if (quantity <= 0) {
      setSelectedItems(prev => {
        const newItems = { ...prev };
        delete newItems[itemId];
        return newItems;
      });
    } else {
      setSelectedItems(prev => ({ ...prev, [itemId]: quantity }));
    }
  };

  const calculateOrderTotals = () => {
    if (!menuItems) return { subtotal: 0, tax: 0, deliveryFee: 3.99, total: 3.99 };

    const subtotal = Object.entries(selectedItems).reduce((sum, [itemId, quantity]) => {
      const item = menuItems.find(item => item.id === parseInt(itemId));
      if (item) {
        return sum + (parseFloat(item.price) * quantity);
      }
      return sum;
    }, 0);

    const tax = subtotal * 0.08; // 8% tax
    const deliveryFee = 3.99;
    const total = subtotal + tax + deliveryFee;

    return { subtotal, tax, deliveryFee, total };
  };

  const { subtotal, tax, deliveryFee, total } = calculateOrderTotals();

  const onSubmit = (data: OrderFormData) => {
    if (Object.keys(selectedItems).length === 0) {
      toast({
        title: "No Items Selected",
        description: "Please select at least one item to order.",
        variant: "destructive",
      });
      return;
    }

    const orderItems = Object.entries(selectedItems).map(([itemId, quantity]) => ({
      itemId: parseInt(itemId),
      quantity,
    }));

    const orderData = {
      ...data,
      items: JSON.stringify(orderItems),
      subtotal: subtotal.toString(),
      tax: tax.toString(),
      deliveryFee: deliveryFee.toString(),
      total: total.toString(),
      status: "pending",
    };

    createOrderMutation.mutate(orderData);
  };

  // Generate time slots
  const timeSlots = [
    "11:00", "11:30", "12:00", "12:30", "13:00", "13:30",
    "18:00", "18:30", "19:00", "19:30"
  ];

  const formatTimeSlot = (time: string) => {
    const [hour, minute] = time.split(":");
    const hourNum = parseInt(hour);
    const isPM = hourNum >= 12;
    const displayHour = hourNum > 12 ? hourNum - 12 : hourNum === 0 ? 12 : hourNum;
    const nextHour = hourNum + 1 > 12 ? (hourNum + 1) - 12 : hourNum + 1;
    const nextPeriod = hourNum + 1 >= 12 && hourNum < 12 ? "PM" : isPM ? "PM" : "AM";
    
    return `${displayHour}:${minute} ${isPM ? "PM" : "AM"} - ${nextHour}:${minute} ${nextPeriod}`;
  };

  // Set minimum date to today
  const today = new Date().toISOString().split('T')[0];

  return (
    <>
      <MenuSection 
        selectedItems={selectedItems}
        onItemToggle={handleItemToggle}
        onQuantityChange={handleQuantityChange}
      />
      
      <section id="order" className="py-20 bg-card dark:bg-card">
        <div className="container mx-auto px-6">
          <div className="text-center mb-16">
            <h2 className="font-playfair text-4xl md:text-5xl font-bold mb-4 text-foreground">
              Place Your Order
            </h2>
            <p className="text-xl text-muted-foreground max-w-2xl mx-auto">
              Select your favorite dishes and choose a convenient delivery time. No phone calls or emails needed!
            </p>
          </div>

          <div className="max-w-4xl mx-auto">
            <Card className="bg-muted/50 dark:bg-muted/20 border border-border">
              <CardContent className="p-8">
                <Form {...form}>
                  <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    {/* Customer Information */}
                    <div className="grid md:grid-cols-2 gap-6">
                      <FormField
                        control={form.control}
                        name="customerName"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-foreground">Full Name *</FormLabel>
                            <FormControl>
                              <Input placeholder="Enter your full name" {...field} />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                      
                      <FormField
                        control={form.control}
                        name="customerPhone"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-foreground">Phone Number *</FormLabel>
                            <FormControl>
                              <Input placeholder="(555) 123-4567" {...field} />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                      
                      <FormField
                        control={form.control}
                        name="customerEmail"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-foreground">Email Address *</FormLabel>
                            <FormControl>
                              <Input type="email" placeholder="your.email@example.com" {...field} />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                      
                      <FormField
                        control={form.control}
                        name="customerAddress"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-foreground">Delivery Address *</FormLabel>
                            <FormControl>
                              <Input placeholder="Street address" {...field} />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>

                    {/* Delivery Date and Time */}
                    <div className="grid md:grid-cols-2 gap-6">
                      <FormField
                        control={form.control}
                        name="deliveryDate"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-foreground">Delivery Date *</FormLabel>
                            <FormControl>
                              <Input type="date" min={today} {...field} />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                      
                      <FormField
                        control={form.control}
                        name="deliveryTime"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel className="text-foreground">Delivery Time *</FormLabel>
                            <Select onValueChange={field.onChange} defaultValue={field.value}>
                              <FormControl>
                                <SelectTrigger>
                                  <SelectValue placeholder="Select delivery time" />
                                </SelectTrigger>
                              </FormControl>
                              <SelectContent>
                                {timeSlots.map((time) => (
                                  <SelectItem key={time} value={time}>
                                    {formatTimeSlot(time)}
                                  </SelectItem>
                                ))}
                              </SelectContent>
                            </Select>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>

                    {/* Special Instructions */}
                    <FormField
                      control={form.control}
                      name="specialInstructions"
                      render={({ field }) => (
                        <FormItem>
                          <FormLabel className="text-foreground">Special Instructions (Optional)</FormLabel>
                          <FormControl>
                            <Textarea 
                              placeholder="Any special dietary requirements or cooking preferences..." 
                              rows={3}
                              {...field} 
                            />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />

                    {/* Order Summary */}
                    <Card className="bg-card dark:bg-card border border-border">
                      <CardContent className="p-6">
                        <h3 className="text-lg font-semibold text-foreground mb-4">Order Summary</h3>
                        <div className="space-y-2 text-muted-foreground">
                          <div className="flex justify-between">
                            <span>Subtotal:</span>
                            <span>${subtotal.toFixed(2)}</span>
                          </div>
                          <div className="flex justify-between">
                            <span>Delivery Fee:</span>
                            <span>${deliveryFee.toFixed(2)}</span>
                          </div>
                          <div className="flex justify-between">
                            <span>Tax:</span>
                            <span>${tax.toFixed(2)}</span>
                          </div>
                          <hr className="border-border" />
                          <div className="flex justify-between font-bold text-foreground">
                            <span>Total:</span>
                            <span>${total.toFixed(2)}</span>
                          </div>
                        </div>
                      </CardContent>
                    </Card>

                    {/* Submit Button */}
                    <Button 
                      type="submit" 
                      disabled={createOrderMutation.isPending}
                      className="w-full bg-primary hover:bg-primary/90 text-primary-foreground py-4 text-lg font-semibold"
                    >
                      {createOrderMutation.isPending ? (
                        <>
                          <i className="fas fa-spinner fa-spin mr-2"></i>
                          Processing Order...
                        </>
                      ) : (
                        <>
                          <i className="fas fa-check-circle mr-2"></i>
                          Confirm Order & Schedule Delivery
                        </>
                      )}
                    </Button>
                  </form>
                </Form>
              </CardContent>
            </Card>
          </div>
        </div>
      </section>
    </>
  );
}
