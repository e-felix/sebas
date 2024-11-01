package command

import (
	"reflect"
	"testing"
)

func TestCreateCommandWithEmptyArgs(t *testing.T) {
	cmd := "ls"
	args := make([]string, 0)
	expected := &Command{Cmd: cmd, Args: args}

	command := CreateCommand(cmd, args)

	if !reflect.DeepEqual(command, expected) {
		t.Fatalf("Found %v, expected %v", command, expected)
	}
}
