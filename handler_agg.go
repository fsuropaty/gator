package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/fsuropaty/gator-go/internal/database"
	"github.com/fsuropaty/gator-go/rss"
	"github.com/google/uuid"
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

	for range ticker.C {
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

	for _, post := range rssFeed.Channel.Item {
		pDate, err := parseDate(post.PubDate)
		if err != nil {
			return err
		}

		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			Title:       post.Title,
			Url:         post.Link,
			Description: post.Description,
			PublishedAt: pDate,
			FeedID:      nextFeed.ID,
		}

		createdPost, err := s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			return fmt.Errorf("Failed to create post: %w", err)
		}

		if createdPost.ID == uuid.Nil {
			fmt.Printf("Skipped existing post: %s\n", post.Title)
		} else {
			fmt.Printf("Created new post: %s\n", post.Title)
		}

	}

	fmt.Printf("Feed %s fetched successfully\n", nextFeed.Url)

	return nil
}

func parseDate(dateStr string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z,          // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC1123,           // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC822,            // "02 Jan 06 15:04 MST"
		"2006-01-02T15:04:05Z", // ISO 8601
		"Mon, 2 Jan 2006 15:04:05 -0700",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, dateStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("Couldn't parse date: %s", dateStr)
}
