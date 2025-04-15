export CGO_ENABLED=1
export GO111MODULE=on

setup:
	go get github.com/mattn/go-sqlite3

install:
	sudo apt-get update
	sudo apt-get install sqlite

build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o cmd/app/main.go

docker-build:
	CGO_ENABLED=1 docker build -t go-project-core .

all: setup install build docker-build