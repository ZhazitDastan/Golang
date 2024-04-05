package main

import (
	"context"
	"testing"
	"time"
)

func TestQueryDatabase(t *testing.T) {
	t.Run("Complete successfully", func(t *testing.T) {
		ctx := context.Background()
		err := QueryDatabase(ctx)
		if err != nil {
			t.Errorf("QueryDatabase returned an unexpected error: %v", err)
		}
	})

	t.Run("Times out", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		err := QueryDatabase(ctx)
		if err == nil {
			t.Error("QueryDatabase didn't return an error.")
		}
	})
}
