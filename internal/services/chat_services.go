package services

import (
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/model"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/repository"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/utils"
)

func CreateChatServices(body *model.Chats) interface{} {
	response, err := repository.CreateChatRepository(body)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}
	return utils.SuccessResponse{
		Code:    201,
		Message: "Chat created successfully",
		Data:    response,
	}

}
