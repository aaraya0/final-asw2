package controllers

import (
	"fmt"
	"net/http"

	"github.com/aaraya0/final-asw2/services/items/config"
	dtos "github.com/aaraya0/final-asw2/services/items/dtos"
	service "github.com/aaraya0/final-asw2/services/items/services"
	client "github.com/aaraya0/final-asw2/services/items/services/repositories"

	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	itemService = service.NewItemService(
		client.NewItemInterface(config.MONGOHOST, config.MONGOPORT, config.MONGOCOLLECTION),
		client.NewMemcachedInterface(config.MEMCACHEDHOST, config.MEMCACHEDPORT),
		client.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT),
	)
)

func GetItem(c *gin.Context) {
	var itemDto dtos.ItemDto
	id := c.Param("item_id")
	itemDto, err := itemService.GetItem(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, itemDto)
}

func InsertItem(c *gin.Context) {
	var itemDto dtos.ItemDto
	err := c.BindJSON(&itemDto)

	// Error Parsing json param
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemDto, er := itemService.InsertItem(itemDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, itemDto)
}

func QueueItems(c *gin.Context) {
	var itemsDto dtos.ItemsDto
	err := c.BindJSON(&itemsDto)
	// Error Parsing json param
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := itemService.QueueItems(itemsDto)

	// Error Queueing
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, itemsDto)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("item_id")
	err := itemService.DeleteItem(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func GetItemsByUId(c *gin.Context) {
	var itemsDto dtos.ItemsDto
	id, _ := strconv.Atoi(c.Param("id"))
	itemsDto, err := itemService.GetItemsByUId(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, itemsDto)

}

func DeleteUserItems(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := itemService.DeleteUserItems(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func DownloadImage(c *gin.Context) {
	id := c.Param("item_id")
	err := itemService.DownloadImage(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
