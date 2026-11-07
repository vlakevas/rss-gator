package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("'%s' expects 1 argument, none were given", cmd.Name)
	}

	if len(cmd.Args) > 1 {
		return fmt.Errorf("'%s' expects 1 argument, too many were given", cmd.Name)
	}

	if err := s.cfg.SetUser(cmd.Args[0]); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Printf("info: user '%s' has been set!\n", s.cfg.CurrentUserName)

	return nil

}
