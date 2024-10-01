package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/haidar1337/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("username is required. gator register <username>")
	}

	u, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		os.Exit(1)
	}

	err = s.config.SetUser(u.Name)
	if err != nil {
		return err
	}

	fmt.Println("user created successfully")
	printUserDetails(u)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("username is required. gator login <username>")
	}

	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("could not find user %s", username)
	}

	err = s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user in config %w", err)
	}

	fmt.Printf("switched to user %s successfully\n", username)
	return nil
}

func printUserDetails(u database.User) {
	fmt.Printf("** ID: %v\n", u.ID)
	fmt.Printf("** Name: %v\n", u.Name)
}
