package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"app/handlers"
	"app/middleware"
	"app/models"
)

var DB *gorm.DB

func main() {
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Web Template v1.0",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, X-CSRF-Token",
		AllowCredentials: true,
	}))

	// Rate limiting
	app.Use(middleware.RateLimiter())

	// CSRF protection
	app.Use(middleware.CSRF())

	// Initialize database (optional - comment out if not using database)
	initDatabase()

	// Serve static files (frontend)
	// In dev, STATIC_PATH points to mounted volume; in prod, to copied files
	staticPath := os.Getenv("STATIC_PATH")
	if staticPath == "" {
		staticPath = "./static"
	}
	app.Static("/", staticPath)

	// API routes
	api := app.Group("/api")

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Example CRUD endpoints
	api.Get("/items", handlers.GetItems)
	api.Post("/items", handlers.CreateItem)
	api.Get("/items/:id", handlers.GetItem)
	api.Put("/items/:id", handlers.UpdateItem)
	api.Delete("/items/:id", handlers.DeleteItem)

	// Contact endpoints (with stricter rate limiting for POST)
	api.Post("/contact", middleware.StrictRateLimiter(), handlers.CreateContact)

	// Admin authentication endpoints
	api.Post("/admin/login", handlers.AdminLogin)
	api.Post("/admin/logout", handlers.AdminLogout)
	api.Get("/admin/check", middleware.AdminAuth(), handlers.CheckAuth)

	// Protected admin endpoints for contact management
	admin := api.Group("/admin", middleware.AdminAuth())
	admin.Get("/contacts", handlers.GetContacts)
	admin.Get("/contacts/:id", handlers.GetContact)
	admin.Put("/contacts/:id", handlers.UpdateContact)
	admin.Delete("/contacts/:id", handlers.DeleteContact)

	// Fallback to index.html for client-side routing
	app.Use(func(c *fiber.Ctx) error {
		return c.SendFile(staticPath + "/index.html")
	})

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

// initDatabase initializes the database connection
// Comment out this function if not using a database
func initDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Println("DATABASE_URL not set, skipping database initialization")
		return
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}

	log.Println("Database connected successfully")

	// Auto-migrate models
	err = DB.AutoMigrate(&models.Item{}, &models.Contact{})
	if err != nil {
		log.Printf("Failed to migrate database: %v", err)
	}

	// Store DB in handlers package
	handlers.DB = DB
}
