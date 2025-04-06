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

func handlerAgg(s *state, cmd command) error {
	feed, _ := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	fmt.Println(feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		log.Println("Please enter a feed name and ur to add")
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}
	username := s.config.CurrentUserName
	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    username,
	}
	_, err := s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		fmt.Println("Failed to add feed to database.")
		log.Printf("failed to add feed %v to database. %v\n", cmd.args[0], err)
		os.Exit(1)
	}

	log.Printf("Feed %v successfully added to database\n", feed.Name)
	fmt.Println("Feed successfully added to database")
	return nil
}
