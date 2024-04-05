package main

import (
	"context"
	"fmt"
	"time"
)

func QueryDatabase(ctx context.Context) error {

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Successful")
		return nil
	case <-ctx.Done():
		fmt.Println("Timeout")
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := QueryDatabase(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
