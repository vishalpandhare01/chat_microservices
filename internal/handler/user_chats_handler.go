package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/repository"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/services"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/utils"
)

// add chat handler
func CreateUserChatHandler(c *fiber.Ctx) error {
	var body *repository.StartChatBody

	userId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	body.UserAId = userId

	response := services.CreateChatServices(body)

	switch r := response.(type) {
	case utils.ErrorResponse:
		return c.Status(r.Code).JSON(fiber.Map{
			"message": r.Message,
		})
	case utils.SuccessResponse:
		return c.Status(r.Code).JSON(fiber.Map{
			"message": r.Message,
			"data":    r.Data,
		})
	default:
		return c.Status(500).JSON(fiber.Map{
			"message": "Somthing wrong in services",
		})

	}

}

// get chat messages by chat id
func GetChatMessagesByChatIdHandler(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	chatId := c.Params("chatId")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	response := services.GetMessagesListServices(pageStr, limitStr, chatId, userId)

	switch r := response.(type) {
	case utils.ErrorResponse:
		return c.Status(r.Code).JSON(fiber.Map{
			"message": r.Message,
		})
	case utils.SuccessResponse:
		return c.Status(r.Code).JSON(fiber.Map{
			"message": r.Message,
			"data":    r.Data,
		})
	default:
		return c.Status(500).JSON(fiber.Map{
			"message": "Somthing wrong in services",
		})

	}

}

// get chat messages by chat id
func GetChatListOfUserHandler(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	response := services.GetChatListServices(userId)

	switch r := response.(type) {
	case utils.ErrorResponse:
		return c.Status(r.Code).JSON(fiber.Map{
			"message": r.Message,
		})
	case utils.SuccessResponse:
		return c.Status(r.Code).JSON(fiber.Map{
			"message": r.Message,
			"data":    r.Data,
		})
	default:
		return c.Status(500).JSON(fiber.Map{
			"message": "Somthing wrong in services",
		})

	}

}

// delete user chat messages by chat id
func DeleteChatOfUserHandler(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}
	chatUserId := c.Params("chatUserId")
	response := services.RemoveUserFromChatListServices(userId, chatUserId)

	switch r := response.(type) {
	case utils.ErrorResponse:
		return c.Status(r.Code).JSON(fiber.Map{
			"message": r.Message,
		})
	case utils.SuccessResponse:
		return c.Status(r.Code).JSON(fiber.Map{
			"message": r.Message,
			"data":    r.Data,
		})
	default:
		return c.Status(500).JSON(fiber.Map{
			"message": "Somthing wrong in services",
		})

	}

}
