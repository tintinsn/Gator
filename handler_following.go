package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/tintinsn/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	if len(cmd.args) > 1 {
		return errors.New("argument not required")
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error getting following: %w", err)
	}

	for _, feed := range following {
		fmt.Println(feed.FeedName)
	}

	return nil
}
