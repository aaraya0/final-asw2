package messageController

import (
	"net/http"
	"strconv"

	"github.com/aaraya0/final-asw2/services/messages/config"
	dto "github.com/aaraya0/final-asw2/services/messages/dtos"
	service "github.com/aaraya0/final-asw2/services/messages/services"
	client "github.com/aaraya0/final-asw2/services/messages/services/repositories"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	messageService = service.NewMessageServiceImpl(
		client.NewMessageInterface(config.SQLUSER, config.SQLPASS, config.SQLHOST, config.SQLPORT, config.SQLDB),
		client.NewQueueClient("guest", "guest", "localhost", 5672),
	)
)

func GetMessageById(c *gin.Context) {
	log.Debug("Message id: " + c.Param("id"))

	// Get Back Message

	var messageDto dto.MessageDto
	id, _ := strconv.Atoi(c.Param("id"))
	messageDto, err := messageService.GetMessageById(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, messageDto)
}

func DeleteMessageById(c *gin.Context) {
	log.Debug("Message id: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	err := messageService.DeleteMessageById(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func GetMessagesByUserId(c *gin.Context) {
	log.Debug("User id: " + c.Param("id"))

	// Get Back Messages

	var messagesDto dto.MessagesDto
	id, _ := strconv.Atoi(c.Param("id"))
	messagesDto, err := messageService.GetMessagesByUserId(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, messagesDto)
}

func GetMessages(c *gin.Context) {

	var messagesDto dto.MessagesDto
	messagesDto, err := messageService.GetMessages()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, messagesDto)
}

func MessageInsert(c *gin.Context) {
	var messageDto dto.MessageDto
	err := c.BindJSON(&messageDto)

	log.Debug(messageDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	messageDto, er := messageService.InsertMessage(messageDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, messageDto)
}

func GetMessagesByItemId(c *gin.Context) {
	log.Debug("Item id: " + c.Param("id"))

	// Get Back Messages

	var messagesDto dto.MessagesDto
	id := c.Param("id")
	messagesDto, err := messageService.GetMessagesByItemId(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, messagesDto)
}

func DeleteUserMessages(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := messageService.DeleteUserMessages(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
