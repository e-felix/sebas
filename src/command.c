#include "command.h"
#include <stdlib.h>

Command *create_command(char *name, char *cmd, char *args) {
  Command *command = malloc(sizeof(Command));
  command->name = name;
  command->cmd = cmd;
  command->args = args;

  return command;
}

void modify_command_name(Command *command, char *new_name) {
  command->name = new_name;
}

void modify_command_cmd(Command *command, char *new_cmd) {
  command->cmd = new_cmd;
}

void modify_command_args(Command *command, char *new_args) {
  command->args = new_args;
}

void delete_command(Command *command) { free(command); }
