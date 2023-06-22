package services

import (
	dtos "final-asw2/services/items/dtos"
	e "final-asw2/services/items/utils/errors"
)

type ItemService interface {
	GetItem(id string) (dtos.ItemDto, e.ApiError)
	InsertItem(item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
	QueueItems(items dtos.ItemsDto) e.ApiError
}
