package repositories

import (
	"fmt"

	"github.com/aaraya0/final-asw2/services/users/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserClient struct {
	Db *gorm.DB
}

func NewUserInterface(DBUser string, DBPass string, DBHost string, DBPort int, DBName string) *UserClient {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&parseTime=True", DBUser, DBPass, DBHost, DBPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Error initializing SQL: %v", err))
	}
	err = createDatabaseAndSchema(db, DBName)
	if err != nil {
		panic(fmt.Sprintf("Error creating db: %v", err))
	}

	return &UserClient{
		Db: db,
	}

}
func (s *UserClient) StartDbEngine() error {
	err := s.Db.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("error migrating database tables: %w", err)
	}

	log.Info("Finished migrating database tables")
	return nil
}

func createDatabaseAndSchema(db *gorm.DB, dbName string) error {
	// Create the database
	err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci", dbName)).Error
	if err != nil {
		return fmt.Errorf("error creating database: %w", err)
	}

	// Switch to the created database
	err = db.Exec(fmt.Sprintf("USE `%s`", dbName)).Error
	if err != nil {
		return fmt.Errorf("error switching to database: %w", err)
	}

	return nil
}

func (s *UserClient) GetUserById(id int) model.User {
	var user model.User
	s.Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func (s *UserClient) GetUserByUname(username string) (model.User, error) {
	var user model.User
	result := s.Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (s *UserClient) GetUsers() model.Users {
	var users model.Users
	s.Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}

func (s *UserClient) InsertUser(user model.User) model.User {
	result := s.Db.Create(&user)

	if result.Error != nil {
		log.Error(result.Error)
		user.ID = 0
		return user
	}
	log.Debug("User Created: ", user.ID)
	return user
}

func (s *UserClient) DeleteUser(user model.User) error {
	result := s.Db.Delete(&user)
	return result.Error
}

func (s *UserClient) UpdateUser(user model.User) (model.User, error) {
	result := s.Db.Save(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}
