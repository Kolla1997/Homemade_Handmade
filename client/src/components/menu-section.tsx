import { useQuery } from "@tanstack/react-query";
import type { MenuItem } from "@shared/schema";
import { Card, CardContent } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";

interface MenuSectionProps {
  selectedItems: { [key: number]: number };
  onItemToggle: (itemId: number, checked: boolean) => void;
  onQuantityChange: (itemId: number, quantity: number) => void;
}

export function MenuSection({ selectedItems, onItemToggle, onQuantityChange }: MenuSectionProps) {
  const { data: menuItems, isLoading } = useQuery<MenuItem[]>({
    queryKey: ["/api/menu"],
  });

  if (isLoading) {
    return (
      <div className="py-20 bg-vintage-cream dark:bg-background">
        <div className="container mx-auto px-6">
          <div className="text-center mb-16">
            <Skeleton className="h-12 w-64 mx-auto mb-4" />
            <Skeleton className="h-6 w-96 mx-auto" />
          </div>
          <div className="grid gap-12">
            {[1, 2, 3].map((i) => (
              <Card key={i} className="p-8">
                <Skeleton className="h-8 w-32 mb-6" />
                <div className="grid md:grid-cols-2 gap-6">
                  {[1, 2, 3, 4].map((j) => (
                    <div key={j} className="border-b border-muted pb-4">
                      <Skeleton className="h-6 w-48 mb-2" />
                      <Skeleton className="h-4 w-full" />
                    </div>
                  ))}
                </div>
              </Card>
            ))}
          </div>
        </div>
      </div>
    );
  }

  if (!menuItems || menuItems.length === 0) {
    return (
      <div className="py-20 bg-vintage-cream dark:bg-background">
        <div className="container mx-auto px-6 text-center">
          <h2 className="font-playfair text-4xl font-bold mb-4 text-foreground">Our Authentic Menu</h2>
          <p className="text-muted-foreground">Menu items are currently unavailable. Please check back later.</p>
        </div>
      </div>
    );
  }

  const categories = Array.from(new Set(menuItems.map(item => item.category)));

  const getCategoryIcon = (category: string) => {
    switch (category) {
      case "Starters": return "fas fa-seedling";
      case "Main Courses": return "fas fa-fire";
      case "Desserts": return "fas fa-birthday-cake";
      default: return "fas fa-utensils";
    }
  };

  return (
    <section id="menu" className="py-20 bg-vintage-cream dark:bg-background">
      <div className="container mx-auto px-6">
        <div className="text-center mb-16">
          <h2 className="font-playfair text-4xl md:text-5xl font-bold mb-4 text-foreground">
            Our Authentic Menu
          </h2>
          <p className="text-xl text-muted-foreground max-w-2xl mx-auto">
            Each dish is crafted with traditional spices and cooking methods, bringing you the true taste of Indian home cooking.
          </p>
        </div>

        <div className="grid gap-12">
          {categories.map((category) => {
            const categoryItems = menuItems.filter(item => item.category === category);
            
            return (
              <Card key={category} className="bg-card dark:bg-card border border-border shadow-lg">
                <CardContent className="p-8">
                  <div className="flex items-center mb-6">
                    <i className={`${getCategoryIcon(category)} text-primary text-2xl mr-3`}></i>
                    <h3 className="font-playfair text-3xl font-bold text-foreground">{category}</h3>
                  </div>
                  <div className="grid md:grid-cols-2 gap-6">
                    {categoryItems.map((item) => (
                      <div key={item.id} className="flex justify-between items-start border-b border-border pb-4">
                        <div className="flex-1">
                          <div className="flex items-center gap-3 mb-1">
                            <input
                              type="checkbox"
                              checked={!!selectedItems[item.id]}
                              onChange={(e) => onItemToggle(item.id, e.target.checked)}
                              className="text-primary focus:ring-primary/50"
                            />
                            <h4 className="font-semibold text-foreground">{item.name}</h4>
                          </div>
                          <p className="text-sm text-muted-foreground mb-2">{item.description}</p>
                          <div className="flex items-center justify-between">
                            <span className="text-primary font-bold">${item.price}</span>
                            {selectedItems[item.id] && (
                              <div className="flex items-center gap-2">
                                <button
                                  type="button"
                                  onClick={() => onQuantityChange(item.id, Math.max(0, selectedItems[item.id] - 1))}
                                  className="px-2 py-1 text-foreground border border-border rounded hover:bg-muted"
                                >
                                  -
                                </button>
                                <span className="px-3 text-foreground">{selectedItems[item.id]}</span>
                                <button
                                  type="button"
                                  onClick={() => onQuantityChange(item.id, selectedItems[item.id] + 1)}
                                  className="px-2 py-1 text-foreground border border-border rounded hover:bg-muted"
                                >
                                  +
                                </button>
                              </div>
                            )}
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                </CardContent>
              </Card>
            );
          })}
        </div>
      </div>
    </section>
  );
}
