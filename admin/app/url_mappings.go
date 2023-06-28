package app

import (
	dockerController "github.com/aaraya0/final-asw2/admin/controllers"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Mapeo productos
	router.POST("/container/:image", dockerController.CreateContainer)

	log.Info("Finishing mapping configurations")
}
