# kafka-producer-consumer
Create kafka-producer &amp; kafka-consumer on kubernetes

# References
- https://www.sohamkamani.com/golang/working-with-kafka/
- https://platform9.com/blog/how-to-set-up-and-run-kafka-on-kubernetes/
- https://sourcegraph.com/github.com/segmentio/kafka-go@df148e48c83fae30b4dd134617dfa5fcb55965b3/-/blob/client_test.go?L144

# Step 1: Setup & run kafka on kubernetes
Follow the instructions [here](https://platform9.com/blog/how-to-set-up-and-run-kafka-on-kubernetes/)

# Step 2: Clone this project
```
git clone https://github.com/l4mb0san/kafka-producer-consumer.git
cd kafka-producer-consumer
```

# Step 3: Build kafka-produce image
```
docker build -t <your-repository>/kafka-producer:<tag> -f cmd/producer/Dockerfile .
```
- Change containers.image at cmd/producer/deployment.yaml
```
    spec:
      containers:
        - name: producer
          image: <your-repository>/kafka-producer:<tag>
```
- Change env.KAFKA_TOPIC if you create with another name
```
          env:
            - name: KAFKA_TOPIC
              value: "messages"
```

# Step 4: Build kafka-consumer image
```
docker build -t <your-repository>/kafka-consumer:<tag> -f cmd/consumer/Dockerfile .
```
- Change containers.image at cmd/consumer/deployment.yaml
```
    spec:
      containers:
        - name: consumer
          image: <your-repository>/kafka-consumer:<tag>
```
- Change env.KAFKA_TOPIC & env.CONSUMER_GROUP if you create with another name
```
          env:
            - name: KAFKA_TOPIC
              value: "messages"
            - name: CONSUMER_GROUP
              value: "securecode"
```

# Step 5: Apply to k8s
```
kubectl apply -f cmd/producer/deployment.yaml -f cmd/producer/service.yaml
kubectl apply -f cmd/consumer/deployment.yaml
```

# Step 6: Test
```
kubectl port-forward svc/l4mb0-kafka-producer 8989

curl -XPOST http://localhost:8989 -d '{"key":"cat","value":"meow moew"}
```
