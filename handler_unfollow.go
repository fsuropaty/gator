package main

import (
	"context"
	"fmt"

	"github.com/fsuropaty/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedsByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Couldn't get the feed: %w", err)
	}

	params := database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Couldn't unfollow the feed: %w", err)
	}

	fmt.Printf("Unfollow succedeed: %s\n", url)

	return nil
}
