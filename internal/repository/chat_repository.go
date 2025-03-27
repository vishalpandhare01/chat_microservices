package repository

import (
	"github.com/vishalpandhare01/myschool_chat_microservices/initializer"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/model"
)

func CreateChatRepository(body *model.Chats) (*model.Chats, error) {
	if err := initializer.DB.Create(&body).Error; err != nil {
		return nil, err
	}
	return body, nil
}

func DeleteChatRepository(id string) (interface{}, error) {
	var data *model.Chats
	if err := initializer.DB.Delete(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
