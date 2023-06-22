package repositories

import (
	"final-asw2/services/messages/model"
)

type Client interface {
	GetMessageById(id int) model.Message
	GetMessagesByUserId(id int) (model.Messages, error)
	GetMessages() model.Messages
	InsertMessage(message model.Message) model.Message
}
