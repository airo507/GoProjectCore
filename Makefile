export CGO_ENABLED=1
export GO111MODULE=on

LOCAL_BIN:=$(CURDIR)/bin

install-protoc:
	sudo apt install protobuf-compiler
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-protoc:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	#make install-protoc
	#make get-protoc
	#mkdir gen/blog
	protoc -I=grpc --go_out=./gen/blog/ --go_opt=paths=source_relative --go-grpc_out=./gen/blog/ --go-grpc_opt=paths=source_relative grpc/blog/blog.proto


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