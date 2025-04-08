package main

import (
	"context"
	"log"

	"github.com/Ziyad678/Gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error{
	
	return func(s *state, cmd command)error{
		user, err := s.db.GetUser(context.Background(),s.config.CurrentUserName)
		if err!=nil{
			log.Printf("Can't get user %v from database\n",s.config.CurrentUserName)
			return err
		}
		return handler(s,cmd,user)
	}
}