package internal

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/handler"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/middleware"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/api/v1/start-chat", middleware.Authentication, handler.CreateUserChatHandler)
	app.Get("/api/v1/chat/:chatId", middleware.Authentication, handler.GetChatMessagesByChatIdHandler)

	// WebSocket route
	app.Get("/api/v1/chat", websocket.New(handler.HandleWebSocket))
	fmt.Println("All routes connected 🚀")
}
