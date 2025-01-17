package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fsuropaty/gator-go/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}

	currentUser := s.cfg.CurrentUserName
	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("Current user not found: %w", err)
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("Couldn't create feed: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Couldn't create feed follow record: %w", err)
	}

	fmt.Printf("Feed created: %v (%v)\n", feed.Name, feed.Url)
	fmt.Printf("Feed has been added to follow record: %v", feedFollow.FeedName)

	return nil
}
