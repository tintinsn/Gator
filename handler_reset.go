package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) > 1 {
		return errors.New("not required username")
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("fail to delete users: %w", err)
	}

	fmt.Println("Users has been delete")
	return nil
}
