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

func handlerFollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 1 {
		log.Println("Please enter a URL to follow")
		return fmt.Errorf("usage: %s <URL>", cmd.Name)
	}
	feed, err := s.db.GetFeedByURL(context.Background(),cmd.Args[0])
	if err!=nil{
		log.Printf("Can't get feed %v from database\n",s.config.CurrentUserName)
		return err
	}
	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:      user.ID,
		FeedID: feed.ID,
	}
	feedFromDB, err := s.db.CreateFeedFollow(context.Background(), feedFollow) 
	if err != nil {
		fmt.Println("Failed to follow feed.")
		log.Printf("failed to follow feed %v in database. %v\n", cmd.Args[0], err)
		os.Exit(1)
	}
	fmt.Println("Feed follow created:")
	printFeedFollow(feedFromDB.UserName, feedFromDB.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(),user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get followed feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("You dont follow any feeds.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, ff := range feeds {
		fmt.Printf("* %s\n", ff.FeedName)
	}

	return nil
}


func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}


func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 1 {
		log.Println("Please enter a URL to unfollow")
		return fmt.Errorf("usage: %s <URL>", cmd.Name)
	}
	feed, err := s.db.GetFeedByURL(context.Background(),cmd.Args[0])
	if err!=nil{
		log.Printf("Can't get feed %v from database\n",s.config.CurrentUserName)
		return err
	}
	toDelete := database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}
	err = s.db.DeleteFeedFollow(context.Background(),toDelete)
	if err!=nil{
		log.Printf("Can't delete feed follow from database\n")
		return err
	}
	return nil
}