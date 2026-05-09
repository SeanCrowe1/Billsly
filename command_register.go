package main

import (
	"context"
	"fmt"
	"time"

	"Billsly/internal/database"

	"github.com/google/uuid"
)

func commandRegister(s *state, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: register <username>")
	}

	name := args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("Couldn't create user: %v", name)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
