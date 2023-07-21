package repositories

import (
	"github.com/aaraya0/final-asw2/services/messages/model"
	errors "github.com/aaraya0/final-asw2/services/messages/utils/errors"
)

type Client interface {
	GetMessageById(id int) model.Message
	GetMessagesByUserId(id int) (model.Messages, error)
	GetMessagesByItemId(id string) (model.Messages, error)
	GetMessages() model.Messages
	InsertMessage(message model.Message) model.Message
	DeleteUserMessages(id int) errors.ApiError
}
