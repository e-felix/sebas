#include "project.h"
#include <stdio.h>
#include <stdlib.h>

Project *create_project(char *name, char *absolute_path) {
  Project *project = malloc(sizeof(Project));

  if (project == NULL) {
    printf("Error: could not allocate memory for new project \"%s\"", name);
    exit(EXIT_FAILURE);
  }

  project->name = name;
  project->absolute_path = absolute_path;

  return project;
}

void modify_project_name(Project *project, char *new_name) {
  project->name = new_name;
}

void modify_project_absolute_path(Project *project, char *new_absolute_path) {
  project->absolute_path = new_absolute_path;
}

void delete_project(Project *project) { project = NULL; }
