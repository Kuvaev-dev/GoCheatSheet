// Cpntext is a powerful tool in Go for managing deadlines, cancellation signals,
// and request-scoped values across API boundaries and between processes.
// It allows developers to control the lifecycle of operations, ensuring that resources are properly
// released when they are no longer needed. In this example, we will explore various ways
// to use context in Go, including creating different types of contexts,
// handling cancellation, and passing values through contexts.

package main

import (
	"context"
	"time"
)

func contextFund() {
	// 1. Empty Context
	ctx := context.Background()
	println(ctx)

	// 2. Unknown Context
	unknownCtx := context.TODO()
	println(unknownCtx)

	// 3. Context with Cancel
	ctxWithCancel, cancel := context.WithCancel(context.Background())
	println(ctxWithCancel)
	cancel() // cancel the context

	// 4. Context with Timeout
	ctxWithTimeout, cancelTimeout := context.WithTimeout(context.Background(), 0)
	println(ctxWithTimeout)
	cancelTimeout() // cancel the context

	// 5. Context with Deadline
	ctxWithDeadline, cancelDeadline := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	println(ctxWithDeadline)
	cancelDeadline() // cancel the context

	// 6. Context with Value
	ctxWithValue := context.WithValue(context.Background(), "key", "value")
	println(ctxWithValue.Value("key"))

	// 7. Context Cancellation Reaction
	err := contextCancelReaction(ctxWithCancel)
	if err != nil {
		println("Context cancelled:", err)
	}

	// 8. Send value through context
	contextWithValue(context.Background(), "greeting", "Hello, Context!")
}

// 7. Context Cancellation Reaction
func contextCancelReaction(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // return the error from the context
	default:
		// do some work
		return nil
	}
}

// 8. Send value through context
func contextWithValue(ctx context.Context, key, value string) {
	ctxWithValue := context.WithValue(ctx, key, value)
	println(ctxWithValue.Value(key))
}
