package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 1 {
		return errors.New("not required argument")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	for _, feed := range feeds {
		username, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Error: %w, err", err)
		}

		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(username)

	}
	return nil
}
