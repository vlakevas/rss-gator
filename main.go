package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/vlakevas/rss-gator/internal/config"
	"github.com/vlakevas/rss-gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	main_cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	dbUrl := main_cfg.DBURL
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	main_state := state{ // main_state as in the state the is initialized here not implying that other states will be
		cfg: &main_cfg,
		db:  database.New(db),
	}

	main_commands := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}

	main_commands.register("login", handlerLogin)
	main_commands.register("register", handlerRegister)
	main_commands.register("reset", handlerReset)
	main_commands.register("users", handlerListUsers)

	arguments := os.Args

	if len(arguments) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	login := command{
		Name: arguments[1],
		Args: arguments[2:],
	}

	if err := main_commands.run(&main_state, login); err != nil {
		log.Fatal(err)
	}
}
