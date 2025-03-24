package main

import (
	"context"

	"github.com/lcv-back/crypto-dashboard-backend/internal/pubsub"
)

func main() {
	ctx := context.Background()
	pubsub.RunWorker(ctx, "localhost:6379")
}
