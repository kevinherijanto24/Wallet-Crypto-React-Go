package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"Wallet-Crypto-Crud/backend/database"
	"Wallet-Crypto-Crud/backend/routes"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS middleware for all origins
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", 
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Initialize Database
	database.ConnectDB()

	// Enable logger middleware
	app.Use(logger.New())

	// Register routes
	routes.RegisterWalletRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
