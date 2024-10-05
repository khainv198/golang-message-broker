package natspubsub

import (
	"context"
	"encoding/json"

	messagebroker "github.com/khainv198/golang-message-broker"
	"github.com/nats-io/nats.go"
)

type publisher struct {
	nc *nats.Conn
}

func newPublisher(nc *nats.Conn) messagebroker.Publisher {
	return &publisher{nc: nc}
}

func (p *publisher) Publish(ctx context.Context, subject string, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.nc.Publish(subject, dataBytes)
}
