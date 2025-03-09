FROM golang:1.24.1-alpine3.21 AS builder

COPY . /github.com/Ghaarp/chat-server/cmd
WORKDIR /github.com/Ghaarp/chat-server/cmd

RUN go mod download
RUN go build -o ./bin/chat_service cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Ghaarp/chat-server/cmd/bin/chat_service .
CMD ["./chat_service"]