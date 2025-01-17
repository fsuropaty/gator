package main

import "github.com/fsuropaty/gator-go/internal/database"

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) {

}
