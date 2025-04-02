package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return errors.New("reset db: Too many arguments provided")
	}

	err := s.db.RemoveAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("remove db: %w", err)
	}

	fmt.Println("Reset DB: Success")
	return nil
}
