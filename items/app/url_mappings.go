package app

import (
	itemController "github.com/aaraya0/arq-software/final-asw2/items/controllers"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Items Mapping
	router.GET("/items/:item_id", itemController.GetItem)
	router.POST("/item", itemController.InsertItem)
	router.POST("/items", itemController.QueueItems)

	log.Info("Finishing mappings configurations")
}
