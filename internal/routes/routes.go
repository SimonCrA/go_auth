package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/simoncra/goauth/config"
	"github.com/simoncra/goauth/internal/handlers"
	"github.com/simoncra/goauth/internal/middlewares"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	app.Get("/", handlers.HomeHandler)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// declare the middleware to validate the JWT token
	jwt := middlewares.NewAuthMiddleware(config.AppSecret)

	v1.Post("/users", jwt, handlers.CreateUserHandler(db))

	v1.Get("/users", handlers.GetUsersHandler(db))
	v1.Post("/login", handlers.LoginHandler(db))
	// v1.Get("/users/:id")
	// v1.Put("/users/:id")
}
