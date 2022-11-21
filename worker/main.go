package main

import (
	worker "github.com/aaraya0/arq-software/final-asw2/worker/controllers"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Starting worker")
	worker.StartWorker()
}
