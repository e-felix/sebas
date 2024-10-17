#include "database.h"
#include "../lib/database/sqlite.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct _database {
    enum DATABASE_DBMS dbms;
    char *dsn;
    void *instance;
} Database;

static int query_result_callback(void *NotUsed, int argc, char **argv, char **azColName) {
  printf("Callback called\n");
  for (int i = 0; i < argc; i++) {
    printf("%s = %s\n", azColName[i], argv[i] ? argv[i] : NULL);
  }
  printf("\n");

  return 0;
}

Database *database_get(enum DATABASE_DBMS dbms, char* dsn) {
    Database *database = malloc(sizeof(Database));

    database->dbms = dbms;
    database->dsn = SEBAS_SQLITE_DB_PATH;

    if (dsn != NULL && strlen(dsn) > 0) {
        database->dsn = dsn;
    }

    switch (database->dbms) {
        case DBMS_SQLITE: {
            database->instance = s_sqlite_get_instance(database->dsn);
            break;
        }
        default: {
            printf("Database SGB is not supported.\n");
            exit(EXIT_FAILURE);
        }
    }

    if (database->instance == NULL) {
      printf("Database could not be created\n");
      exit(EXIT_FAILURE);
    }

    return database;
}

void database_close(Database* db) {
    free(db);
}

int database_init(Database *db) {
  sqlite3 *sqlite_db = (sqlite3 *) db->instance;

  s_sqlite_run_query(sqlite_db,
                "CREATE TABLE IF NOT EXISTS col_test (id INT PRIMARY KEY, name "
                "VARCHAR(255))", query_result_callback);

    int i = 0;
    char *sql = "INSERT INTO col_test VALUES (%d, 'test %d')";
    char *query = malloc(strlen(sql) + sizeof(i) * 2);

    while (i <= 10) {
      sprintf(query, sql, i, i);
      s_sqlite_run_query(sqlite_db, query, query_result_callback);
      i++;
    }
    free(query);

  s_sqlite_run_query(sqlite_db, "SELECT * FROM col_test", query_result_callback);

  s_sqlite_close_instance(sqlite_db);
  return 0;
}
