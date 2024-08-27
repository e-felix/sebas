SOURCE_DIR=cli
ENTRY_FILE=main.go
DEST_DIR=build
BIN=sebas

.PHONY: clean

all:
	go mod tidy

install:
	go mod tidy

build: $(SOURCE_DIR)/$(ENTRY_FILE)
	go build -o $(DEST_DIR)/$(BIN) -v $(SOURCE_DIR)/$(ENTRY_FILE)

test:
	go test -v ./...

clean:
	rm -rf $(DEST_DIR)