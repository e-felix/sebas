BIN=sebas
DIST=build

.PHONY: install test

all: app_run

install:
	go get ./...

test: install
	go test ./...

# Gui related
.PHONY: app_cmd app_build app_run clean build_dir

app_cmd:
	go install cogentcore.org/core/cmd/core@main

app_build: app_cmd install
	core build

app_run: app_cmd install
	core run


clean:
	rm -rf ${DIST}

