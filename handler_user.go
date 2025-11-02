package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tintinsn/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("username is required")
	}

	username := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user '%s' not found", username)
		}
		return fmt.Errorf("error getting user: %w", err)
	}

	if err := s.cfg.SetUser(username); err != nil {
		return errors.New("cannot set username")
	}

	fmt.Println("User has been set to:", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("username is required")
	}

	username := cmd.args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			fmt.Println("Error: User with that name already exists")
			os.Exit(1)
		}
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	fmt.Printf("ID:         %v\n", user.ID)
	fmt.Printf("Name:       %v\n", user.Name)
	fmt.Printf("Created At: %v\n", user.CreatedAt)

	return nil

}

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.args) > 1 {
		return errors.New("not required username")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("fail to getting users: %w", err)
	}

	for _, user := range users {
		if s.cfg.CurrentUserName == user.Name {
			fmt.Println(user.Name, "(current)")
		} else {
			fmt.Println(user.Name)
		}

	}

	return nil

}
