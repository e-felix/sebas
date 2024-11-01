package project

import (
	"reflect"
	"testing"

	. "github.com/e-felix/sebas/internal/command"
	. "github.com/e-felix/sebas/internal/env"
)

func TestNewProject(t *testing.T) {
	expected := &Project{Id: 1, Name: "MyProject", Envs: make([]Env, 0), Cmds: make([]Command, 0)}

	newProject := NewProject("MyProject")

	assertEqual(expected, newProject, t)
}

func TestUpdateProject(t *testing.T) {
	newName := "MyNewName"
	expected := &Project{Id: 1, Name: newName, Envs: make([]Env, 0)}

	newProject := NewProject("MyProject")
	newProject.Update(newName)

	assertEqual(expected, newProject, t)
}

func TestAddEnv(t *testing.T) {
	env := Env{Key: "FOO", Value: "BAR"}
	envs := []Env{env}
	expected := &Project{Id: 1, Name: "MyProject", Envs: envs}

	newProject := NewProject("MyProject")
	newProject.AddEnv(env)

	assertEqual(expected, newProject, t)
}

func assertEqual(expected *Project, actual *Project, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Found %v, expected %v", actual, expected)
	}
}
