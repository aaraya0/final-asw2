package repositories

import (
	e "github.com/aaraya0/arq-software/final-asw2/worker/utils/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type QueueClient struct {
	Connection *amqp.Connection
}

func (qc *QueueClient) ProcessMessages(qname string, process func(string)) e.ApiError {
	channel, err := qc.Connection.Channel()
	message, err := channel.Consume(qname,
		"items",
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		return e.NewBadRequestApiError("Failed to register a consumer")
	}

	var forever chan struct{}
	go func() {
		for d := range message {
			process(string(d.Body))
		}
	}()
	<-forever
	return nil
}
