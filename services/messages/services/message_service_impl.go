package services

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	dto "github.com/aaraya0/final-asw2/services/messages/dtos"
	"github.com/aaraya0/final-asw2/services/messages/model"
	client "github.com/aaraya0/final-asw2/services/messages/services/repositories"
	e "github.com/aaraya0/final-asw2/services/messages/utils/errors"
)

type MessageServiceImpl struct {
	messageDB *client.MessageClient
	queue     *client.QueueClient
}

func NewMessageServiceImpl(
	messageDB *client.MessageClient,
	queue *client.QueueClient,
) *MessageServiceImpl {
	messageDB.StartDbEngine()
	return &MessageServiceImpl{
		messageDB: messageDB,
		queue:     queue,
	}
}

func (s *MessageServiceImpl) GetMessageById(id int) (dto.MessageDto, e.ApiError) {

	var message = s.messageDB.GetMessageById(id)
	var messageDto dto.MessageDto

	if message.ID == 0 {
		return messageDto, e.NewBadRequestApiError("message not found")
	}
	messageDto.MessageId = message.ID
	messageDto.UserId = message.UserId
	messageDto.ItemId = message.ItemId
	messageDto.Body = message.Body
	messageDto.CreatedAt = message.CreatedAt
	messageDto.System = message.System
	return messageDto, nil
}

func (s *MessageServiceImpl) DeleteMessageById(id int) e.ApiError {

	var message model.Message
	message.ID = id
	err := s.messageDB.DeleteMessageById(message)
	if err != nil {
		return e.NewInternalServerApiError("Error deleting message", err)
	}
	return nil
}

func (s *MessageServiceImpl) GetMessagesByUserId(id int) (dto.MessagesDto, e.ApiError) {

	var messagesDto dto.MessagesDto
	var messages, err = s.messageDB.GetMessagesByUserId(id)

	if err != nil {
		return messagesDto, e.NewBadRequestApiError(err.Error())
	}

	for _, message := range messages {
		var messageDto dto.MessageDto
		messageDto.CreatedAt = message.CreatedAt
		messageDto.UserId = message.UserId
		messageDto.ItemId = message.ItemId
		messageDto.Body = message.Body
		messageDto.MessageId = message.ID
		messageDto.System = message.System

		messagesDto = append(messagesDto, messageDto)
	}
	return messagesDto, nil
}

func (s *MessageServiceImpl) GetMessages() (dto.MessagesDto, e.ApiError) {

	var messages = s.messageDB.GetMessages()
	var messagesDto dto.MessagesDto

	for _, message := range messages {
		var messageDto dto.MessageDto
		messageDto.CreatedAt = message.CreatedAt
		messageDto.UserId = message.UserId
		messageDto.ItemId = message.ItemId
		messageDto.Body = message.Body
		messageDto.MessageId = message.ID
		messageDto.System = message.System

		messagesDto = append(messagesDto, messageDto)
	}

	return messagesDto, nil
}

func (s *MessageServiceImpl) InsertMessage(messageDto dto.MessageDto) (dto.MessageDto, e.ApiError) {

	var message model.Message

	message.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	message.UserId = messageDto.UserId
	message.ItemId = messageDto.ItemId
	message.Body = messageDto.Body
	message.ID = messageDto.MessageId
	message.System = messageDto.System

	message = s.messageDB.InsertMessage(message)

	messageDto.MessageId = message.ID
	messageDto.CreatedAt = message.CreatedAt
	s.queue.SendMessage(strconv.Itoa(message.UserId), message.ItemId, fmt.Sprintf("%d", message.ID))

	return messageDto, nil
}

func (s *MessageServiceImpl) GetMessagesByItemId(id string) (dto.MessagesDto, e.ApiError) {

	var messagesDto dto.MessagesDto
	var messages, err = s.messageDB.GetMessagesByItemId(id)

	if err != nil {
		return messagesDto, e.NewBadRequestApiError(err.Error())
	}

	for _, message := range messages {
		var messageDto dto.MessageDto
		messageDto.CreatedAt = message.CreatedAt
		messageDto.UserId = message.UserId
		messageDto.ItemId = message.ItemId
		messageDto.Body = message.Body
		messageDto.MessageId = message.ID
		messageDto.System = message.System

		messagesDto = append(messagesDto, messageDto)
	}
	return messagesDto, nil
}

func (s *MessageServiceImpl) DeleteUserMessages(id int) e.ApiError {
	messages, err := s.GetMessagesByUserId(id)
	if err != nil {
		log.Error(err)
		return err
	}
	for i := range messages {
		var message dto.MessageDto
		message = messages[i]
		go func() {
			err := s.DeleteMessageById(message.MessageId)
			if err != nil {
				log.Error(err)
			}
			err = s.queue.SendMessage(strconv.Itoa(message.MessageId), "delete", fmt.Sprintf("%s.delete", strconv.Itoa(message.MessageId)))
			log.Error(err)
		}()
	}
	return nil
}
