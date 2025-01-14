package main

import (
	"context"
	"fmt"

	"github.com/fsuropaty/gator-go/rss"
)

func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()
	feedURL := "https://www.wagslane.dev/index.xml"

	rssFeed, err := rss.FetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("Error fetching the RSS feed: %w", err)
	}

	fmt.Printf("Feed Title: %s\n", rssFeed.Channel.Title)
	fmt.Printf("Feed Link: %s\n", rssFeed.Channel.Link)
	fmt.Printf("Feed Description: %s\n", rssFeed.Channel.Description)

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Article Title: %s\n", item.Title)
		fmt.Printf("Article Link: %s\n", item.Link)
		fmt.Printf("Article Description: %s\n", item.Description)
		fmt.Printf("Published Date: %s\n\n", item.PubDate)
	}

	return nil

}
