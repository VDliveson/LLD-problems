package internal

import (
	"encoding/json"
	"sync"
	"time"
)

type Message struct {
	ID        string          `json:"id"`
	Topic     string          `json:"topic"`
	Payload   json.RawMessage `json:"payload"`
	Timestamp time.Time       `json:"timestamp"`
}

type MessageQueue struct {
	Name        string
	Messages    []Message
	Subscribers map[string]*Subscriber
	mu          sync.Mutex
}

type MessageQueueService struct {
	queues map[string]*MessageQueue
	mu     sync.RWMutex
}

type Subscriber struct {
	ID        string
	BatchSize int
	Callback  func([]Message) error
}
