package handler

import (
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/model"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/services"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/utils"
)

// User struct for login and registration
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Message struct for chat
type Message struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

func HandleWebSocket(c *websocket.Conn) {
	defer c.Close()

	// Uncomment and handle JWT Token if needed
	// token := c.Locals("Authorization").(string)
	// claims, err := validateToken(token)
	// if err != nil || !claims.Valid {
	//     c.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))
	//     return
	// }
	// userID := claims.Claims.(jwt.MapClaims)["user_id"].(float64)

	// Mocked userID for the example
	// userID := "1234567890"

	for {
		var msg *model.Messages
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			break // Optionally break the loop if error occurs.
		}

		log.Printf("Received message: %+v\n", msg)

		// Call your service to process the message
		response := services.CreateMessagesServices(msg)

		switch r := response.(type) {
		case utils.ErrorResponse:
			log.Println("Error Response:", r.Message)
			return
		case utils.SuccessResponse:
			log.Println("Success Response:", r.Data)
			c.WriteMessage(websocket.TextMessage, []byte(msg.Message)) // Or msg.Message if you want to echo
			return
		default:
			// Default case if response type is unrecognized
			log.Println("Error:")
			return
		}
	}
}
