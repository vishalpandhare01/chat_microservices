package initializer

import (
	"fmt"
	"log"

	"github.com/vishalpandhare01/myschool_chat_microservices/internal/model"
)

func MigrateTables() {
	err := DB.AutoMigrate(
		model.Chats{},
		model.ChatParticipants{},
		model.Messages{},
	)
	if err != nil {
		log.Fatal("Migration failed")
	}
	fmt.Println("Table migrate successfully")
}
