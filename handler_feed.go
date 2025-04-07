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

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		log.Println("Please enter a feed name and ur to add")
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	user , err:=s.db.GetUser(context.Background(),s.config.CurrentUserName)
	if err != nil {
		fmt.Println("Couldn't get user from database.")
		log.Printf("failed to get user %v from database. %v\n",s.config.CurrentUserName, err)
		os.Exit(1)
	}
	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}
	_, err = s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		fmt.Println("Failed to add feed to database.")
		log.Printf("failed to add feed %v to database. %v\n", cmd.Args[0], err)
		os.Exit(1)
	}

	log.Printf("Feed %v successfully added to database\n", feed.Name)
	fmt.Println("Feed successfully added to database")
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		printFeed(feed, user)
		fmt.Println("=====================================")
	}

	return nil
}
func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
}