package messaging

import (
	"auth-service/internal/messaging/events"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	Conn *nats.Conn
}

func (p *Publisher) PublishUserCreated(
	event events.UserCreatedEvent,
) error {
	data, err := json.Marshal(event)

	if err != nil {
		return err
	}

	return p.Conn.Publish(
		"user_created",
		data,
	)
}