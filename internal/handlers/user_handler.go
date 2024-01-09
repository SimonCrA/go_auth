package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/simoncra/goauth/internal/models"
	"gorm.io/gorm"
)

// getUsersHandler retrieves a lsit of users
func GetUsersHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving users from the database")
		}

		return c.JSON(users)
	}
}

func CreateUserHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newUser models.User

		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		if err := validateUserInput(&newUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if err := db.Create(&newUser).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error creating a new user")
		}

		return c.JSON(newUser)
	}
}

func validateUserInput(user *models.User) error {
	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		var validationErrors []string

		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return fmt.Errorf("validation error: %s", validationErrors)
	}
	return nil
}
