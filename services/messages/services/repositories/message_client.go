package repositories

import (
	"fmt"

	"github.com/aaraya0/final-asw2/services/messages/model"
	e "github.com/aaraya0/final-asw2/services/messages/utils/errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MessageClient struct {
	Db *gorm.DB
}

func NewMessageInterface(DBUser string, DBPass string, DBHost string, DBPort int, DBName string) *MessageClient {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&parseTime=True", DBUser, DBPass, DBHost, DBPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Error initializing SQL: %v", err))
	}
	err = createDatabaseAndSchema(db, DBName)
	if err != nil {
		panic(fmt.Sprintf("Error creating db: %v", err))
	}

	return &MessageClient{
		Db: db,
	}

}
func (s *MessageClient) StartDbEngine() error {
	err := s.Db.AutoMigrate(&model.Message{})
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

func (s *MessageClient) GetMessageById(id int) model.Message {
	var message model.Message
	s.Db.Where("id = ?", id).First(&message)
	log.Debug("Message: ", message)

	return message
}

func (s *MessageClient) DeleteMessageById(message model.Message) e.ApiError {
	result := s.Db.Delete(&message)
	return e.NewInternalServerApiError("Error deleting", result.Error)
}

func (s *MessageClient) GetMessagesByUserId(id int) (model.Messages, error) {
	var messages model.Messages
	result := s.Db.Find(&messages).Where("user_id = ?", id)
	if result.Error != nil {
		return messages, result.Error
	}

	return messages, nil
}

func (s *MessageClient) GetMessages() model.Messages {
	var messages model.Messages
	s.Db.Find(&messages)

	log.Debug("Messages: ", messages)

	return messages
}

func (s *MessageClient) InsertMessage(message model.Message) model.Message {
	result := s.Db.Create(&message)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Message Created: ", message.ID)
	return message
}

func (s *MessageClient) GetMessagesByItemId(id string) (model.Messages, error) {
	var messages model.Messages
	result := s.Db.Find(&messages).Where("item_id = ?", id)
	if result.Error != nil {
		return messages, result.Error
	}

	return messages, nil
}
