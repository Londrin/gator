package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Londrin/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
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
