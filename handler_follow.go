package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/haidar1337/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("a url is required. gator follow <url>")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("feed does not exist: %w", err)
	}

	currentUser := s.config.CurrentUser
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("current user was not registered in database: %w", err)
	}

	createdFeedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to follow feed: %w", err)
	}

	fmt.Println("successfully followed feed")
	printFeedFollow(createdFeedFollow)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	currentUser := s.config.CurrentUser
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("current user was not registered in database: %w", err)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get user followed feeds: %w", err)
	}
	if len(feeds) == 0 {
		fmt.Println("not following any feed")
	} else {
		printUserFeeds(feeds)
	}

	return nil
}

func printUserFeeds(feeds []database.GetFeedFollowsForUserRow) {
	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf("Feed: %s\nUser: %s\n", feedFollow.UserName, feedFollow.FeedName)
}
