package main

import (
	"fmt"
	"log"
	internal "message-queue/internal"
	"net/http"
)

func main() {
	queue := internal.NewMessageQueueService()
	err := queue.CreateQueue("orders")
	if err != nil {
		panic(err)
	}

	// Subscriber 1
	queue.Subscribe("orders", "sub-1", 2, func(msgs []internal.Message) error {
		log.Println("Sub1 received:")
		for _, m := range msgs {
			log.Println(string(m.Payload))
		}
		return nil
	})

	// Subscriber 2 (with failure simulation)
	queue.Subscribe("orders", "sub-2", 3, func(msgs []internal.Message) error {
		log.Println("Sub2 processing batch...")
		return fmt.Errorf("failed")
	})

	// Publish messages
	for i := 0; i < 1; i++ {
		log.Println("Published message")
		queue.Publish("orders", fmt.Sprintf(`{"order_id": %d}`, i))
	}
	http.ListenAndServe(":8080", nil)
}
