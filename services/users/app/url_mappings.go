package app

import (
	userController "final-asw2/services/users/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Users Mapping
	router.GET("/users/:id", userController.GetUserById)
	router.GET("/users", userController.GetUsers)
	router.POST("/user", userController.UserInsert)
	router.DELETE("/user/:id", userController.DeleteUser)

	// Login Mapping
	router.POST("/login", userController.Login)

	log.Info("Finishing mappings configurations")
}
