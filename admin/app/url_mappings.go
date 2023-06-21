package app

import (
	dockerController "admin/controllers"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Mapeo productos
	router.POST("/container/:image", dockerController.CreateContainer)

	log.Info("Finishing mapping configurations")
}
