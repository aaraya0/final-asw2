package services

import (
	dtos "github.com/aaraya0/arq-software/final-asw2/items/dtos"
	e "github.com/aaraya0/arq-software/final-asw2/items/utils/errors"
)

type ItemService interface {
	GetItem(id string) (dtos.ItemDto, e.ApiError)
	InsertItem(item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
	QueueItems(items dtos.ItemsDto) e.ApiError
}
