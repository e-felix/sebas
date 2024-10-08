CC=gcc
CFLAGS=-std=c99 -pedantic -Wall
BIN=sebas
DEST=bin/
SRC=sebas.c src/project.c src/database.c
LIBS=lib/database/sqlite.c
CLIBS=-lsqlite3

.PHONY: clean build run dir

all: run

dir:
	mkdir -p $(DEST)

build: dir $(SRC) $(LIBS)
	$(CC) $(SRC) $(LIBS) $(CFLAGS) $(CLIBS) -o $(DEST)$(BIN)

run: build
	./$(DEST)$(BIN)

clean:
	rm -rf $(DEST)
