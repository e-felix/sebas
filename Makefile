SOURCE_DIR=cli
ENTRY_FILE=main.go
DEST_DIR=build
BIN=sebas

.PHONY: install build test clean

all: install

install:
	go get -v ./...

build: $(SOURCE_DIR)/$(ENTRY_FILE)
	go build -o $(DEST_DIR)/$(BIN) -v $(SOURCE_DIR)/$(ENTRY_FILE)

test: install
	go test -v ./...

clean:
	rm -rf $(DEST_DIR)