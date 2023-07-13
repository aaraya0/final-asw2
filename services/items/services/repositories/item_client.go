package repositories

import (
	"context"
	"fmt"

	dto "github.com/aaraya0/final-asw2/services/items/dtos"
	"github.com/aaraya0/final-asw2/services/items/model"
	e "github.com/aaraya0/final-asw2/services/items/utils/errors"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ItemClient struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection string
}

func NewItemInterface(host string, port int, collection string) *ItemClient {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://root:root@%s:%d/?authSource=admin&authMechanism=SCRAM-SHA-256", host, port)))
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	names, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(fmt.Sprintf("Error getting dbs from MongoDB: %v", err))
	}

	fmt.Println("[MongoDB] Initialized connection")
	fmt.Println(fmt.Sprintf("[MongoDB] Available databases: %s", names))

	return &ItemClient{
		Client:     client,
		Database:   client.Database("publicaciones"),
		Collection: collection,
	}
}

func (s *ItemClient) GetItem(id string) (dto.ItemDto, e.ApiError) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.ItemDto{}, e.NewBadRequestApiError(fmt.Sprintf("error getting item %s invalid id", id))
	}
	result := s.Database.Collection(s.Collection).FindOne(context.TODO(), bson.M{
		"_id": objectID,
	})
	if result.Err() == mongo.ErrNoDocuments {
		return dto.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	var item model.Item
	if err := result.Decode(&item); err != nil {
		return dto.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}
	return dto.ItemDto{
		ItemId:      id,
		Titulo:      item.Titulo,
		Ubicacion:   item.Ubicacion,
		Vendedor:    item.Vendedor,
		Descripcion: item.Descripcion,
		Mts2:        item.Mts2,
		Imagen:      item.Imagen,
		Precio:      item.Precio,
		Clase:       item.Clase,
		UsuarioId:   item.UsuarioId,
	}, nil

}

func (s *ItemClient) InsertItem(item dto.ItemDto) (dto.ItemDto, e.ApiError) {
	result, err := s.Database.Collection(s.Collection).InsertOne(context.TODO(), model.Item{
		ItemId:      primitive.NewObjectID(),
		Titulo:      item.Titulo,
		Ubicacion:   item.Ubicacion,
		Vendedor:    item.Vendedor,
		Descripcion: item.Descripcion,
		Clase:       item.Clase,
		Mts2:        item.Mts2,
		Precio:      item.Precio,
		Imagen:      item.Imagen,
		UsuarioId:   item.UsuarioId,
	})

	if err != nil {
		return item, e.NewInternalServerApiError(fmt.Sprintf("error inserting to mongo %s", item.ItemId), err)
	}

	item.ItemId = result.InsertedID.(primitive.ObjectID).Hex()

	return item, nil
}

func (s *ItemClient) GetItemsByUId(id int) (dto.ItemsDto, e.ApiError) {
	result, err := s.Database.Collection(s.Collection).Find(context.TODO(), bson.D{{Key: "usuario_id", Value: id}})
	if err != nil {
		return dto.ItemsDto{}, e.NewInternalServerApiError(fmt.Sprintf("error executing query: %v", err), err)
	}

	if result.Err() == mongo.ErrNoDocuments {
		return dto.ItemsDto{}, e.NewNotFoundApiError(fmt.Sprintf("user %d not found", id))
	}

	var items model.Items
	if err := result.All(context.TODO(), &items); err != nil {
		return dto.ItemsDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %d", id), err)
	}

	var itemsDto dto.ItemsDto
	for i := range items {
		item := items[i]
		itemsDto = append(itemsDto,
			dto.ItemDto{
				ItemId:      item.ItemId.Hex(),
				Titulo:      item.Titulo,
				Ubicacion:   item.Ubicacion,
				Vendedor:    item.Vendedor,
				Descripcion: item.Descripcion,
				Clase:       item.Clase,
				Mts2:        item.Mts2,
				Precio:      item.Precio,
				Imagen:      item.Imagen,
				UsuarioId:   item.UsuarioId,
			})
	}

	return itemsDto, nil
}

func (s *ItemClient) DeleteItem(id string) e.ApiError {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return e.NewBadRequestApiError(fmt.Sprintf("error deleting item %s invalid id", id))
	}

	result, err := s.Database.Collection(s.Collection).DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		log.Error(err)
		return e.NewInternalServerApiError("error deleting item", err)
	}
	log.Debug(result.DeletedCount)

	result, err = s.Database.Collection(s.Collection).DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		log.Error(err)
		return e.NewInternalServerApiError("error deleting item", err)
	}
	log.Debug(result.DeletedCount)
	return nil
}
