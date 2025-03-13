package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Expected Command args not provided: Username required")
	}
	userName := cmd.Args[0]

	usr, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("Login - Get User: Unable to get user: %w", err)
	}

	err = s.cfg.SetUser(usr.Name)
	if err != nil {
		return fmt.Errorf("Login - SetUser: Unable to set username: %w", err)
	}

	fmt.Printf("User Set: %s\n", userName)
	return nil
}
