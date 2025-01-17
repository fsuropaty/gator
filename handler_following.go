package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {

	currentUserName := s.cfg.CurrentUserName

	currentUser, err := s.db.GetUser(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("Couldn't get user id: %w", err)
	}

	followedFeed, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("Couldn't get feed follows for user: %w", err)
	}

	for _, feed := range followedFeed {
		fmt.Println(feed.FeedName)
	}

	return nil
}
