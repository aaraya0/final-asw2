package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aaraya0/final-asw2/services/items/config"
	dto "github.com/aaraya0/final-asw2/services/items/dtos"
	client "github.com/aaraya0/final-asw2/services/items/services/repositories"
	e "github.com/aaraya0/final-asw2/services/items/utils/errors"
	json "github.com/json-iterator/go"
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
	err = s.queue.SendMessage(insertItem.ItemId, "create", insertItem.ItemId)
	log.Debug(err)
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
			err = s.queue.SendMessage(item.ItemId, "create", item.ItemId)
			log.Debug(err)
		}()
	}
	return nil
}
func (s *ItemService) DeleteUserItems(id int) e.ApiError {
	items, err := s.GetItemsByUId(id)
	if err != nil {
		log.Error(err)
		return err
	}
	for i := range items {
		var item dto.ItemDto
		item = items[i]
		go func() {
			err := s.item.DeleteItem(item.ItemId)
			if err != nil {
				log.Error(err)
			}
			err = s.queue.SendMessage(item.ItemId, "delete", fmt.Sprintf("%s.delete", item.ItemId))
			log.Error(err)
		}()
	}
	return nil
}

func (s *ItemService) DeleteItem(id string) e.ApiError {

	err := s.item.DeleteItem(id)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.memcached.Delete(id)
	if err != nil {
		log.Error("Error deleting from cache", err)
	}
	err = s.queue.SendMessage(id, "delete", fmt.Sprintf("%s.delete", id))
	log.Error(err)

	return nil
}

func (s *ItemService) GetItemsByUId(id int) (dto.ItemsDto, e.ApiError) {

	var itemsDto dto.ItemsDto

	var itemsRespDto dto.ItemsDto

	itemsDto, err := s.item.GetItemsByUId(id)
	if err != nil {
		log.Debug("Error getting items from mongo")
		return itemsDto, err
	}

	for i := range itemsDto {
		item, err := s.GetUserById(itemsDto[i].UsuarioId, itemsDto[i])
		if err != nil {
			return itemsDto, e.NewBadRequestApiError("error getting user for item")
		}
		itemsDto[i].Usuario = item.Usuario
		itemsDto[i].UNombre = item.UNombre
		itemsDto[i].UApellido = item.UApellido
		itemsDto[i].UEmail = item.UEmail
		itemsRespDto = append(itemsRespDto, itemsDto[i])
	}

	return itemsRespDto, nil
}

func (s *ItemService) GetUserById(id int, itemDto dto.ItemDto) (dto.ItemDto, e.ApiError) {
	var userDto dto.UserDto
	var itemRDto dto.ItemDto

	var er e.ApiError
	er = nil

	userDto, err := s.memcached.GetUserById(id)
	if err != nil {
		resp, err := http.Get(fmt.Sprintf("http://%s:%d/%s/%d", config.USERSHOST, config.USERSPORT, config.USERSENDPOINT, id))
		if err != nil {
			return itemRDto, e.NewInternalServerApiError("Error getting user from user service", err)
		}
		err = json.NewDecoder(resp.Body).Decode(&userDto)
		if err != nil {
			return itemRDto, e.NewInternalServerApiError("Error decoding userDto", err)
		}

		userDto, err = s.memcached.InsertUser(userDto)
		if err != nil {
			er = e.NewInternalServerApiError("Error inserting user to memcached", err)
		}
	}

	itemRDto.Usuario = userDto.Username
	itemRDto.UNombre = userDto.FirstName
	itemRDto.UApellido = userDto.LastName
	itemRDto.UEmail = userDto.Email
	return itemRDto, er
}

func (s *ItemService) DownloadImage(id string) e.ApiError {

	itemDto, err := s.item.GetItem(id)

	if err != nil {
		return e.NewInternalServerApiError("error getting item", err)
	}

	// Obtener la imagen desde la URL

	resp, _ := http.Get(itemDto.Imagen)
	if err != nil {
		return e.NewInternalServerApiError("error downloading image", err)
	}
	defer resp.Body.Close()

	// Crear el archivo en la carpeta "images" con el nombre del ID del item y la extensión de la imagen
	filePath := filepath.Join("../../frontend/src/images", id+".png") // Puedes usar ".png" u otra extensión según la imagen que esperas recibir.

	// Crear el archivo
	file, _ := os.Create(filePath)
	if err != nil {
		return e.NewInternalServerApiError("error creating image file", err)
	}
	defer file.Close()

	// Copiar el contenido de la imagen descargada al archivo
	_, _ = io.Copy(file, resp.Body)
	if err != nil {
		return e.NewInternalServerApiError("error saving image to file", err)
	}

	return nil
}
