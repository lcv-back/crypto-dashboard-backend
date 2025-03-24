package pubsub

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

func RunWorker(ctx context.Context, redisAddr string) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc("dashboard", processTask)
	mux.HandleFunc("whale", processTask)
	mux.HandleFunc("portfolio", processTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("Could not run worker: %v", err)
	}
}

func processTask(ctx context.Context, task *asynq.Task) error {
	var payload TaskPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	handler, err := CreateHandler(payload.Topic)
	if err != nil {
		return err
	}

	return handler.Handle(ctx, payload.Message)
}
