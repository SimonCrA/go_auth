package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simoncra/goauth/internal/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/", handlers.HomeHandler)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/users", handlers.CreateUserHandler(db))
	v1.Get("/users", handlers.GetUsersHandler(db))
	// v1.Get("/users/:id")
	// v1.Put("/users/:id")
}
