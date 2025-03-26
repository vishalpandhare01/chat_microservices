package handler

import (
	"log"

	"github.com/gofiber/websocket/v2"
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

	// Get user ID from JWT token
	// token := c.Locals("Authorization").(string)
	// claims, err := validateToken(token)
	// if err != nil || !claims.Valid {
	// 	c.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))
	// 	return
	// }

	// userID := claims.Claims.(jwt.MapClaims)["user_id"].(float64)
	// userID := "1234567890"

	for {
		var msg Message
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		// Insert message into database
		// _, err = db.Exec("INSERT INTO messages (user_id, message) VALUES (?, ?)", userID, msg.Message)
		// if err != nil {
		// 	c.WriteMessage(websocket.TextMessage, []byte("Error saving message"))
		// 	return
		// }

		// Send message to all connected WebSocket clients (broadcast)
		c.WriteMessage(websocket.TextMessage, []byte(msg.Message))
	}
}
