package controllers

import (
	"final-asw2/services/images_worker/config"
	"final-asw2/services/images_worker/services"
	client "final-asw2/services/images_worker/services/repositories"
	con "final-asw2/services/images_worker/utils/connections"
)

var (
	Worker = services.NewWorker(
		(*client.QueueClient)(con.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)),
	)
)

func StartWorker() {

	Worker.TopicWorker("*.create")

}
