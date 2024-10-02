#ifndef SEBAS_PROJECT_H
#define SEBAS_PROJECT_H

typedef struct Project {
    int id;
    char *name;
    char *absolute_path;
} Project;

Project *create_project(char *name, char *absolute_path);

void modify_project_name(Project *project, char *new_name);

void modify_project_absolute_path(Project *project, char *new_absolute_path);

void delete_project(Project *project);

#endif
