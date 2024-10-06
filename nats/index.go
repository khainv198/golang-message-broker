package adapternats

import (
	messagebroker "github.com/khainv198/golang-message-broker"
	"github.com/nats-io/nats.go"
)

func New(url string) (messagebroker.Publisher, messagebroker.Subscriber, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, nil, err
	}

	pub := newPublisher(nc)
	sub := newSubscriber(nc)

	return pub, sub, nil
}
