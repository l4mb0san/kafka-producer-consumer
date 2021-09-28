# Dockerfile References: https://docs.docker.com/engine/reference/builder/
FROM golang:1.17.1 as builder
LABEL maintainer="TaiLeHuu"
RUN export PATH=$PATH:/go/bin

WORKDIR /app
COPY ../ .
RUN go mod download
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o main test/main.go


######## Start a new stage from scratch #######
FROM alpine:latest
RUN apk update
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add curl

WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT ["./main"]
