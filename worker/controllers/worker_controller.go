package worker

import (
	"github.com/aaraya0/arq-software/final-asw2/worker/config"
	"github.com/aaraya0/arq-software/final-asw2/worker/services"
	client "github.com/aaraya0/arq-software/final-asw2/worker/services/repositories"
	con "github.com/aaraya0/arq-software/final-asw2/worker/utils/connections"
)

var (
	Worker = services.NewWorker(
		(*client.QueueClient)(con.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)),
	)
)

func StartWorker() {

	Worker.QueueWorker("solr")

}
