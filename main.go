package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/fsuropaty/gator-go/internal/config"
	"github.com/fsuropaty/gator-go/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	var s state

	dbURL := "postgres://postgres:postgres@localhost:5432/gator"

	db, err := sql.Open("postgres", dbURL)
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	dbQueries := database.New(db)
	s.db = dbQueries

	conf, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}

	s.cfg = &conf

	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments provided")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
