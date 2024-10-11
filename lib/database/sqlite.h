#ifndef SEBAS_SQLITE_H
#define SEBAS_SQLITE_H

#define SEBAS_SQLITE_DB_PATH "./bin/sqlite3.db"
#include <sqlite3.h>

sqlite3 *s_sqlite_get_instance(char *dsn);
void s_sqlite_close_instance(sqlite3 *db);

int s_sqlite_run_query(sqlite3 *db, char *query, int (*callback)(void*, int, char**, char**));

#endif
