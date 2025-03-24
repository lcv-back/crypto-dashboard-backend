package pubsub

import "context"

// Handler defines the general interface for handling topics
type Handler interface {
	Handle(ctx context.Context, message string) error
}
