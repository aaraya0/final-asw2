package dtos

type MessageDto struct {
	MessageId   int    `json:"message_id"`
	UserId      int    `json:"user_id"`
	Body        string `json:"body"`
	ItemId      string `json:"item_id"`
	CreatedAt   string `json:"created_at"`
	Transmitter string `json:"transmitter"`
}

type MessagesDto []MessageDto
