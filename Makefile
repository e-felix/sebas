BIN=sebas

.PHONY: install test

all: app_run

install:
	go get ./...

test: install
	go test ./...

# Gui related
.PHONY: app_cmd app_build app_run

app_cmd:
	go install cogentcore.org/core/cmd/core@main

app_build: install
	core build

app_run: app_cmd install
	core run
