package solrController

import (
	"net/http"

	"github.com/aaraya0/arq-software/final-asw2/search/config"
	dto "github.com/aaraya0/arq-software/final-asw2/search/dtos"
	"github.com/aaraya0/arq-software/final-asw2/search/services"
	client "github.com/aaraya0/arq-software/final-asw2/search/services/repositories"
	con "github.com/aaraya0/arq-software/final-asw2/search/utils/connections"
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

func Add(c *gin.Context) {
	id := c.Param("id")
	err := Solr.Add(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, err)
}
