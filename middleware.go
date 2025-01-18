package main

import (
	"context"
	"fmt"

	"github.com/fsuropaty/gator-go/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, c command) error {

		currentUser := s.cfg.CurrentUserName

		user, err := s.db.GetUser(context.Background(), currentUser)
		if err != nil {
			return fmt.Errorf("Current user not found: %w", err)
		}

		handler(s, c, user)
		return nil
	}
}
