package adapterkafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	messagebroker "github.com/khainv198/golang-message-broker"
	"github.com/segmentio/kafka-go"
)

type subscriber struct {
	reader *kafka.Reader
}

func newSubscriber(cfg kafka.ReaderConfig) messagebroker.Subscriber {
	cfg.Dialer = &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	cfg.Logger = log.New(os.Stdout, fmt.Sprintf("kafka writer topic %s: ", cfg.Topic), 0)

	reader := kafka.NewReader(cfg)

	return &subscriber{reader: reader}
}

func (s subscriber) Subscribe(ctx context.Context, subject string, handler messagebroker.HandlerFunc) error {
	go func() {
		for {
			m, err := s.reader.FetchMessage(ctx)
			if err != nil {
				break
			}
			handler(ctx, subject, m.Value)
			if err := s.reader.CommitMessages(ctx, m); err != nil {
				log.Fatal("failed to commit messages:", err)
			}
		}
	}()

	return nil
}
