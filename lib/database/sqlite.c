#include "sqlite.h"
#include <stdio.h>

sqlite3 *db;

sqlite3* s_sqlite_get_instance(char *dsn) {
    int rc;
    rc = sqlite3_open(dsn, &db);
    if (rc) {
      fprintf(stderr, "Can't open database \"%s\"\n", sqlite3_errmsg(db));
      sqlite3_close(db);
      return NULL;
    }

    return db;
}

void s_sqlite_close_instance(sqlite3 *db) {
    sqlite3_close(db);
}

int s_sqlite_run_query(sqlite3 *db, char *query, int (*callback)(void*, int, char**, char**)) {
    int rc;
    char *errorMsg;
    rc = sqlite3_exec(db, query, callback, 0, &errorMsg);
    if (rc != SQLITE_OK) {
      fprintf(stderr, "%s\n", sqlite3_errmsg(db));
      sqlite3_close(db);
      sqlite3_free(errorMsg);
      return 1;
    }

    return 0;
}
