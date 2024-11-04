BIN=sebas

.PHONY: install test build run

all: run

install:
	go get ./...  && \
	go mod tidy

test: install
	go test ./...

build: install
	go build

run: install
	go run -v .