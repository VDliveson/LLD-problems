package internal

import "log"

func (q *MessageQueue) Enqueue(message Message) {
	q.Messages = append(q.Messages, message)
}

func (q *MessageQueue) Dequeue() (Message, bool) {
	if len(q.Messages) == 0 {
		return Message{}, false
	}
	message := q.Messages[0]
	q.Messages = q.Messages[1:]
	return message, true
}

func (q *MessageQueue) notifySubscribers() {
	q.mu.Lock()
	messages := q.Messages
	q.Messages = nil
	q.mu.Unlock()

	for _, subscriber := range q.Subscribers {
		go func() {
			batchSize := subscriber.BatchSize
			if batchSize <= 0 {
				batchSize = 1
			}

			for i := 0; i < len(messages); i += batchSize {
				end := i + batchSize
				if end > len(messages) {
					end = len(messages)
				}
				batch := messages[i:end]
				val := ExponentialBackoff(0, func() error {
					return subscriber.Callback(batch)
				})
				if val == -1 {
					log.Printf("Failed to process batch %v for subscriber %s after multiple attempts", (i / batchSize), subscriber.ID)
				}
			}
		}()
	}

}
