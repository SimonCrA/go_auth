package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/simoncra/goauth/config"
	"github.com/simoncra/goauth/internal/models"
	"github.com/simoncra/goauth/internal/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// initialize the database
	db, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// auto-migrate the database models
	db.AutoMigrate(&models.User{})

	// create the fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	// set up middlewares
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	// Set up routes
	routes.SetupRoutes(app, db)

	// catch 404 errors
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Ups, estas perdido?  parece que esta ruta no existe!")
	})

	err = app.Listen(":" + config.AppPort)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
