BIN=Sebas
DIST=build

.PHONY: install test

all: app_run

install:
	go get ./...

test: install
	go test ./...

# Gui related
.PHONY: app_cmd app_build app_run clean

app_cmd:
	go install cogentcore.org/core/cmd/core@main

app_build: install
	core build -o ${DIST} -name ${BIN}

app_run: app_cmd install
	core run -o ${DIST} -name ${BIN}

clean:
	rm -rf ./${DIST}
