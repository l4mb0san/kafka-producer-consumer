package pkg

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{Address},
		Topic:    Topic,
		GroupID:  Group,
	})

	for {
		msg, err := r.FetchMessage(ctx)
		if err != nil {
			fmt.Println("error fetching message:", err)
			return
		}
		if err = r.CommitMessages(context.Background(), msg); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("received:", string(msg.Value))
	}
}
