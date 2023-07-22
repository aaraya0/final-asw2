package app

import (
	itemController "github.com/aaraya0/final-asw2/services/items/controllers"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Items Mapping
	router.GET("/items/:item_id", itemController.GetItem)
	router.POST("/item", itemController.InsertItem)
	router.POST("/items", itemController.QueueItems)
	router.DELETE("/item/:item_id", itemController.DeleteItem)

	router.DELETE("/users/:id/items", itemController.DeleteUserItems)
	router.GET("/users/:id/items", itemController.GetItemsByUId)

	router.GET("/items/image/:item_id", itemController.DownloadImage)

	log.Info("Finishing mappings configurations")
}
