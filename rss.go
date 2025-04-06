package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"log"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	log.Println("Entering fetchFeed function")

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		log.Println("error creating request.", err)
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("website respond with error.", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body.", err)
	}

	rssFeed := RSSFeed{}
	xml.Unmarshal(body, &rssFeed)
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for _, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return &rssFeed, nil
}
