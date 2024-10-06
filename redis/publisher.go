package adapterredis

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	messagebroker "github.com/khainv198/golang-message-broker"
)

type publisher struct {
	rc *redis.Client
}

func newPublisher(rc *redis.Client) messagebroker.Publisher {
	return &publisher{rc: rc}
}

func (p publisher) Publish(ctx context.Context, subject string, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.rc.Publish(ctx, subject, string(dataBytes)).Err()
}
