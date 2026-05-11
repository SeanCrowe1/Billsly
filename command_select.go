package main

import (
	"context"
	"fmt"
)

func commandSelect(s *state, args ...string) error {
	var err error

	switch args[0] {
	case "user":
		if len(args) != 2 {
			return fmt.Errorf("Usage: select user <name>")
		}
		err = selectUser(s, args[1])
	case "bill":
		if len(args) != 2 {
			return fmt.Errorf("Usage: select bill <name>")
		}
		err = selectBill(s, args[1])
	case "bills":
		if len(args) != 1 {
			return fmt.Errorf("Usage: select bills")
		}
		err = selectUserBills(s)
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

func selectUserBills(s *state) error {
	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Failed to get user '%v': %v", s.cfg.CurrentUserName, err)
	}

	inBills, err := s.db.GetInTransactionsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Failed to get income transactions for user '%v': %v", s.cfg.CurrentUserName, err)
	}

	outBills, err := s.db.GetOutTransactionsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Failed to get expenditure transactions for user '%v': %v", s.cfg.CurrentUserName, err)
	}

	fmt.Println()
	fmt.Println("Income:")

	for i, bill := range inBills {
		fmt.Println()
		printShortTransaction(bill)
		if i == len(inBills)-1 {
			fmt.Println()
		}
	}

	fmt.Println("Expenditures:")

	for i, bill := range outBills {
		fmt.Println()
		printShortTransaction(bill)
		if i == len(outBills)-1 {
			fmt.Println()
		}
	}

	return nil
}
