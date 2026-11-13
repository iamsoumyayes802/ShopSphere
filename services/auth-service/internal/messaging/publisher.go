package messaging

import (
	"auth-service/internal/events"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	Conn *nats.Conn
}

func NewPublisher(conn *nats.Conn) *Publisher {
	return &Publisher{
		Conn: conn,
	}
}

func (p *Publisher) Publish(
	subject string,
	event events.UserCreatedEvent,
) error {
	data, err := json.Marshal(event)

	if err != nil {
		return err
	}

	return p.Conn.Publish(
		subject,
		data,
	)
}