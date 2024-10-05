package redispubsub

import (
	"context"

	"github.com/go-redis/redis/v8"
	messagebroker "github.com/khainv198/golang-message-broker"
)

type subscriber struct {
	rc *redis.Client
}

func newSubscriber(rc *redis.Client) messagebroker.Subscriber {
	return &subscriber{rc: rc}
}

func (s *subscriber) Subscribe(ctx context.Context, subject string, handler messagebroker.HandlerFunc) error {
	sub := s.rc.Subscribe(ctx, subject)

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			return err
		}

		err = handler(ctx, subject, msg.Payload)
		if err != nil {
			return err
		}
	}
}
