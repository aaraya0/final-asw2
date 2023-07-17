package controllers

import (
	"github.com/aaraya0/final-asw2/services/solr_worker/config"
	"github.com/aaraya0/final-asw2/services/solr_worker/services"
	client "github.com/aaraya0/final-asw2/services/solr_worker/services/repositories"
	con "github.com/aaraya0/final-asw2/services/solr_worker/utils/connections"
)

var (
	Worker = services.NewWorker(
		(*client.QueueClient)(con.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)),
	)
)

func StartWorker() {

	Worker.TopicWorker("*.*")
}
