package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/fsuropaty/gator-go/internal/database"
	"github.com/fsuropaty/gator-go/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %v <1s,1m,1h, etc>", cmd.Name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])

	if err != nil {
		return fmt.Errorf("Invalid duration format: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)

	for ; ; <-ticker.C {
		if err := scrapeFeed(s); err != nil {
			fmt.Printf("Failed to scrape the feed: %v\n", err)
		}
	}

	return nil

}

func scrapeFeed(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())

	if err == sql.ErrNoRows {
		fmt.Println("No feeds to fetch, waiting ...")
		return nil
	} else if err != nil {
		return fmt.Errorf("Couldn't get the next feed: %w", err)
	}

	params := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now()},
		UpdatedAt:     time.Now(),
		ID:            nextFeed.ID,
	}

	_, err = s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Couldn't mark the feed: %w", err)
	}
	rssFeed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("Couldn't fetch feed: %w", err)
	}

	fmt.Printf("Feed Title: %s\n", rssFeed.Channel.Title)
	fmt.Printf("Feed Link: %s\n", rssFeed.Channel.Link)
	fmt.Printf("Feed Description: %s\n", rssFeed.Channel.Description)

	for _, item := range rssFeed.Channel.Item {
		fmt.Println("-----------------------------------------------")
		fmt.Printf("Article Title: %s\n", item.Title)
		fmt.Printf("Article Link: %s\n", item.Link)
		fmt.Printf("Article Description: %s\n", item.Description)
		fmt.Printf("Published Date: %s\n\n", item.PubDate)
	}

	return nil
}
