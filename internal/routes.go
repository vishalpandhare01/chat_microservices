package internal

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/handler"
)

func SetupRoutes(app *fiber.App) {

	// WebSocket route
	app.Get("/api/v1/chat", websocket.New(handler.HandleWebSocket))
	fmt.Println("All routes connected 🚀")
}
