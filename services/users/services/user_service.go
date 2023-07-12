package services

import (
	dto "github.com/aaraya0/final-asw2/services/users/dtos"
	e "github.com/aaraya0/final-asw2/services/users/utils/errors"
)

type UserServices interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	DeleteUser(id int) e.ApiError
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
	Login(loginDto dto.LoginDto) (dto.LoginResponseDto, e.ApiError)
	UpdateUser(id int, userDto dto.UserDto) (dto.UserDto, e.ApiError)
}
