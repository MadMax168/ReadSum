package handlers

import (
	"errors"
	"strings"

	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/MadMax168/Readsum/config"
	"github.com/MadMax168/Readsum/models"
	customerrors "github.com/MadMax168/Readsum/errors"
)

func Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return customerrors.NewBadRequestError("Invalid request body")
	}

	input.Name = strings.TrimSpace(strings.ToLower(input.Name))
	input.Email = strings.TrimSpace(strings.ToLower(input.Email))

	if input.Name == "" || input.Email == "" || input.Password == "" {
		return customerrors.NewBadRequestError("Name, email, and password are required")
	}

	if len(input.Password) < 6 {
		return customerrors.NewBadRequestError("Password must be at least 6 characters")
	}

	var existingUser models.User

	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return customerrors.NewConflictError("Email already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return customerrors.NewInternalServerError("Database error")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return customerrors.NewInternalServerError("Failed to process password")
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return customerrors.NewInternalServerError("Failed to create user")
	}

	//DTO
	type UserResponse struct {
		ID    uint    `json:"id"`
		Name  string  `json:"name"`
		Email string `json:"email"`
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"data": UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		"message": "User registered successfully",
	})
}