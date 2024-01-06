package main

import (
	"log"

	"github.com/brendanjcarlson/visql/server/src/pkg/config"
	"github.com/brendanjcarlson/visql/server/src/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables from files
	config.MustLoadEnv(".env")

	// Connect to database and defer closing connection
	dbClient := database.MustConnect()
	defer dbClient.MustClose()

	// Configure Fiber app
	appConfig := fiber.Config{
		AppName: config.GetOrDefault("APP_NAME", "visql"),
	}
	// Initialize Fiber app
	app := fiber.New(appConfig)

    // Configure global middleware

    // Bootstrap domains

    // Bootstrap routes

	// Run Fiber app
	log.Fatal(app.Listen(config.GetOrDefault("PORT", ":8080")))
}
