package pubsub

import (
	"context"
	"log"
)

type DashboardHandler struct{}

func (h *DashboardHandler) Handle(ctx context.Context, message string) error {
	log.Printf("Dashboard received: %s", message)
	return nil
}
