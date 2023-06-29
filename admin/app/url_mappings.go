package app

import (
	dockerController "github.com/aaraya0/final-asw2/admin/controllers"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.POST("/container/:image/:name", dockerController.CreateContainer)
	router.DELETE("/container/:containerID", dockerController.RemoveContainer)
	router.GET("/containers", dockerController.ListContainers)

	log.Info("Finishing mapping configurations")
}
