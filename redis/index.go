package adapterredis

import (
	"context"

	"github.com/go-redis/redis/v8"
	messagebroker "github.com/khainv198/golang-message-broker"
)

func New(ctx context.Context, opts *redis.Options) (messagebroker.Publisher, messagebroker.Subscriber, error) {
	client := redis.NewClient(opts)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, nil, err
	}

	pub := newPublisher(client)
	sub := newSubscriber(client)
	return pub, sub, nil
}
