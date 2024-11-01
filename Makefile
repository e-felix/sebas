BIN=sebas

.PHONY: install test

all: install

install:
	go get -v ./...

test: install
	go test -v ./...
