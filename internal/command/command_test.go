package command

import (
	"reflect"
	"testing"
)

func TestNewCommand(t *testing.T) {
	cmd := "echo"
	args := make([]string, 0)
	args = append(args, "Hello")
	expected := &Command{Cmd: cmd, Args: args}

	command := NewCommand(cmd, args)

	if !reflect.DeepEqual(command, expected) {
		t.Fatalf("Found %v, expected %v", command, expected)
	}
}

func TestCommandUpdateCmd(t *testing.T) {
	cmd := "echo"
	args := make([]string, 0)
	args = append(args, "Hello")

	newCmd := "ls"
	expected := &Command{Cmd: newCmd, Args: args}

	command := &Command{Cmd: cmd, Args: args}
	command.UpdateCmd(newCmd)

	if !reflect.DeepEqual(command, expected) {
		t.Fatalf("Found %v, expected %v", command, expected)
	}
}

func TestCommandUpdateArgs(t *testing.T) {
	cmd := "echo"
	args := make([]string, 0)
	args = append(args, "Hello")

	newArgs := make([]string, 0)
	newArgs = append(newArgs, "Hi")
	expected := &Command{Cmd: cmd, Args: newArgs}

	command := &Command{Cmd: cmd, Args: args}
	command.UpdateArgs(newArgs)

	if !reflect.DeepEqual(command, expected) {
		t.Fatalf("Found %v, expected %v", command, expected)
	}
}
