package controllers

import (
	"net/http"

	"github.com/aaraya0/final-asw2/services/search/config"
	"github.com/aaraya0/final-asw2/services/search/dto"
	"github.com/aaraya0/final-asw2/services/search/services"
	client "github.com/aaraya0/final-asw2/services/search/services/repositories"
	con "github.com/aaraya0/final-asw2/services/search/utils/connections"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	Solr = services.NewSolrServiceImpl(
		(*client.SolrClient)(con.NewSolrClient(config.SOLRHOST, config.SOLRPORT, config.SOLRCOLLECTION)),
	)
)

func GetQuery(c *gin.Context) {
	var itemsDto dto.ItemsDto
	query := c.Param("searchQuery")

	itemsDto, err := Solr.GetQuery(query)
	if err != nil {
		log.Debug(itemsDto)
		c.JSON(http.StatusBadRequest, itemsDto)
		return
	}

	c.JSON(http.StatusOK, itemsDto)

}

func GetQueryAllFields(c *gin.Context) {
	var itemsDto dto.ItemsDto
	query := c.Param("searchQuery")

	itemsDto, err := Solr.GetQueryAllFields(query)
	if err != nil {
		log.Debug(itemsDto)
		c.JSON(http.StatusBadRequest, itemsDto)
		return
	}

	c.JSON(http.StatusOK, itemsDto)

}

func AddFromId(c *gin.Context) {
	id := c.Param("id")
	err := Solr.AddFromId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, err)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	err := Solr.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, err)
}
