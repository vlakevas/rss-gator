package main

import (
	"log"
	"os"

	"github.com/vlakevas/rss-gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	main_cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	main_state := state{ // main state as in the state the is initialized here not implying that other states will be
		cfg: &main_cfg,
	}

	main_commands := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}

	main_commands.register("login", handlerLogin)

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
