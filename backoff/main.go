package main

import (
	"errors"
	"fmt"
	backoff "github.com/cenkalti/backoff/v4"
	"time"
)

func main() {
	var pre int64 = time.Now().Unix()
	operation := func() error {
		fmt.Println("Attempting operation...", time.Now().Unix(), time.Now().Unix()-pre)
		pre = time.Now().Unix()
		return errors.New("some error")
	}

	err := backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		fmt.Println("Operation failed after retries:", err)
	}
}
