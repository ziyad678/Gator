package main

import (
	"log"
)

type command struct {
	name string
	args []string
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
	log.Printf("Entering run function. running %v command", cmd.name)
	err := handlerLogin(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
