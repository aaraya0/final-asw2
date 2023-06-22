package services

import (
	"encoding/json"
	dto "final-asw2/services/items/dtos"
	"final-asw2/services/search/config"
	client "final-asw2/services/search/services/repositories"
	e "final-asw2/services/search/utils/errors"
	"fmt"
	"io"
	"net/http"
	"strings"

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
	queryParams := strings.Split(query, "_")
	field, query := queryParams[0], queryParams[1]
	itemsDto, err := s.solr.GetQuery(query, field)
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("Solr failed")
	}
	return itemsDto, nil
}

func (s *SolrService) GetQueryAllFields(query string) (dto.ItemsDto, e.ApiError) {
	var itemsDto dto.ItemsDto
	itemsDto, err := s.solr.GetQueryAllFields(query)
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("Solr failed")
	}
	return itemsDto, nil
}

func (s *SolrService) AddFromId(id string) e.ApiError {
	var itemDto dto.ItemDto
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/items/%s", config.ITEMSHOST, config.ITEMSPORT, id))
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

func (s *SolrService) Delete(id string) e.ApiError {
	err := s.solr.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
