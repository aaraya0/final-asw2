package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ItemId      primitive.ObjectID `bson:"_id"`
	Titulo      string             `bson:"title"`
	Ubicacion   string             `bson:"location"`
	Vendedor    string             `bson:"seller"`
	Descripcion string             `bson:"description"`
	Clase       string             `bson:"class"`
	Mts2        int                `bson:"mts2"`
	Precio      int                `bson:"price"`
	Imagen      string             `bson:"img_url"`
}
type Items []Item
