package adapternats

import (
	"context"

	messagebroker "github.com/khainv198/golang-message-broker"
	"github.com/nats-io/nats.go"
)

type subscriber struct {
	nc *nats.Conn
}

func newSubscriber(nc *nats.Conn) messagebroker.Subscriber {
	return &subscriber{nc: nc}
}

func (s subscriber) Subscribe(ctx context.Context, subject string, handler messagebroker.HandlerFunc) error {
	_, err := s.nc.Subscribe(subject, func(msg *nats.Msg) {
		handler(ctx, subject, msg.Data)
	})
	return err
}
