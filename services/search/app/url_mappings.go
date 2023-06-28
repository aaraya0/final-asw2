package app

import (
	solrController "github.com/aaraya0/final-asw2/services/search/controllers"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/search=:searchQuery", solrController.GetQuery)
	router.GET("/searchAll=:searchQuery", solrController.GetQueryAllFields)
	router.GET("/items/:id", solrController.AddFromId)

	router.DELETE("/items/:id", solrController.Delete)

	log.Info("Finishing mappings configurations")
}
