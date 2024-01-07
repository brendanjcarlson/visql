package main

import (
	"log"

	"github.com/brendanjcarlson/visql/server/src/pkg/config"
	"github.com/brendanjcarlson/visql/server/src/pkg/database"
	"github.com/brendanjcarlson/visql/server/src/pkg/domains/account"
	"github.com/brendanjcarlson/visql/server/src/pkg/domains/auth"
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
		AppName: config.MustGet("APP_NAME"),
	}
	// Initialize Fiber app
	app := fiber.New(appConfig)

	// Configure global middleware

	// Bootstrap domains
	// account
	accountRepository := account.NewRepository(dbClient)

	// session
	// sessionRepository := session.NewRepository(dbClient)

	// auth
	authService := auth.NewService(accountRepository)
	authController := auth.NewController(authService)
	auth.SetupRoutes(app, authController)

	// Bootstrap routes

	// Run Fiber app
	log.Fatal(app.Listen(config.MustGet("PORT")))
}
