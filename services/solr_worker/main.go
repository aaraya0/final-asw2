package main

import (
	worker "final-asw2/services/solr_worker/controllers"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Starting worker")
	worker.StartWorker()
}
