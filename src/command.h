#ifndef SEBAS_COMMAND_H
#define SEBAS_COMMAND_H

typedef struct Command {
  char *name;
  char *cmd;
  char *args;
} Command;

Command *create_command(char *name, char *cmd, char *args);

void modify_command_name(Command *command, char *new_name);

void modify_command_cmd(Command *command, char *new_cmd);

void modify_command_args(Command *command, char *new_args);

void delete_command(Command *command);

#endif
