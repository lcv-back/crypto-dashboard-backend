package pubsub

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

// TaskPayload là dữ liệu gửi vào hàng đợi
type TaskPayload struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

// NewTask tạo task mới cho Asynq
func NewTask(topic, message string) (*asynq.Task, error) {
	payload, err := json.Marshal(TaskPayload{Topic: topic, Message: message})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(topic, payload), nil
}
