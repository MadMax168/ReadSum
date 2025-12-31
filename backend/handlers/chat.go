package handlers

import (
	"github.com/MadMax168/Readsum/config"
	"github.com/MadMax168/Readsum/customerrors"
	"github.com/MadMax168/Readsum/models"
	"github.com/gofiber/fiber/v2"
)

type resp struct {
	Index uint   `json:"index"`
	Title string `json:"title"`
}

func GetChat(c *fiber.Ctx) error {
	UID, ok := c.Locals("userID").(uint)
	if !ok || UID == 0 {
		return customerrors.NewUnauthorizedError("Authentication token is invalid or missing user ID.")
	}

	var cxs []models.Chat
	if err := config.DB.Where("user_id = ?", UID).Find(&cxs).Error; err != nil {
		return customerrors.NewInternalServerError("Not found any chat in sever")
	}

	var response []resp
	for _, i := range cxs {
		response = append(response, resp{
			Index: i.ID,
			Title: i.Title,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    response,
	})
}

func Create(c *fiber.Ctx) error {
	UID, ok := c.Locals("userID").(uint)
	if !ok || UID == 0 {
		return customerrors.NewUnauthorizedError("Authentication token is invalid or missing user ID.")
	}

	chat := new(models.Chat)
	if err := c.BodyParser(&chat); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	chat.UserID = UID
	config.DB.Create(&chat)
	return c.Status(200).JSON(chat)
}

func UpdChat(c *fiber.Ctx) error {
	cx := new(models.Chat)

	UID, ok := c.Locals("userID").(uint)
	if !ok || UID == 0 {
		return customerrors.NewUnauthorizedError("Authentication token is invalid or missing user ID.")
	}

	CID := c.Params("chatID")
	if err := c.BodyParser(cx); err != nil {
		return customerrors.NewBadRequestError("Can not find this chat")
	}

	config.DB.Where("userID = ? and id = ?", UID, CID).Updates(&cx)
	return c.Status(200).JSON(cx)
}

func DelChat(c *fiber.Ctx) error {
	CID := c.Params("chatID")
	var cx models.Chat

	UID, ok := c.Locals("userID").(uint)
	if !ok {
		return c.SendStatus(401)
	}

	result := config.DB.Where("user_id = ? AND id = ?", UID, CID).Delete(&cx)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).SendString("Deleted Success!!")
}
