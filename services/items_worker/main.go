package main

import (
	worker "github.com/aaraya0/final-asw2/services/items_worker/controllers"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Starting worker")
	worker.StartWorker()
}
