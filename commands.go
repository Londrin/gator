package main

import (
	"errors"
	"fmt"

	"github.com/Londrin/rss-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Expected Command args not provided: Username required")
	}
	userName := cmd.args[0]

	err := s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("SetUser: Unable to set username: %w", err)
	}

	fmt.Printf("User Set: %s\n", userName)
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	cmdToRun := c.cmds[cmd.name]
	err := cmdToRun(s, cmd)
	if err != nil {
		return fmt.Errorf("Run Command Failed: %v - Error: %v", cmd.name, err)
	}

	return nil
}
