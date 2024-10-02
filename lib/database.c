#include "database.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static int query_result_callback(void *NotUsed, int argc, char **argv, char **azColName) {
  for (int i = 0; i < argc; i++) {
    printf("%s = %s\n", azColName[i], argv[i] ? argv[i] : NULL);
  }
  printf("\n");

  return 0;
}

int database_init(void) {
  sqlite3 *db;
  int rc;

  rc = sqlite3_open(SQLITE_DB_PATH, &db);
  if (rc) {
    fprintf(stderr, "Can't open database \"%s\"\n", sqlite3_errmsg(db));
    sqlite3_close(db);
    return 1;
  }

  rc = database_run_query(db,
                "CREATE TABLE IF NOT EXISTS col_test (id INT PRIMARY KEY, name "
                "VARCHAR(255))", query_result_callback);

  {
    int i = 1;
    char *sql = "INSERT INTO col_test VALUES (%d, 'test %d')";
    char *query = malloc(strlen(sql) + sizeof(i) * 2);

    while (i <= 10) {
      sprintf(query, sql, i, i);
      rc = database_run_query(db, query, query_result_callback);
      i++;
    }
    free(query);
  }

  rc = database_run_query(db, "SELECT * FROM col_test", query_result_callback);

  sqlite3_close(db);
  return 0;
}

int database_run_query(sqlite3 *db, char *query, int (*callback)(void*, int, char**, char**)) {
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
