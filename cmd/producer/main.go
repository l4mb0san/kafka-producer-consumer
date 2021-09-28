package main

import (
	kps "KafkaProducerConsumer/pkg"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"io"
	"net/http"
	"strings"
)

var (
	ctx = context.Background()
	kw = &kafka.Writer{
		Addr:         kafka.TCP(kps.Address),
		Topic:        kps.Topic,
	}
)

func main() {

	port := ":8989"
	r := mux.NewRouter()

	// Special APIs
	r.HandleFunc("/", Handler).Methods(http.MethodPost)

	fmt.Println("Server is running on port 8989")
	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println("Main: Error http.ListenAndServe")
		return
	}
}

func Handler(w http.ResponseWriter, req *http.Request) {
	var (
		buf = new(strings.Builder)
		cont = kps.Content{}
	)

	_, _ = io.Copy(buf, req.Body)

	_ = json.Unmarshal([]byte(buf.String()), &cont)
	fmt.Println(cont)

	go kps.Produce(ctx, kw, &cont)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"OK\"}"))
}

