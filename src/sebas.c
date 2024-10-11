#include "sebas.h"
#include "database.h"
#include "project.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void sebas_run(void) {
  Project *gb = create_project(
      "Gambettesbox", "$HOME/Projets/mylittle/sylius/gambettesbox-sylius");

  if (gb == NULL) {
    printf("Error creating gb");
    exit(EXIT_FAILURE);
  }

  printf("Project name '%s'\n", gb->name);
  printf("Project path '%s'\n", gb->absolute_path);

  delete_project(gb);

  enum DATABASE_DBMS dbms = SQLITE;
  Database *db = database_get(dbms, NULL);
  printf("dsn = %s\n", db->dsn);
  printf("DB exists: %s\n", db->instance != NULL ? "Yes" : "No");
  database_init(db);
  printf("DB initiated\n");
  database_close(db);
  printf("DB closed\n");
}