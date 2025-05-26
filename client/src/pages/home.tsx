import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { useTheme } from "@/components/theme-provider";
import { OrderForm } from "@/components/order-form";
import { ContactForm } from "@/components/contact-form";

export default function Home() {
  const { theme, setTheme } = useTheme();

  const scrollToSection = (sectionId: string) => {
    const element = document.getElementById(sectionId);
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' });
    }
  };

  return (
    <div className="min-h-screen bg-background text-foreground">
      {/* Header */}
      <header className="sticky top-0 z-50 bg-background/95 dark:bg-background/95 backdrop-blur-sm border-b border-border">
        <nav className="container mx-auto px-6 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-2">
              <i className="fas fa-utensils text-primary text-2xl"></i>
              <h1 className="font-playfair text-2xl font-bold text-foreground">Spice Heritage</h1>
            </div>
            
            <div className="hidden md:flex items-center space-x-8">
              <button onClick={() => scrollToSection('home')} className="hover:text-primary transition-colors">Home</button>
              <button onClick={() => scrollToSection('about')} className="hover:text-primary transition-colors">About</button>
              <button onClick={() => scrollToSection('menu')} className="hover:text-primary transition-colors">Menu</button>
              <button onClick={() => scrollToSection('order')} className="hover:text-primary transition-colors">Order</button>
              <button onClick={() => scrollToSection('contact')} className="hover:text-primary transition-colors">Contact</button>
            </div>

            <div className="flex items-center space-x-4">
              <Button
                variant="ghost"
                size="icon"
                onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
                className="bg-muted/20 hover:bg-muted/30"
              >
                <i className={`fas ${theme === "dark" ? "fa-sun" : "fa-moon"} text-foreground`}></i>
              </Button>
              
              <Button variant="ghost" size="icon" className="md:hidden bg-muted/20">
                <i className="fas fa-bars text-foreground"></i>
              </Button>
            </div>
          </div>
        </nav>
      </header>

      {/* Hero Section */}
      <section id="home" className="relative min-h-screen flex items-center justify-center overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-br from-muted/20 to-primary/20 dark:from-muted/30 dark:to-primary/30"></div>
        <img 
          src="https://images.unsplash.com/photo-1596040033229-a9821ebd058d?ixlib=rb-4.0.3&auto=format&fit=crop&w=1920&h=1080" 
          alt="Traditional Indian spices and cooking ingredients" 
          className="absolute inset-0 w-full h-full object-cover opacity-30 dark:opacity-20" 
        />
        
        <div className="relative z-10 text-center px-6 max-w-4xl mx-auto">
          <h1 className="font-playfair text-5xl md:text-7xl font-bold mb-6 text-foreground leading-tight">
            Taste Authentic<br />
            <span className="text-primary">Indian Home Cooking</span>
          </h1>
          <p className="text-xl md:text-2xl mb-8 text-muted-foreground max-w-2xl mx-auto font-light">
            Experience the warmth of traditional recipes passed down through generations, 
            now delivered fresh to your doorstep.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Button 
              onClick={() => scrollToSection('order')}
              size="lg"
              className="bg-primary hover:bg-primary/90 text-primary-foreground px-8 py-4 text-lg font-semibold transform hover:scale-105 transition-all"
            >
              <i className="fas fa-shopping-cart mr-2"></i>
              Order Now
            </Button>
            <Button 
              onClick={() => scrollToSection('menu')}
              variant="outline"
              size="lg"
              className="border-2 border-foreground text-foreground hover:bg-foreground hover:text-background px-8 py-4 text-lg font-semibold transition-all"
            >
              View Menu
            </Button>
          </div>
        </div>
      </section>

      {/* About Section */}
      <section id="about" className="py-20 bg-card dark:bg-card">
        <div className="container mx-auto px-6">
          <div className="grid md:grid-cols-2 gap-12 items-center">
            <div className="order-2 md:order-1">
              <h2 className="font-playfair text-4xl font-bold mb-6 text-foreground">
                Our Heritage Story
              </h2>
              <div className="space-y-4 text-muted-foreground leading-relaxed">
                <p>
                  Founded by Priya Sharma in 2018, Spice Heritage began as a way to share the authentic flavors 
                  of her grandmother's kitchen with the local community. Growing up in Mumbai, Priya learned the 
                  art of traditional Indian cooking from her grandmother, who emphasized the importance of using 
                  fresh spices and time-honored techniques.
                </p>
                <p>
                  After moving to the United States, Priya noticed a lack of truly authentic Indian home-style 
                  cooking. Most restaurants focused on popular dishes, but missed the soul of everyday Indian 
                  meals that families enjoy at home. This inspired her to start Spice Heritage, bringing the 
                  warmth and authenticity of Indian home cooking to your table.
                </p>
                <p>
                  Every dish is prepared with the same care and attention that Priya's grandmother taught her, 
                  using traditional recipes and the finest ingredients sourced both locally and from trusted 
                  suppliers in India.
                </p>
              </div>
              <Card className="mt-8 bg-muted/10 dark:bg-muted/20 border border-muted">
                <CardContent className="p-6">
                  <div className="flex items-center mb-3">
                    <i className="fas fa-clock text-primary mr-3"></i>
                    <h3 className="font-semibold text-foreground">Operating Hours</h3>
                  </div>
                  <p className="text-muted-foreground">
                    Monday - Saturday: 10:00 AM - 8:00 PM<br />
                    Sunday: Closed
                  </p>
                </CardContent>
              </Card>
            </div>
            
            <div className="order-1 md:order-2">
              <div className="relative">
                <img 
                  src="https://pixabay.com/get/g2b8faec52e6607ffab916af0171527c21e4a72572ee34dfc848b429bff9de6555fb744eeae6260ec0b5311834e56a7ebd304b33f4cfabbe3146b56900774c8af_1280.jpg" 
                  alt="Indian woman cooking in traditional kitchen" 
                  className="rounded-2xl shadow-xl w-full object-cover h-96 md:h-full border-4 border-muted" 
                />
                <div className="absolute -bottom-6 -right-6 w-24 h-24 bg-primary rounded-full flex items-center justify-center shadow-lg">
                  <i className="fas fa-heart text-primary-foreground text-2xl"></i>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Order Form (includes Menu Section) */}
      <OrderForm />

      {/* Contact Form */}
      <ContactForm />

      {/* Footer */}
      <footer className="bg-foreground dark:bg-muted text-background dark:text-foreground py-12">
        <div className="container mx-auto px-6">
          <div className="grid md:grid-cols-4 gap-8 mb-8">
            <div>
              <div className="flex items-center space-x-2 mb-4">
                <i className="fas fa-utensils text-primary text-2xl"></i>
                <h3 className="font-playfair text-2xl font-bold">Spice Heritage</h3>
              </div>
              <p className="text-background/80 dark:text-foreground/80 leading-relaxed">
                Bringing authentic Indian home cooking to your doorstep with love, tradition, and the finest ingredients.
              </p>
            </div>
            
            <div>
              <h4 className="font-semibold text-lg mb-4">Quick Links</h4>
              <ul className="space-y-2 text-background/80 dark:text-foreground/80">
                <li><button onClick={() => scrollToSection('home')} className="hover:text-primary transition-colors">Home</button></li>
                <li><button onClick={() => scrollToSection('about')} className="hover:text-primary transition-colors">About</button></li>
                <li><button onClick={() => scrollToSection('menu')} className="hover:text-primary transition-colors">Menu</button></li>
                <li><button onClick={() => scrollToSection('order')} className="hover:text-primary transition-colors">Order</button></li>
                <li><button onClick={() => scrollToSection('contact')} className="hover:text-primary transition-colors">Contact</button></li>
              </ul>
            </div>
            
            <div>
              <h4 className="font-semibold text-lg mb-4">Services</h4>
              <ul className="space-y-2 text-background/80 dark:text-foreground/80">
                <li>Home Delivery</li>
                <li>Catering Services</li>
                <li>Custom Orders</li>
                <li>Party Platters</li>
                <li>Gift Cards</li>
              </ul>
            </div>
            
            <div>
              <h4 className="font-semibold text-lg mb-4">Contact Info</h4>
              <div className="space-y-2 text-background/80 dark:text-foreground/80">
                <p>(555) 123-SPICE</p>
                <p>hello@spiceheritage.com</p>
                <p>123 Spice Street<br />Food City, FC 12345</p>
              </div>
            </div>
          </div>
          
          <div className="border-t border-background/20 dark:border-foreground/20 pt-8 flex flex-col md:flex-row justify-between items-center">
            <p className="text-background/60 dark:text-foreground/60 text-sm">
              © 2024 Spice Heritage. All rights reserved.
            </p>
            <div className="flex space-x-4 mt-4 md:mt-0">
              <a href="#" className="text-background/60 dark:text-foreground/60 hover:text-primary transition-colors">
                Privacy Policy
              </a>
              <a href="#" className="text-background/60 dark:text-foreground/60 hover:text-primary transition-colors">
                Terms of Service
              </a>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}
