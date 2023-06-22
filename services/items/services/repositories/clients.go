package repositories

import (
	dto "final-asw2/items/dtos"
	"final-asw2/items/utils/errors"
)

type Client interface {
	GetItem(id string) (dto.ItemDto, errors.ApiError)
	InsertItem(dto.ItemDto) (dto.ItemDto, errors.ApiError)
	//Update(book dto.ItemDto) (dto.ItemDto, errors.ApiError)
	//Delete(id string) errors.ApiError
}
