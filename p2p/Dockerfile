# syntax=docker/dockerfile:1
FROM golang:1.21

WORKDIR /app
ADD . /app

COPY go.mod ./

RUN go build -o p2p
CMD ["./p2p"]
