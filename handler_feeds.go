package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsList(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't get feeds list: %w", err)
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(feed.UserName)
		fmt.Println()
	}

	return nil
}
