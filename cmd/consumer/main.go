package main

import (
	kps "KafkaProducerConsumer/pkg"
	"context"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking

	kps.Consume(ctx)
}


