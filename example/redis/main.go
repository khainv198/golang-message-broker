package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	messagebroker "github.com/khainv198/golang-message-broker"
	redispubsub "github.com/khainv198/golang-message-broker/redis"
)

func main() {
	ctx := context.Background()
	channel := "messages"

	publisher, subscriber, err := redispubsub.New(ctx, &redis.Options{
		Addr: "localhost:6379",
	})
	if err != nil {
		log.Fatal("new redis publisher and subscriber error: ", err)
	}

	usecase := messagebroker.New(publisher, subscriber)

	go func() {
		for i := 0; i < 100; i++ {
			msg := fmt.Sprintf("hello redis pub sub from message id: %d", i)
			err := usecase.Publish.Execute(ctx, channel, map[string]interface{}{"message": msg})
			if err != nil {
				log.Fatal("publish message to redis error: ", err)
			}

			log.Printf("published message id %d to redis ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	usecase.Subscribe.Execute(ctx, channel, func(ctx context.Context, subject string, data interface{}) error {
		log.Print("subscribe message data: ", data)
		return nil
	})
}
