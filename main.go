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
	dbURL := "postgres://postgres:postgres@localhost:5432/gator"

	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}

	db, err := sql.Open("postgres", dbURL)
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	dbQueries := database.New(db)

	programState := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(&programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
