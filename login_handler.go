package main

import (
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		log.Println("Please enter a username to login")
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	log.Printf("User %v switched successfully!\n", cmd.args[0])
	fmt.Println("User switched successfully!")
	return nil
}
