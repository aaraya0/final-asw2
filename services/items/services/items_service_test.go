package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"

	dto "github.com/aaraya0/final-asw2/services/items/dtos"
	"github.com/aaraya0/final-asw2/services/items/model"
)

type ItemClientInterface struct {
	mock.Mock
}

func (m *ItemClientInterface) GetItem(id string) model.Item {
	ret := m.Called(id)
	return ret.Get(0).(model.Item)
}

// ... Adaptar otros métodos del cliente si es necesario ...

func TestGetItem(t *testing.T) {
	mockItemClient := new(ItemClientInterface)
	mockMemcachedClient := new(MemcachedClientInterface)
	mockQueueClient := new(QueueClientInterface)

	var item model.Item
	item.ItemId = "1"
	item.Titulo = "Test_Item"
	item.Ubicacion = "Test_Ubicacion"
	// ... Agregar otros campos ...

	var itemDto dto.ItemDto
	itemDto.ItemId = "1"
	itemDto.Titulo = "Test_Item"
	itemDto.Ubicacion = "Test_Ubicacion"
	// ... Agregar otros campos ...

	mockItemClient.On("GetItem", "1").Return(item)
	mockMemcachedClient.On("GetItem", "1").Return(itemDto)
	service := NewItemService(mockItemClient, mockMemcachedClient, mockQueueClient)
	res, err := service.GetItem("1")
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, res, itemDto)
}

// ... Adaptar otros casos de prueba según sea necesario ...
