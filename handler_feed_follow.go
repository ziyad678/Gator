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

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		log.Println("Please enter a URL to follow")
		return fmt.Errorf("usage: %s <URL>", cmd.Name)
	}
	user := database.CreateFeedF{
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
