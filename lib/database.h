#ifndef SEBAS_DATABASE_H
#define SEBAS_DATABASE_H

#define SQLITE_DB_PATH "/tmp/sebas/data/sqlite3.db"

#include <sqlite3.h>

int database_init(void);
int database_run_query(sqlite3 *db, char *query, int (*callback)(void*, int, char**, char**));

#endif
