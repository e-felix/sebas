#include "src/project.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "lib/database.h"

int main(void) {
  Project *gb = create_project(
      "Gambettesbox", "$HOME/Projets/mylittle/sylius/gambettesbox-sylius");

  if (gb == NULL) {
    printf("Error creating gb");
    return EXIT_FAILURE;
  }

  printf("Project name '%s'\n", gb->name);
  printf("Project path '%s'\n", gb->absolute_path);

  delete_project(gb);

  return EXIT_SUCCESS;
}

static int query_result_callback(void *NotUsed, int argc, char **argv, char **azColName) {
  for (int i = 0; i < argc; i++) {
    printf("%s = %s\n", azColName[i], argv[i] ? argv[i] : NULL);
  }
  printf("\n");

  return 0;
}
