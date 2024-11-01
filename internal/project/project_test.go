package project

import (
	"reflect"
	"testing"

	. "github.com/e-felix/sebas/internal/command"
	. "github.com/e-felix/sebas/internal/env"
)

var t *testing.T

func TestNewProject(t *testing.T) {
	expected := &Project{Id: 1, Name: "MyProject", Envs: make([]Env, 0), Cmds: make([]Command, 0)}

	project := NewProject("MyProject")

	if !reflect.DeepEqual(project, expected) {
		t.Fatalf("Found %v, expected %v", project, expected)
	}
}

func TestUpdateProject(t *testing.T) {
	newName := "MyNewName"
	expected := &Project{Id: 1, Name: newName, Envs: make([]Env, 0), Cmds: make([]Command, 0)}

	project := NewProject("MyProject")
	project.Update(newName)

	if !reflect.DeepEqual(project, expected) {
		t.Fatalf("Found %v, expected %v", project, expected)
	}
}

func TestAddEnv(t *testing.T) {
	env := Env{Key: "FOO", Value: "BAR"}
	expected := &Project{Id: 1, Name: "MyProject", Envs: []Env{env}, Cmds: make([]Command, 0)}

	project := NewProject("MyProject")
	project.AddEnv(env)

	if !reflect.DeepEqual(project, expected) {
		t.Fatalf("Found %v, expected %v", project, expected)
	}
}

func TestAddCommand(t *testing.T) {
	cmd := Command{Cmd: "echo", Args: []string{"Hello", "World"}}
	expected := &Project{Id: 1, Name: "MyProject", Envs: make([]Env, 0), Cmds: []Command{cmd}}

	project := NewProject("MyProject")
	project.AddCmd(cmd)

	if !reflect.DeepEqual(project, expected) {
		t.Fatalf("Found %v, expected %v", project, expected)
	}
}

func TestRemoveCommand(t *testing.T) {
	cmd := Command{Cmd: "echo", Args: []string{"Hello", "World"}}
	expected := &Project{Id: 1, Name: "MyProject", Envs: make([]Env, 0), Cmds: []Command{cmd}}

	project := NewProject("MyProject")
	project.AddCmd(cmd)
	project.RemoveCmd(cmd)

	if !reflect.DeepEqual(project, expected) {
		t.Fatalf("Found %v, expected %v", project, expected)
	}
}
