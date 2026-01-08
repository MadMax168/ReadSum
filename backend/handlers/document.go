package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/MadMax168/Readsum/models"
	"github.com/MadMax168/Readsum/config"
	"github.com/MadMax168/Readsum/customerrors"
)

func CreateDocument(c* fiber.Ctx) error {
	UID := c.Locals("userID").(uint)

	input := new(models.Document)
	if err := c.BodyParser(input); err != nil {
		return customerrors.NewBadRequestError("Invalid request body")
	}

	if input.FileUrl == "" {
		return customerrors.NewBadRequestError("File URL is required")
	}

	input.UserID = UID
	input.Status = "process"
	input.Summary = "{}"
	if err := config.DB.Create(input).Error; err != nil {
		fmt.Println("GORM Error:", err)
		return customerrors.NewInternalServerError("Failed to save data")
	}

	//sent to AI summarize
	return c.Status(201).JSON(input)
}

func GetDocuments(c *fiber.Ctx) error {
	UID := c.Locals("userID").(uint)
	var docs []models.Document

	if err := config.DB.Where("user_id = ?", UID).Find(&docs).Error; err != nil {
		return customerrors.NewInternalServerError("Could not fetch documents")
	}

	return c.JSON(docs)
}

func GetDocumentByID(c *fiber.Ctx) error {
	UID := c.Locals("userID").(uint)
	DID := c.Params("id")

	var doc models.Document

	if err := config.DB.Where("id = ? AND user_id = ?", DID, UID).First(&doc).Error; err != nil {
		fmt.Println("Error: ", err)
		return customerrors.NewNotFoundError("Document not found")
	}

	return c.JSON(doc)
}

func DeleteDocument(c *fiber.Ctx) error {
	UID := c.Locals("userID").(uint)
	DID := c.Params("id")

	result := config.DB.Where("id = ? AND user_id = ?", DID, UID).Delete(&models.Document{})

	if result.RowsAffected == 0 {
		return customerrors.NewNotFoundError("Document not found")
	}

	if result.Error != nil {
		return customerrors.NewInternalServerError("Failed to delete document")
	}

	return c.Status(200).JSON(fiber.Map{ "message": "Deleted document"})
}