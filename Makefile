export CGO_ENABLED=1
export GO111MODULE=on

setup:
	go get github.com/jackc/pgx/v5

build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o cmd/app/main.go

#docker-build:
#	CGO_ENABLED=1 docker build -t go-project-core .

compose-up:
	docker compose up -d --build

compose-down:
	docker compose down

all: setup build compose-up