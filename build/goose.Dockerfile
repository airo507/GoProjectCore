FROM golang:1.24

WORKDIR /migrations

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD goose
