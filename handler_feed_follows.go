package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/tintinsn/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return errors.New("url is required")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("Feed not exists")
		}
		return err
	}
	
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("Error creating feed follow: %w", err)
	}

	fmt.Println("Feed Name: ", feedFollow.FeedName)
	fmt.Println("User: ", feedFollow.UserName)

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("usage: unfollow <url>")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Not found feed : %w", err)
	}

	rows, err := s.db.DeleteFeedFollowByUser(context.Background(), database.DeleteFeedFollowByUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("you weren't following this feed")
	}

	return nil
}
