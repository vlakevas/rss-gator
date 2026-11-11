package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/vlakevas/rss-gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("'%s' expects 1 argument, none were given", cmd.Name)
	}

	if len(cmd.Args) > 1 {
		return fmt.Errorf("'%s' expects 1 argument, too many were given", cmd.Name)
	}

	usr, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(usr.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Printf("info:\n'%s'\n has been set!\n", usr)

	return nil

}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("'%s' expects 1 argument, none were given", cmd.Name)
	}

	if len(cmd.Args) > 1 {
		return fmt.Errorf("'%s' expects 1 argument, too many were given", cmd.Name)
	}

	userData := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	usr, err := s.db.CreateUser(context.Background(), userData)
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(usr.Name); err != nil {
		return err
	}
	fmt.Println(usr)

	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("'%s' doesn't take arguments", cmd.Name)
	}

	if err := s.db.ResetUsers(context.Background()); err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("the user database has been cleaned!")
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("'%s' doesn't take arguments", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return fmt.Errorf("user database is currently empty")
	}

	for _, user := range users {
		userName := fmt.Sprintf("* %s", user.Name)

		if s.cfg.CurrentUserName == user.Name {
			userName += " (current)"
		}

		fmt.Println(userName)
	}
	return nil
}
