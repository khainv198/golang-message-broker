package adapterkafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	messagebroker "github.com/khainv198/golang-message-broker"
	"github.com/segmentio/kafka-go"
)

type publisher struct {
	writer *kafka.Writer
}

func newPublisher(cfg kafka.WriterConfig) messagebroker.Publisher {
	cfg.Logger = log.New(os.Stdout, fmt.Sprintf("kafka writer topic %s: ", cfg.Topic), 0)

	cfg.Dialer = &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	cfg.Async = false

	writer := kafka.NewWriter(cfg)

	return &publisher{writer: writer}
}

func (p publisher) Publish(ctx context.Context, subject string, data interface{}) error {
	var value []byte
	if v, ok := data.(string); ok {
		value = []byte(v)
	} else {
		v, err := json.Marshal(data)
		if err != nil {
			return err
		}

		value = v
	}

	msg := kafka.Message{
		Key:   []byte(uuid.NewString()),
		Value: value,
		Topic: subject,
	}

	err := p.writer.WriteMessages(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}
