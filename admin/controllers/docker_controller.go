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
	log.Debug(image)
	id, err := Docker.CreateContainer(image, "test")
	if err != nil {
		c.JSON(http.StatusBadRequest, id)
		return
	}

	c.JSON(http.StatusOK, id)

}
