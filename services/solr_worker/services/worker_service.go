package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aaraya0/final-asw2/services/solr_worker/config"
	client "github.com/aaraya0/final-asw2/services/solr_worker/services/repositories"

	log "github.com/sirupsen/logrus"
)

type WorkerService struct {
	queue *client.QueueClient
}

func NewWorker(
	queue *client.QueueClient,
) *WorkerService {
	return &WorkerService{
		queue: queue,
	}
}

func (s *WorkerService) TopicWorker(topic string) {
	err := s.queue.ProcessMessages(config.EXCHANGE, topic, func(id string) {
		var resp *http.Response
		var resp2 *http.Response
		var err error
		cli := &http.Client{}
		strs := strings.Split(id, ".")
		if len(strs) < 2 {
			resp, err = http.Get(fmt.Sprintf("http://%s:%d/items/%s", config.SOLRHOST, config.SOLRPORT, id))
		} else {
			if strs[1] == "delete" {
				req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%d/items/%s", config.ITEMSHOST, config.ITEMSPORT, strs[0]), nil)
				if err != nil {
					log.Error(err)
				}
				resp, err = cli.Do(req)
				if err != nil {
					log.Error(err)
					log.Debug(resp)
				}
				req2, err2 := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%d/items/%s", config.SOLRHOST, config.SOLRPORT, strs[0]), nil)
				if err2 != nil {
					log.Error(err2)
				}
				resp2, err2 = cli.Do(req2)
				if err2 != nil {
					log.Error(err2)
					log.Debug(resp2)
				}

			}
		}
		log.Debug("Item sent " + id)
		if err != nil {
			log.Debug("error in get request")
			log.Debug(resp, resp2)
		}
	})
	if err != nil {
		log.Error("Error starting worker processing", err)
	}
}
