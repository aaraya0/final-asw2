package controllers

import (
	"final_asw2/services/solr_worker/config"
	"final_asw2/services/solr_worker/services"
	client "final_asw2/services/solr_worker/services/repositories"
	con "final_asw2/services/solr_worker/utils/connections"
)

var (
	Worker = services.NewWorker(
		(*client.QueueClient)(con.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)),
	)
)

func StartWorker() {

	Worker.TopicWorker("*.*")

}
