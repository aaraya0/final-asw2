package services

import (
	dto "github.com/aaraya0/final-asw2/services/items/dtos"
	client "github.com/aaraya0/final-asw2/services/items/services/repositories"
	e "github.com/aaraya0/final-asw2/services/items/utils/errors"

	log "github.com/sirupsen/logrus"
)

type ItemService struct {
	item      *client.ItemClient
	memcached *client.MemcachedClient
	queue     *client.QueueClient
}

func NewItemService(
	item *client.ItemClient,
	memcached *client.MemcachedClient,
	queue *client.QueueClient,
) *ItemService {
	return &ItemService{
		item:      item,
		memcached: memcached,
		queue:     queue,
	}
}

func (s *ItemService) GetItem(id string) (dto.ItemDto, e.ApiError) {

	var itemDto dto.ItemDto
	itemDto, err := s.memcached.GetItem(id)
	if err == nil {
		log.Debug("memcached")
		return itemDto, nil
	}

	log.Debug("Error getting item from memcached")
	itemDto, err = s.item.GetItem(id)
	if err != nil {
		log.Debug("Error getting item from mongo")
		return itemDto, err
	}

	if itemDto.ItemId == "000000000000000000000000" {
		return itemDto, e.NewBadRequestApiError("item not found")
	}

	_, err = s.memcached.InsertItem(itemDto)
	if err != nil {
		log.Debug("Error inserting in memcached")
	}
	log.Debug("mongo")
	return itemDto, nil

}

func (s *ItemService) InsertItem(itemDto dto.ItemDto) (dto.ItemDto, e.ApiError) {

	var insertItem dto.ItemDto

	insertItem, err := s.item.InsertItem(itemDto)
	if err != nil {
		return itemDto, e.NewBadRequestApiError("error inserting item")
	}

	if insertItem.ItemId == "000000000000000000000000" {
		return itemDto, e.NewBadRequestApiError("error in insert")
	}
	itemDto.ItemId = insertItem.ItemId

	itemDto, err2 := s.memcached.InsertItem(itemDto)
	if err2 != nil {
		return itemDto, e.NewBadRequestApiError("Error inserting in memcached")
	}
	return itemDto, nil
}

func (s *ItemService) QueueItems(itemsDto dto.ItemsDto) e.ApiError {
	for i := range itemsDto {
		item := itemsDto[i]
		go func() {
			item, err := s.item.InsertItem(item)
			if err != nil {
				log.Debug(err)
			}
			err = s.queue.SendMessage("solr", item.ItemId) //se envia a solr el id de cada elemento que se añade a mongodb
			log.Debug(err)
		}()
	}
	return nil
}
