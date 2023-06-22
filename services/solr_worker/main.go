package main

import (
	"worker/controllers/worker"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Starting worker")
	worker.StartWorker()
}
