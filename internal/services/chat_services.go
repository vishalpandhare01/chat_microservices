package services

import (
	"fmt"

	"github.com/vishalpandhare01/myschool_chat_microservices/internal/model"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/repository"
	"github.com/vishalpandhare01/myschool_chat_microservices/internal/utils"
)

// create chats beetween 2 users
func CreateChatServices(body *repository.StartChatBody) interface{} {
	//todo check users exist by user1 , user 2 microservices

	response, err := repository.CheckChatExistRepository(body)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	if response != nil {
		return utils.SuccessResponse{
			Code:    200,
			Message: "Chat Already Exist",
			Data:    response,
		}
	}
	responseData, err := repository.CreateChatRepository(body)

	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}
	return utils.SuccessResponse{
		Code:    201,
		Message: "Chat created successfully",
		Data:    responseData,
	}

}

// create messages services
func CreateMessagesServices(body *model.Messages) interface{} {
	//todo sendor id exist or not check by user microservices

	fmt.Println("Body: ", body)

	//text' , 'image' ,'video' ,'document
	if body.MessageType != "image" && body.MessageType != "text" && body.MessageType != "video" && body.MessageType != "document" {
		return utils.ErrorResponse{
			Code:    500,
			Message: " message type should be'text' , 'image' ,'video' ,'document",
		}
	}

	if body.ChatID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "chat id required",
		}
	}

	if !repository.CheckChatIdWithSendorIdExistRepository(body.ChatID, body.SenderID) {
		return utils.ErrorResponse{
			Code:    403,
			Message: "Sendor not belongs to chat",
		}
	}

	if body.SenderID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "SenderID required",
		}
	}

	response, err := repository.CreateMessagesRepository(body)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    201,
		Message: "message created successfully",
		Data:    response,
	}

}

// get messages list services
func GetMessagesListServices(pageStr string, limitStr string, chatId string, userId string) interface{} {

	if !repository.CheckChatIdWithSendorIdExistRepository(chatId, userId) {
		return utils.ErrorResponse{
			Code:    403,
			Message: "You are Not Authorized!!",
		}
	}

	response, err := repository.GetMessagesListRepository(pageStr, limitStr, chatId)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "success",
		Data:    response,
	}
}

// get list of chat of user
func GetChatListServices(userId string) interface{} {
	response, err := repository.GetUserChatListRepository(userId)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "success",
		Data:    response,
	}
}

// delete chat
func RemoveUserFromChatListServices(userid string, chatUserId string) interface{} {
	response, err := repository.RemoveUserFromChatRepository(userid, chatUserId)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "chat deleted successfully",
		Data:    response,
	}
}
