package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/haidar1337/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) >= 1 {
		parsed, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return errors.New("invalid limit, must be a numberr")
		}
		limit = parsed
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return fmt.Errorf("could not get posts for user: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("* %s: %s\n", post.Title, post.Description.String)
	}

	return nil
}
