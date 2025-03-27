package repository

import (
	"fmt"

	"github.com/vishalpandhare01/myschool_chat_microservices/initializer"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/model"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/utils"
	funcation "github.com/vishalpandhare01/myschool_chat_microservices/internal/utils/function"
	"gorm.io/gorm"
)

type StartChatBody struct {
	UserAId  string
	UserBId  string
	ChatName string
	IsGroup  bool
}

// 2 user chat exist check
func CheckChatExistRepository(body *StartChatBody) (*model.ChatParticipants, error) {
	var chatsParticipants model.ChatParticipants

	// Find a chat where both UserAId and UserBId are participants in the same chat
	err := initializer.DB.
		Table("chat_participants").
		Where("chat_id IN (?)",
			initializer.DB.
				Table("chat_participants").
				Select("chat_id").
				Where("user_id = ?", body.UserAId)).
		Where("chat_id IN (?)",
			initializer.DB.
				Table("chat_participants").
				Select("chat_id").
				Where("user_id = ?", body.UserBId)).
		First(&chatsParticipants).Error

	if err != nil {
		// If no result found, return nil or any custom error handling
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No chat exists between the two users
		}
		return nil, err // Return any other error
	}

	// Return the found chat participants (this can be used to return chat details)
	return &chatsParticipants, nil
}

// check chat id and sender id exist or not
func CheckChatIdWithSendorIdExistRepository(chatId string, sendorId string) bool {
	var chatsParticipants model.ChatParticipants

	if err := initializer.DB.Where("user_id = ? AND chat_id = ?", sendorId, chatId).First(&chatsParticipants).Error; err != nil {
		return false
	}

	return true
}

func CreateChatRepository(body *StartChatBody) (*model.Chats, error) {
	var chats model.Chats

	// Start a transaction
	tx := initializer.DB.Begin()

	// Rollback in case of any failure
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	chats.ChatName = body.ChatName

	// Create the chat within the transaction
	if err := tx.Create(&chats).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Prepare participants data
	participants := []model.ChatParticipants{
		{ChatID: chats.ID, UserID: body.UserAId}, // User A
		{ChatID: chats.ID, UserID: body.UserBId}, // User B
	}

	// Insert participants into the chat_participants table within the transaction
	if err := tx.Create(&participants).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Return the created chat data
	return &chats, nil
}

func DeleteChatRepository(id string) (interface{}, error) {
	var data *model.Chats
	if err := initializer.DB.Delete(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func CreateMessagesRepository(body *model.Messages) (*model.Messages, error) {
	if err := initializer.DB.Create(&body).Error; err != nil {
		return nil, err
	}
	return body, nil
}

func GetMessagesListRepository(pageStr string, limitStr string, chatId string) (interface{}, error) {
	var data []model.Messages
	var totalData []model.Messages

	offset, limitInt := funcation.Pagination(pageStr, limitStr)
	fmt.Println("limitInt", limitInt, "offset", offset)

	if err := initializer.DB.Where("chat_id = ?", chatId).Find(&totalData).Error; err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	// offset = (offset - 1) * limitInt
	initializer.DB.Where("chat_id = ?", chatId).Limit(limitInt).Offset(offset).Order("id DESC")
	if err := initializer.DB.Where("chat_id = ?", chatId).Find(&data).Order("created_at").Error; err != nil {
		return nil, err
	}

	responseData := utils.SuccessListResponse{
		Total:   len(totalData),
		Perpage: limitInt,
		Page:    offset,
		Data:    data,
	}

	return responseData, nil
}
