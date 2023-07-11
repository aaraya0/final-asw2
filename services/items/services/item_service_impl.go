package services

import (
	dtos "github.com/aaraya0/final-asw2/services/items/dtos"
	e "github.com/aaraya0/final-asw2/services/items/utils/errors"
)

type ItemServices interface {
	GetItem(id string) (dtos.ItemDto, e.ApiError)
	InsertItem(item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
	QueueItems(items dtos.ItemsDto) e.ApiError
}
