package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Londrin/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("Expected Command args not provided: Username required")
	}

	user := cmd.Args[0]
	err := s.cfg.SetUser(user)
	if err != nil {
		return fmt.Errorf("Login - SetUser: Unable to set username: %w", err)
	}

	fmt.Printf("User Set: %s\n", user)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("list Users: Unable to return users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)

	}
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("Register User: Provide username to register")
	}

	userName := cmd.Args[0]

	usr, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	})
	if err != nil {
		return fmt.Errorf("Register User: Unable to register user: %w", err)
	}

	err = s.cfg.SetUser(usr.Name)
	if err != nil {
		return fmt.Errorf("Unable to set user in config: %w", err)
	}

	fmt.Printf("ID: %v\nCreated: %v\nUpdated: %v\nName: %v\n", usr.ID, usr.CreatedAt, usr.UpdatedAt, usr.Name)

	return nil
}
