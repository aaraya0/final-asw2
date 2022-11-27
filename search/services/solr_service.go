package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aaraya0/arq-software/final-asw2/search/config"
	dto "github.com/aaraya0/arq-software/final-asw2/search/dtos"
	client "github.com/aaraya0/arq-software/final-asw2/search/services/repositories"
	e "github.com/aaraya0/arq-software/final-asw2/search/utils/errors"
	log "github.com/sirupsen/logrus"
)

type SolrService struct {
	solr *client.SolrClient
}

func NewSolrServiceImpl(
	solr *client.SolrClient,
) *SolrService {
	return &SolrService{
		solr: solr,
	}
}

func (s *SolrService) GetQuery(query string) (dto.ItemsDto, e.ApiError) {
	var itemsDto dto.ItemsDto
	itemsDto, err := s.solr.GetQuery(query)
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("Solr failed")
	}
	return itemsDto, nil
}

func (s *SolrService) Add(id string) e.ApiError {
	var itemDto dto.ItemDto
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/item/%s", config.ITEMSHOST, config.ITEMSPORT, id))
	if err != nil {
		log.Debugf("error getting item %s", id)
		return e.NewBadRequestApiError("error getting item " + id)
	}
	var body []byte
	body, _ = io.ReadAll(resp.Body)
	log.Debugf("%s", body)
	err = json.Unmarshal(body, &itemDto)
	if err != nil {
		log.Debugf("error in unmarshal of item %s", id)
		return e.NewBadRequestApiError("error in unmarshal of item")
	}
	er := s.solr.Add(itemDto)
	if er != nil {
		log.Debugf("error adding to solr")
		return e.NewInternalServerApiError("Adding to Solr error", err)
	}
	return nil
}
