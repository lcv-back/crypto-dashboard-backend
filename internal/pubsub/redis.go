package pubsub

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

type RedisPubSub struct {
	pubsubClient *redis.Client
	asynqClient  *asynq.Client
	handler      Handler
}

func InitRedisPubSub(ctx context.Context, redisAddr, name string) (*RedisPubSub, error) {
	pubsubClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	if err := pubsubClient.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})

	handler, err := CreateHandler(name)
	if err != nil {
		return nil, err
	}

	ps := &RedisPubSub{
		pubsubClient: pubsubClient,
		asynqClient:  asynqClient,
		handler:      handler,
	}

	go ps.subscribe(ctx, name)
	return ps, nil
}

func (ps *RedisPubSub) subscribe(ctx context.Context, topic string) {
	pubsub := ps.pubsubClient.Subscribe(ctx, topic)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			return
		}

		task, err := NewTask(topic, msg.Payload)
		if err != nil {
			log.Printf("Error creating task: %v", err)
			continue
		}

		if _, err := ps.asynqClient.Enqueue(task); err != nil {
			log.Printf("Error enqueuing task: %v", err)
		}
	}
}

func (ps *RedisPubSub) Close() {
	ps.pubsubClient.Close()
	ps.asynqClient.Close()
}
