package app

import (
	solrController "github.com/aaraya0/arq-software/final-asw2/search/controllers"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping

	router.GET("/search=:searchQuery", solrController.GetQuery)
	router.GET("/items/:id", solrController.Add)

	log.Info("Finishing mappings configurations")
}
