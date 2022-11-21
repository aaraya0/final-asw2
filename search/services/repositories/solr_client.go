package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aaraya0/arq-software/final-asw2/search/config"
	dto "github.com/aaraya0/arq-software/final-asw2/search/dtos"
	e "github.com/aaraya0/arq-software/final-asw2/search/utils/errors"
	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func (sc *SolrClient) GetQuery(query string) (dto.ItemsDto, e.ApiError) {
	var response dto.SolrResponseDto
	var itemsDto dto.ItemsDto

	q, err := http.Get(
		fmt.Sprintf("http://%s:%d/solr/items/select?q=title%s%s%sclass%s%s%slocation%s%s%sseller%s%s%sdescription%s%s%s",
			config.SOLRHOST, config.SOLRPORT,
			"%3A", query, "%0A",
			"%3A", query, "%0A",
			"%3A", query, "%0A",
			"%3A", query, "%0A",
			"%3A", query, "%0A"))
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("error getting from solr")
	}

	var body []byte
	body, err = io.ReadAll(q.Body)
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("error reading body")
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("error in unmarshal")
	}

	itemsDto = response.Response.Docs
	return itemsDto, nil
}

func (sc *SolrClient) Add(itemDto dto.ItemDto) e.ApiError {
	var addItemDto dto.AddDto
	addItemDto.Add = dto.DocDto{Doc: itemDto}
	data, err := json.Marshal(addItemDto)

	reader := bytes.NewReader(data)
	if err != nil {
		return e.NewBadRequestApiError("Error getting json")
	}
	resp, err := sc.Client.Update(context.TODO(), sc.Collection, solr.JSON, reader)
	logger.Debug(resp)
	if err != nil {
		return e.NewBadRequestApiError("Error in solr")
	}

	er := sc.Client.Commit(context.TODO(), sc.Collection)
	if er != nil {
		logger.Debug("Error committing load")
		return e.NewInternalServerApiError("Error committing to solr", er)
	}
	return nil
}
