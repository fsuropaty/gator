package main

import (
	"context"
	"fmt"

	"github.com/fsuropaty/gator-go/rss"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsList(context.Background())

	return nil
}
