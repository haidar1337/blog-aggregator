package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/haidar1337/gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("name and url of feed are required, gator addfeed <name> <url>")
	}

	currentUser := s.config.CurrentUser
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("current user %s could not be retrieved from database: %w", currentUser, err)
	}

	feed_name := cmd.args[0]
	feed_url := cmd.args[1]
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		FeedName:  feed_name,
		FeedUrl:   feed_url,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("could not add feed: %w", err)
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to follow feed: %w", err)
	}

	fmt.Println("feed created successfully")
	fmt.Println(feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsWithUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds from database: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found")
	} else {
		printFeeds(feeds)
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	req.Header.Add("user-agent", "gator")
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out := RSSFeed{}
	err = xml.Unmarshal(respBody, &out)
	if err != nil {
		return nil, err
	}
	out.Channel.Title = html.UnescapeString(out.Channel.Title)
	out.Channel.Description = html.UnescapeString(out.Channel.Description)
	for i := 1; i < len(out.Channel.Item); i++ {
		out.Channel.Item[i].Title = html.UnescapeString(out.Channel.Item[i].Title)
		out.Channel.Item[i].Description = html.UnescapeString(out.Channel.Item[i].Description)
	}

	return &out, nil
}

func printFeeds(feeds []database.GetFeedsWithUsersRow) {
	for _, feed := range feeds {
		fmt.Printf("* %s: %s, Created by: %s\n", feed.FeedName, feed.FeedUrl, feed.Name)
	}
}
