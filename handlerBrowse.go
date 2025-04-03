package main

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/Londrin/rss-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2

	if len(cmd.Args) == 1 {
		if specificedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = specificedLimit
		} else {
			return fmt.Errorf("usage: %s <feed name> <post limit #>", cmd.Name)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("browse error - unable to get posts: %w", err)
	}

	regex := regexp.MustCompile("<[^>]+>")
	fmt.Printf("Found %d posts for user %s:\n\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("--- Title: %s\n", post.Title)
		fmt.Printf("--- Description: %v\n", regex.ReplaceAllString(post.Description, ""))
		fmt.Printf("--- Url: %s\n", post.Url)
		fmt.Printf("--- Published: %v\n\n", post.PublishedAt.Format("Mon Jan 2, 2006 15:04:05"))
	}

	return nil
}
