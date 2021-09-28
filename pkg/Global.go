package pkg

import "os"

var (
	Topic   = os.Getenv("KAFKA_TOPIC")
	Address = os.Getenv("KAFKA_ADDR")
	Group   = os.Getenv("CONSUMER_GROUP")
)

type Content struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
