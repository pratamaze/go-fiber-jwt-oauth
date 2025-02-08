package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	"go-fiber-auth/config"
	"go-fiber-auth/routes"
    "go-fiber-auth/migrations"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	config.InitDatabase()
    migrations.AddIsAdminColumn()

    config.SeedAdminUser()


	// Create a new Fiber instance
	app := fiber.New()

	// Register routes
	routes.SetupRoutes(app)

	// Middleware for JWT authentication
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
