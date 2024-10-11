CC=gcc
CFLAGS=-std=c99 -pedantic -Wall
BIN=sebas
DEST_FOLDER=bin
SRC=main.c src/sebas.c src/project.c src/database.c
LIBS=lib/database/sqlite.c
CLIBS=-lsqlite3

.PHONY: clean build run dir

all: run

dir:
	mkdir -p $(DEST_FOLDER)

build: dir $(SRC) $(LIBS)
	$(CC) $(SRC) $(LIBS) $(CFLAGS) $(CLIBS) -o $(DEST_FOLDER)/$(BIN)

run: build
	./$(DEST_FOLDER)$(BIN)

clean:
	rm -rf $(DEST_FOLDER)


test: build tests/tests.c tests/project_test.c
	$(CC) tests/tests.c tests/project_test.c $(SRC) $(LIBS) $(CFLAGS) $(CLIBS)