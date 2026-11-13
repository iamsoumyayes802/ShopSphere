package messaging

import (
    "log"

    "github.com/nats-io/nats.go"
)

type NATSClient struct {
	Conn *nats.Conn
}

func NewNATSClient(url string) (*NATSClient, error) {
	nc, err := nats.Connect(
		url,
		nats.Name("auth-service"),
	)

	if err != nil {
		return nil, err
	}

	log.Println("Connected to NATS")

	return &NATSClient{
		Conn: nc,
	}, nil
}

func (n *NATSClient) Close() {
	n.Conn.Drain()
}