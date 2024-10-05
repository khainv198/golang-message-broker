package messagebroker

import (
	"context"
)

type Publisher interface {
	Publish(ctx context.Context, subject string, data interface{}) error
}

type publishMessageUsecase struct {
	Publisher Publisher
}

type PublishMessageUsecase interface {
	Execute(ctx context.Context, subject string, data interface{}) error
}

func (uc publishMessageUsecase) Execute(ctx context.Context, subject string, data interface{}) error {
	return uc.Publisher.Publish(ctx, subject, data)
}

func newPublishMessageUsecase(publisher Publisher) PublishMessageUsecase {
	return &publishMessageUsecase{Publisher: publisher}
}

type HandlerFunc func(ctx context.Context, subject string, data interface{}) error

type Subscriber interface {
	Subscribe(ctx context.Context, subject string, handler HandlerFunc) error
}

type subscribeMessageUsecase struct {
	Subscriber Subscriber
}

type SubscribeMessageUsecase interface {
	Execute(ctx context.Context, subject string, handler HandlerFunc) error
}

func (uc *subscribeMessageUsecase) Execute(ctx context.Context, subject string, handler HandlerFunc) error {
	return uc.Subscriber.Subscribe(ctx, subject, handler)
}

func newSubscribeMessageUsecase(subscriber Subscriber) SubscribeMessageUsecase {
	return &subscribeMessageUsecase{Subscriber: subscriber}
}

type MessageBrokerUsecase struct {
	Publish   PublishMessageUsecase
	Subscribe SubscribeMessageUsecase
}

func New(publisher Publisher, subscriber Subscriber) *MessageBrokerUsecase {
	return &MessageBrokerUsecase{
		Publish:   newPublishMessageUsecase(publisher),
		Subscribe: newSubscribeMessageUsecase(subscriber),
	}
}
