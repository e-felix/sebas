#ifndef SEBAS_DATABASE_H
#define SEBAS_DATABASE_H

#define DBMS_SQLITE 0

enum DATABASE_DBMS {
    SQLITE = DBMS_SQLITE,
};

typedef struct Database {
    enum DATABASE_DBMS dbms;
    char *dsn;
    void *instance;
} Database;

Database *database_get(enum DATABASE_DBMS type, char *dsn);

void database_close(Database *db);

int database_init(Database *db);

#endif
