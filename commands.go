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

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f

}

func (c *commands) run(s *state, cmd command) error {
	if handler, exists := c.handlers[cmd.name]; exists {
		return handler(s, cmd)
	} else {
		return fmt.Errorf("The command is not exists")
	}
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No arguments found")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])

	if err == sql.ErrNoRows {
		fmt.Println("User doesn't exists!")
		os.Exit(1)
	} else if err != nil {
		return fmt.Errorf("Error checking for user: %w", err)
	}

	err = s.cfg.SetUser(cmd.args[0])

	if err != nil {
		return fmt.Errorf("Error set user: %w", err)
	}

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No arguments found")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	_, err := s.db.GetUser(context.Background(), params.Name)

	if err == nil {
		fmt.Println("User already exists!")
		os.Exit(1)
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("Error checking for user: %w", err)
	}

	_, err = s.db.CreateUser(context.Background(), params)

	if err != nil {
		return fmt.Errorf("Failed to create user: %w", err)
	}

	err = s.cfg.SetUser(params.Name)

	fmt.Println("User has been registered: ", params.Name)

	return nil

}

func handlerReset(s *state, cmd command) error {

	err := s.db.DelUsersTable(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to reset users table: %w", err)
	}

	fmt.Println("Table users has been resetted")

	return nil

}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.Users(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to get users list: %w", err)
	}

	for _, user := range users {
		fmt.Println("* ", user)
	}

	return nil

}
