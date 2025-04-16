FROM golang:1.23.6-alpine AS builder

COPY . /usr/src/app/goproject
WORKDIR /usr/src/app/goproject

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /usr/local/bin/app ./cmd/app/main.go


CMD ["app"]