package main

import (
	"errors"
	"log"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	log.Println("Entering Register Command function")
	_, ok := c.registeredCommands[name]
	if ok {
		log.Printf("Fialed to register command. %v already exists", name)
		return
	}
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	log.Printf("Entering run function. running %v command", cmd.Name)
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s,cmd)
}
