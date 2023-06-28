package repositories

import (
	dto "items/dtos"
	errors "items/utils/errors"
)

type Client interface {
	GetItem(id string) (dto.ItemDto, errors.ApiError)
	InsertItem(dto.ItemDto) (dto.ItemDto, errors.ApiError)
	//Update(book dto.ItemDto) (dto.ItemDto, errors.ApiError)
	//Delete(id string) errors.ApiError
}
