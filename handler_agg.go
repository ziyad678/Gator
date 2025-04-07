package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feed, _ := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	fmt.Println(feed)
	return nil
}

