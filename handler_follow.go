package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fsuropaty/gator-go/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feeds, err := s.db.GetFeedsByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Couldn't get feeds by URL: %w", err)
	}

	currentUser := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("Couldn't get current user: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feeds.ID,
		UserID:    user.ID,
	}

	followRecord, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Couldn't create a new feed follow record: %w", err)
	}

	fmt.Println(followRecord.FeedName)
	fmt.Println(followRecord.UserName)

	return nil
}
