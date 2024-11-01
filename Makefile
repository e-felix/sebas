BIN=sebas

.PHONY: install test

all: install

install:
	go get ./...

test: install
	go test ./...
