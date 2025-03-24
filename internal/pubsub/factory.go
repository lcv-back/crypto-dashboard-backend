package pubsub

import "fmt"

// HandlerFactory là struct quản lý việc tạo handler
type HandlerFactory struct {
	creators map[string]func() Handler
}

// NewHandlerFactory tạo một factory mới
func NewHandlerFactory() *HandlerFactory {
	f := &HandlerFactory{
		creators: make(map[string]func() Handler),
	}
	// Đăng ký các handler
	f.creators["dashboard"] = func() Handler { return &DashboardHandler{} }
	return f
}

// Create tạo handler dựa trên tên topic
func (f *HandlerFactory) Create(name string) (Handler, error) {
	creator, ok := f.creators[name]
	if !ok {
		return nil, fmt.Errorf("unknown topic: %s", name)
	}
	return creator(), nil
}

// Hàm tiện ích (dạng Factory Method đơn giản như mã gốc)
func CreateHandler(name string) (Handler, error) {
	factory := NewHandlerFactory()
	return factory.Create(name)
}
