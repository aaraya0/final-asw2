package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aaraya0/final-asw2/services/search/config"
	"github.com/aaraya0/final-asw2/services/search/dto"
	e "github.com/aaraya0/final-asw2/services/search/utils/errors"
	logger "github.com/sirupsen/logrus"
	solr "github.com/stevenferrer/solr-go"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func (sc *SolrClient) GetQuery(query string, field string) (dto.ItemsDto, e.ApiError) {
	var response dto.SolrResponseDto
	var itemsDto dto.ItemsDto
	q, err := http.Get(fmt.Sprintf("http://%s:%d/solr/items/select?q=%s%s%s", config.SOLRHOST, config.SOLRPORT, field, "%3A", query))

	if err != nil {
		return itemsDto, e.NewBadRequestApiError("error getting from solr")
	}

	defer q.Body.Close()
	err = json.NewDecoder(q.Body).Decode(&response)
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("error in unmarshal")
	}
	itemsDto = response.Response.Docs
	return itemsDto, nil
}

func (sc *SolrClient) GetQueryAllFields(query string) (dto.ItemsDto, e.ApiError) {
	var response dto.SolrResponseDto
	var itemsDto dto.ItemsDto

	q, err := http.Get(
		fmt.Sprintf("http://%s:%d/solr/items/query?q=*:*&q.op=OR&indent=true&title:%20%22"+query+"%22=&seller:%20%22"+query+"%22%20=&location:%20%22"+query+"%22%20=&description:%20%22"+query+"%22%20=&useParams=&qt=%2Fselect",
			config.SOLRHOST, config.SOLRPORT))
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

func (sc *SolrClient) Delete(id string) e.ApiError {
	var deleteDto dto.DeleteDto
	deleteDto.Delete = dto.DeleteDoc{Query: fmt.Sprintf("id:%s", id)}
	data, err := json.Marshal(deleteDto)
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
