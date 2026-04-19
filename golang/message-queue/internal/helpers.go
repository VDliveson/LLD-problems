package internal

import (
	"log"
	"time"
)

func ExponentialBackoff(attempt int, callback func() error) int {
	const maxRetries = 3
	if attempt > maxRetries {
		return -1
	}
	err := callback()
	if err != nil {
		log.Printf("Attempt %d failed: %v. Retrying...", attempt+1, err)
		time.Sleep(time.Duration(1<<attempt) * time.Second)
		return ExponentialBackoff(attempt+1, callback)
	}
	return attempt
}
