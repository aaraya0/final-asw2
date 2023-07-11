package app

import (
	userController "github.com/aaraya0/final-asw2/services/users/controllers"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Users Mapping
	router.GET("/users/:id", userController.GetUserById)
	router.GET("/users", userController.GetUsers)
	router.POST("/user", userController.InsertUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	// Login Mapping
	router.POST("/login", userController.Login)

	log.Info("Finishing mappings configurations")
}
