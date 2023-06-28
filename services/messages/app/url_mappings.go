package app

import (
	messageController "github.com/aaraya0/final-asw2/services/messages/controllers"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Messages Mapping
	router.GET("/messages/:id", messageController.GetMessageById)
	router.GET("/users/:id/messages", messageController.GetMessagesByUserId)
	router.GET("/messages", messageController.GetMessages)

	router.DELETE("/messages/:id", messageController.DeleteMessageById)

	router.POST("/message", messageController.MessageInsert)

	log.Info("Finishing mappings configurations")
}
