package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/fsuropaty/gator-go/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	_, err := s.db.GetUser(context.Background(), params.Name)

	if err == nil {
		fmt.Println("User already exists!")
		os.Exit(1)
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("Error checking for user: %w", err)
	}

	user, err := s.db.CreateUser(context.Background(), params)

	if err != nil {
		return fmt.Errorf("Failed to create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)

	fmt.Println("User has been registered")
	printUser(user)

	return nil

}

func handlerLogin(s *state, cmd command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("Couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("Couldn't set user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to get users list: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Println("* ", user.Name)
		}
	}

	return nil

}

func printUser(user database.User) {
	fmt.Printf(" * ID:		%v\n", user.ID)
	fmt.Printf(" * Name:	%v\n", user.Name)
}
