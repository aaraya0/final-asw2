package dto

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
	UsuarioId   int    `json:"usuario_id"`
	Usuario     string `json:"usuario"`
	UNombre     string `json:"usuario_nombre"`
	UApellido   string `json:"usuario_apellido"`
	UEmail      string `json:"usuario_email"`
}
type ItemsDto []ItemDto
