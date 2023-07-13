package services

import (
	"fmt"

	dto "github.com/aaraya0/final-asw2/services/users/dtos"
	"github.com/aaraya0/final-asw2/services/users/model"
	client "github.com/aaraya0/final-asw2/services/users/services/repositories"
	e "github.com/aaraya0/final-asw2/services/users/utils/errors"

	"github.com/golang-jwt/jwt"

	log "github.com/sirupsen/logrus"
)

type UserService struct {
	userDB *client.UserClient
	queue  *client.QueueClient
}

func NewUserService(
	userDB *client.UserClient,
	queue *client.QueueClient,

) *UserService {
	userDB.StartDbEngine()
	return &UserService{
		userDB: userDB,
		queue:  queue,
	}
}

func (s *UserService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user = s.userDB.GetUserById(id)
	var userDto dto.UserDto

	if user.ID == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.FirstName = user.FirstName
	userDto.LastName = user.LastName
	userDto.Username = user.Username
	userDto.UserId = user.ID
	userDto.Email = user.Email
	return userDto, nil
}

func (s *UserService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users = s.userDB.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.FirstName = user.FirstName
		userDto.LastName = user.LastName
		userDto.Username = user.Username
		userDto.Email = user.Email
		userDto.UserId = user.ID

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *UserService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.Username = userDto.Username
	user.Password = userDto.Password
	user.Email = userDto.Email

	user = s.userDB.InsertUser(user)

	userDto.UserId = user.ID

	err := s.queue.SendMessage(userDto.UserId, "create", fmt.Sprintf("New user, user id: %d", userDto.UserId))
	if err != nil {
		return userDto, e.NewInternalServerApiError("error sending created message on user creation", err)
	}
	return userDto, nil
}

func (s *UserService) Login(loginDto dto.LoginDto) (dto.LoginResponseDto, e.ApiError) {

	var user model.User
	user, err := s.userDB.GetUserByUname(loginDto.Username)
	var loginResponseDto dto.LoginResponseDto
	loginResponseDto.UserId = -1
	if err != nil {
		return loginResponseDto, e.NewBadRequestApiError("Usuario no encontrado")
	}
	if user.Password != loginDto.Password && loginDto.Username != "encrypted" {
		return loginResponseDto, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginDto.Username,
		"pass":     loginDto.Password,
	})
	var jwtKey = []byte("secret_key")
	tokenString, _ := token.SignedString(jwtKey)
	if user.Password != tokenString && loginDto.Username == "encrypted" {
		return loginResponseDto, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	loginResponseDto.UserId = user.ID
	loginResponseDto.Token = tokenString
	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}

func (s *UserService) DeleteUser(id int) e.ApiError {

	err := s.queue.SendMessage(id, "delete", fmt.Sprintf("%d", id))
	if err != nil {
		return e.NewInternalServerApiError("Error notifying other systems of user deletion. Canceling delete", err)
	}

	var user model.User
	user.ID = id
	er := s.userDB.DeleteUser(user)
	if er != nil {
		return e.NewInternalServerApiError("Error deleting user", er)
	}

	return nil
}

func (s *UserService) UpdateUser(id int, userDto dto.UserDto) (dto.UserDto, e.ApiError) {
	// Obtener el usuario existente por su ID
	user := s.userDB.GetUserById(id)

	// Actualizar los campos del usuario con los valores proporcionados en userDto
	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.Username = userDto.Username
	user.Email = userDto.Email

	// Guardar los cambios en la base de datos
	updatedUser, err := s.userDB.UpdateUser(user)
	if err != nil {
		return dto.UserDto{}, e.NewInternalServerApiError("Error updating user", err)
	}

	// Crear y devolver el UserDto actualizado
	updatedUserDto := dto.UserDto{
		UserId:    updatedUser.ID,
		FirstName: updatedUser.FirstName,
		LastName:  updatedUser.LastName,
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
	}
	er := s.queue.SendMessage(userDto.UserId, "update", fmt.Sprintf("User updated, user id: %d", userDto.UserId))
	if er != nil {
		return userDto, e.NewInternalServerApiError("error sending created message on user update", er)
	}
	return updatedUserDto, nil
}
