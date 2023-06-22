package services

import (
	"fmt"
	"net/http"

	"github.com/aaraya0/final-asw2/services/images_worker/config"
	client "github.com/aaraya0/final-asw2/services/images_worker/services/repositories"
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
		client := &http.Client{}
		req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%d/users/%s/items", config.ITEMSHOST, config.ITEMSPORT, id), nil)
		log.Debug("Item delete sent " + id)

		if err != nil {
			log.Debug("error in delete request")
		}

		_, err = client.Do(req)
		if err != nil {
			log.Error(err)
		}

	})
	if err != nil {
		log.Error("Error starting worker processing", err)
	}
}