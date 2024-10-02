CC=gcc
CFLAGS=-std=c99 -pedantic -Wall
BIN=sebas
DEST=bin/
SRC=sebas.c src/project.c lib/database.c
LIBS=-lsqlite3

.PHONY: clean build run dir

all: run

dir:
	mkdir -p $(DEST)

build: dir $(SRC)
	$(CC) $(SRC) $(CFLAGS) $(LIBS) -o $(DEST)$(BIN)

run: build
	./$(DEST)$(BIN)

clean:
	rm sebas sqlite3.db
