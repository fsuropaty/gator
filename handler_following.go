package main

import (
	"context"
	"fmt"

	"github.com/fsuropaty/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	followedFeed, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Couldn't get feed follows for user: %w", err)
	}

	for _, feed := range followedFeed {
		fmt.Println(feed.FeedName)
	}

	return nil
}
