package main

import (
	"context"
	"fmt"
)

func commandSelect(s *state, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("Usage: select <user/bill> <name>")
	}

	var err error

	switch args[0] {
	case "user":
		err = selectUser(s, args[1])
	case "bill":
		err = selectBill(s, args[1])
	default:
		err = fmt.Errorf("Invalid argument, usage: <user/bill> <name>")
	}

	if err != nil {
		return err
	}

	return nil
}

func selectUser(s *state, name string) error {
	user, err := s.db.GetUserByName(context.Background(), name)
	if err != nil {
		return fmt.Errorf("Failed to get user '%v': %v", name, err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Failed to set user '%v': %v", user.Name, err)
	}

	printUser(user)

	return nil
}

func selectBill(s *state, name string) error {
	bill, err := s.db.GetTransactionByName(context.Background(), name)
	if err != nil {
		return fmt.Errorf("Failed to get transaction '%v': %v", name, err)
	}

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Failed to get user '%v': %v", name, err)
	}

	if user.ID != bill.UserID {
		return fmt.Errorf("Bill '%v' is not registered for current user", name)
	}

	printTransaction(bill)

	return nil
}
