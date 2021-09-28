package pkg

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func Produce(ctx context.Context, w *kafka.Writer, cont *Content) {

	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(cont.Key),
		Value: []byte(cont.Value),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

	// log a confirmation once the message is written
	fmt.Println("key:", string(cont.Key), "Value:", string(cont.Value))
}
