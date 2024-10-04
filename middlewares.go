package main

import (
	"context"
	"fmt"

	"github.com/haidar1337/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.CurrentUser)
		if err != nil {
			return fmt.Errorf("could not get user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
