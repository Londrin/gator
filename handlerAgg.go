package main

import (
	"fmt"
	"time"

	"github.com/Londrin/rss-aggregator/internal/database"
)

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <duration as string 'e.g 1s, 1m, 1h'>", cmd.Name)
	}

	interval := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(interval)
	if err != nil {
		return fmt.Errorf("agg error - input command unable to be processed: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s, user)
	}
}
