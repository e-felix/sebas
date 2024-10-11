#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include "../src/project.h"

typedef struct TestResult {
    int result_code;
    char* result_message;
} TestResult;

TestResult* test_create_project(void);
// void modify_project_name(Project *project, char *new_name);
// void modify_project_absolute_path(Project *project, char *new_absolute_path);
// void delete_project(Project *project);

void print_result(char* function_name, TestResult *result);

void run_project_tests(void) {
    printf("Starting to run tests for domain: Project\n");
    print_result("test_create_project", test_create_project());
}

void print_result(char* function_name, TestResult *result) {
    printf("Testing \"%s\":...", function_name);

    if (result->result_code == 0) {
        printf("SUCCESSED\n");
    } else {
        printf("FAILED\n");
        printf("\n\n\n");
        printf("%s", result->result_message);
    }

    free(result);
}

TestResult* test_create_project(void)
{
    TestResult *result = malloc(sizeof(TestResult));
    result->result_code = 0;

    Project *project;
    char* expected_name = "my_project";
    char* expected_path = "path/to/my/project";

    project = create_project(expected_name, expected_path);

    bool eq = project->name == expected_name;
    if (!eq) {
        result->result_code = 1;
        char *message = NULL;
        sprintf(message, "Expected \"%s\". Got \"%s\"", expected_name, project->name);
        result->result_message = message;
    }

    return result;
}
