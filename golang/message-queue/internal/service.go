package internal

import (
	"encoding/json"
	"fmt"
	"time"
)

func NewMessageQueueService() *MessageQueueService {
	return &MessageQueueService{
		queues: make(map[string]*MessageQueue),
	}
}

func (s *MessageQueueService) CreateQueue(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.queues[name]; exists {
		return nil
	}

	s.queues[name] = &MessageQueue{Name: name, Messages: []Message{}}
	return nil
}

func (s *MessageQueueService) Publish(queueName string, content string) error {
	s.mu.RLock()
	queue, exists := s.queues[queueName]
	s.mu.RUnlock()
	if !exists {
		return fmt.Errorf("queue %s does not exist", queueName)
	}

	queue.mu.Lock()
	message := Message{
		ID:        fmt.Sprintf("%s-%d", queueName, len(queue.Messages)+1),
		Topic:     queueName,
		Payload:   json.RawMessage(content),
		Timestamp: time.Now(),
	}
	queue.Messages = append(queue.Messages, message)
	queue.mu.Unlock()

	go queue.notifySubscribers()
	return nil
}

func (s *MessageQueueService) GetQueueLength(queueName string) (int, error) {
	s.mu.RLock()
	queue, exists := s.queues[queueName]
	s.mu.RUnlock()
	if !exists {
		return 0, fmt.Errorf("queue %s does not exist", queueName)
	}

	queue.mu.Lock()
	defer queue.mu.Unlock()
	return len(queue.Messages), nil
}

func (s *MessageQueueService) RemoveSubscriber(queueName string, subscriberID string) error {
	s.mu.RLock()
	queue, exists := s.queues[queueName]
	s.mu.RUnlock()
	if !exists {
		return fmt.Errorf("queue %s does not exist", queueName)
	}

	queue.mu.Lock()
	defer queue.mu.Unlock()
	delete(queue.Subscribers, subscriberID)
	return nil
}

func (s *MessageQueueService) Subscribe(queueName string, subscriberID string, batchSize int, callback func([]Message) error) error {
	s.mu.RLock()
	queue, exists := s.queues[queueName]
	s.mu.RUnlock()
	if !exists {
		return fmt.Errorf("queue %s does not exist", queueName)
	}

	queue.mu.Lock()
	defer queue.mu.Unlock()
	if queue.Subscribers == nil {
		queue.Subscribers = make(map[string]*Subscriber)
	}
	queue.Subscribers[subscriberID] = &Subscriber{
		ID:        subscriberID,
		BatchSize: batchSize,
		Callback:  callback,
	}
	return nil
}

func (s *MessageQueueService) ChangeBatchSize(queueName string, subscriberID string, newBatchSize int) error {
	s.mu.RLock()
	queue, exists := s.queues[queueName]
	s.mu.RUnlock()
	if !exists {
		return fmt.Errorf("queue %s does not exist", queueName)
	}

	queue.mu.Lock()
	defer queue.mu.Unlock()
	subscriber, exists := queue.Subscribers[subscriberID]
	if !exists {
		return fmt.Errorf("subscriber %s does not exist in queue %s", subscriberID, queueName)
	}
	subscriber.BatchSize = newBatchSize
	return nil
}
