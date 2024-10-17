package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func main() {
	// Track the last attempt time
	var previousAttempt int64 = time.Now().UnixMilli()

	// The operation to be retried
	operation := func() error {
		currentAttempt := time.Now().UnixMilli()
		fmt.Printf("Attempting operation... time since last attempt: %d ms\n", currentAttempt-previousAttempt)
		previousAttempt = currentAttempt
		// Simulate failure
		return errors.New("service unavailable")
	}

	// Set up exponential backoff with configuration
	b := backoff.NewExponentialBackOff(
		backoff.WithInitialInterval(100*time.Millisecond),
		backoff.WithMaxElapsedTime(30*time.Second),
		backoff.WithMaxInterval(5*time.Second),
		backoff.WithMultiplier(1.5),
		backoff.WithRandomizationFactor(0.1),
	)

	// Attempt to retry the operation using the backoff strategy
	err := backoff.Retry(operation, b)
	if err != nil {
		fmt.Println("Operation failed after retries:", err)
	}
}
