package controllers

import (
	"net/http"

	"github.com/aaraya0/final-asw2/admin/services"
	client "github.com/aaraya0/final-asw2/admin/services/repositories"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	Docker = services.NewDockerServiceImpl(
		client.NewDockerClient(),
	)
)

func CreateContainer(c *gin.Context) {

	image := c.Param("image")
	name := c.Param("name")
	log.Debug(image)
	id, err := Docker.CreateContainer(image, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, id)
		return
	}

	c.JSON(http.StatusOK, id)

}

func RemoveContainer(c *gin.Context) {
	containerID := c.Param("containerID")
	err := Docker.RemoveContainer(containerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Container removed successfully",
	})
}

func ListContainers(c *gin.Context) {
	containers, err := Docker.ListContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list containers"})
		return
	}

	c.JSON(http.StatusOK, containers)
}
