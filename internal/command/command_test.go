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