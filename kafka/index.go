package adapterkafka

import (
	messagebroker "github.com/khainv198/golang-message-broker"
	"github.com/segmentio/kafka-go"
)

func New(writerCfg kafka.WriterConfig, readerCfg kafka.ReaderConfig) (messagebroker.Publisher, messagebroker.Subscriber) {
	pub := newPublisher(writerCfg)
	sub := newSubscriber(readerCfg)

	return pub, sub
}
