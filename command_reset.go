package main

import (
	"context"
	"fmt"
)

func commandReset(s *state, args ...string) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
