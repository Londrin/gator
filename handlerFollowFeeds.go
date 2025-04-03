package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Londrin/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feedFound, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("follow feed - unable to get feed: %w", err)
	}

	feed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedFound.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed Followed Successful:")
	fmt.Printf("Feed Followed: %s\nUser Following: %s\n\n", feed.FeedName, feed.UserName)

	return nil
}

func handlerFollowingFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("following feed - Unable to fetch feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("User is not following any feeds")
		return nil
	}

	fmt.Printf("Feed Follows for %s:\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.Name)
	}

	fmt.Println()

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("unfollow - unable to get feed: %w", err)
	}

	err = s.db.DeleteFeed(context.Background(), database.DeleteFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unfollow - unable to delete feed: %w", err)
	}

	fmt.Println("Feed successfully unfollowed!")
	return nil
}
