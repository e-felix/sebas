package project

import (
	"fmt"

	. "github.com/e-felix/sebas/internal/command"
	. "github.com/e-felix/sebas/internal/env"
)

type Project struct {
	Id   int
	Name string
	Envs []Env
	Cmds []Command
}

func NewProject(name string) *Project {
	return &Project{
		Id:   1,
		Name: name,
		Envs: make([]Env, 0),
		Cmds: make([]Command, 0),
	}
}

func (p *Project) Print() {
	fmt.Printf("Id: %d\n", p.Id)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Println("Envs:")

	for _, env := range p.Envs {
		fmt.Printf("\t%v: %v\n", env.Key, env.Value)
	}
}

func (p *Project) Update(name string) {
	p.Name = name
}

func (p *Project) AddEnv(newEnv Env) bool {
	envs := p.Envs

	envExists := false
	for _, env := range envs {
		if env.Key == newEnv.Key {
			envExists = true
			break
		}

	}

	if !envExists {
		p.Envs = append(p.Envs, newEnv)
		return true
	}

	return false
}
