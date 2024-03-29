package services

import (
	dto "github.com/aaraya0/final-asw2/services/messages/dtos"
	e "github.com/aaraya0/final-asw2/services/messages/utils/errors"
)

type MessageService interface {
	GetMessageById(id int) (dto.MessageDto, e.ApiError)
	GetMessagesByUserId(id int) (dto.MessagesDto, e.ApiError)
	GetMessagesByItemId(id string) (dto.MessagesDto, e.ApiError)
	GetMessages() (dto.MessagesDto, e.ApiError)
	InsertMessage(messageDto dto.MessageDto) (dto.MessageDto, e.ApiError)
	DeleteUserMessages(id int) e.ApiError
}
