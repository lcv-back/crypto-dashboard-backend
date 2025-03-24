package main

import (
	"context"
	"log"

	"github.com/lcv-back/crypto-dashboard-backend/internal/pubsub"
)

func main() {
	ctx := context.Background()
	ps, err := pubsub.InitRedisPubSub(ctx, "localhost:6379", "dashboard")
	if err != nil {
		log.Fatalf("Failed to init Redis Pub/Sub: %v", err)
	}
	defer ps.Close()

	select {}
}
