package services

import (
	"testing"

	dto "github.com/aaraya0/final-asw2/services/users/dtos"
	"github.com/aaraya0/final-asw2/services/users/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserClientMock struct {
	mock.Mock
}

func (m *UserClientMock) GetUserById(id int) model.User {
	ret := m.Called(id)
	return ret.Get(0).(model.User)
}

func (m *UserClientMock) GetUsers() model.Users {
	ret := m.Called()
	return ret.Get(0).(model.Users)
}

func (m *UserClientMock) GetUserByUsername(username string) (model.User, error) {
	ret := m.Called(username)
	return ret.Get(0).(model.User), ret.Error(1)
}

func (m *UserClientMock) InsertUser(user model.User) model.User {
	ret := m.Called(user)
	return ret.Get(0).(model.User)
}

func TestGetUserById(t *testing.T) {
	mockUserClient := new(UserClientMock)

	var user model.User
	user.ID = 1
	user.Username = "test_username"
	user.Password = "test_password"
	user.FirstName = "test_firstname"
	user.LastName = "test_lastname"
	user.Email = "email@email"

	var emptyUser model.User
	emptyUser.ID = 0

	var userDto dto.UserDto
	userDto.UserId = 1
	userDto.Username = "test_username"
	userDto.FirstName = "test_firstname"
	userDto.LastName = "test_lastname"
	userDto.Email = "email@email"

	var emptyDto dto.UserDto

	mockUserClient.On("GetUserById", 1).Return(user)
	mockUserClient.On("GetUserById", 0).Return(emptyUser)
	service := NewUserService(mockUserClient, nil) // Replace 'nil' with a mock for QueueClient if you're using it.

	res, err := service.GetUserById(1)
	res2, err2 := service.GetUserById(0)

	assert.Nil(t, err, "Error should be Nil")
	assert.NotNil(t, err2, "Error should NOT be Nil")

	assert.Equal(t, res, userDto)   // Shouldn't return pass
	assert.Equal(t, res2, emptyDto) // Should be empty
}

func TestGetUsers(t *testing.T) {
	mockUserClient := new(UserClientMock)

	var user model.User
	user.ID = 1
	user.Username = "test_username"
	user.Password = "test_password"
	user.FirstName = "test_firstname"
	user.LastName = "test_lastname"
	user.Email = "email@email"

	var users model.Users
	users = append(users, user)

	mockUserClient.On("GetUsers").Return(users)
	service := NewUserService(mockUserClient, nil) // Replace 'nil' with a mock for QueueClient if you're using it.

	res, err := service.GetUsers()

	assert.Nil(t, err, "Error should be Nil")
	assert.NotEqual(t, 0, len(res)) // Should be empty
}
