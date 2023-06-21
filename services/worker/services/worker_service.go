package services

import (
	"fmt"
	"net/http"

	"github.com/aaraya0/arq-software/final-asw2/worker/config"
	client "github.com/aaraya0/arq-software/final-asw2/worker/services/repositories"

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

func (s *WorkerService) QueueWorker(qname string) {
	err := s.queue.ProcessMessages(qname, func(id string) {
		resp, err := http.Get(fmt.Sprintf("http://%s:%d/item/%s", config.LBHOST, config.LBPORT, id))
		log.Debug("Item sent " + id)
		if err != nil {
			log.Debug("error in get request")
			log.Debug(resp)
		}
	})
	if err != nil {
		log.Error("Error starting worker processing", err)
	}
}
