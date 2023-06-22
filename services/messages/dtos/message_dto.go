package dtos

type ItemDto struct {
	ItemId      string `json:"_id"`
	Titulo      string `json:"title"`
	Ubicacion   string `json:"location"`
	Vendedor    string `json:"seller"`
	Descripcion string `json:"description"`
	Clase       string `json:"class"`
	Mts2        int    `json:"mts2"`
	Precio      int    `json:"price"`
	Imagen      string `json:"img_url"`
}
type ItemsDto []ItemDto

type MessageDto struct {
	MessageId int    `json:"message_id"`
	UserId    int    `json:"user_id"`
	Body      string `json:"body"`
	ItemId    string `json:"item_id"`
	CreatedAt string `json:"created_at"`
	System    bool   `json:"system"`
}

type MessagesDto []MessageDto
