package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Ziyad678/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		log.Println("Please enter a username to login")
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Println("User doesnt exists.")
		log.Printf("failed to create user %v in database. %v\n", cmd.Args[0], err)
		os.Exit(1)
	}

	err = s.config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	log.Printf("User %v switched successfully!\n", cmd.Args[0])
	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		log.Println("Please enter a username to register")
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}
	userDB, err := s.db.CreateUser(context.Background(), user)
	if err != nil {
		fmt.Println("Failed to create user. User already exists")
		log.Printf("failed to create user %v in database. %v\n", user, err)
		os.Exit(1)
	}
	err = s.config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	log.Printf("User %v registered successfully!\n", userDB.Name)
	fmt.Println("User created successfully:")
	printUser(userDB)
	return nil
}

func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Println("Couldn't delete all user from database.")
		log.Printf("failed to delete all users from database. %v\n", err)
		os.Exit(1)
	}

	log.Printf("Users Deleted successfully!\n")
	fmt.Println("Users deleted successfully!")
	return nil
}

func handlerUsers(s *state, cmd command) error {

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Couldn't get all user from database.")
		log.Printf("failed to get all users from database. %v\n", err)
		os.Exit(1)
	}

	for _, user := range users {
		if user.Name == s.config.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}
	return nil
}


func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}