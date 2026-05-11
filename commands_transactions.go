package main

import (
	"Billsly/internal/database"
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var validBanks = []string{"aib", "boi", "ptsb", "rev"}

func commandCreateTransaction(s *state, args ...string) error {
	if len(args) != 5 {
		return fmt.Errorf("Usage: bill <name> <type> <amount> <due_date (dd)> <bank>")
	}

	name := args[0]
	transactionType := args[1]
	if strings.ToLower(transactionType) != "in" && strings.ToLower(transactionType) != "out" {
		return fmt.Errorf("Please enter valid transaction type argument ('in' for receiving amount or 'out' for spending amount)")
	}

	amountParts := strings.Split(args[2], ".")
	if len(amountParts) == 2 && len(amountParts[1]) > 2 {
		return fmt.Errorf("Please enter a valid number with no more than 2 decimal places: %v", args[2])
	}
	amount, err := strconv.ParseFloat(args[2], 2)
	if err != nil {
		return fmt.Errorf("Failed to convert amount argument to valid float: %v", err)
	}

	due_date := args[3]
	date, err := strconv.Atoi(due_date)
	if err != nil {
		return fmt.Errorf("Failed to convert date argument to valid integer: %v", err)
	}
	if date < 1 || date > 31 {
		return fmt.Errorf("Please enter date in the format <dd>, where <dd> is a whole number between 1 and 31: %v", date)
	}
	day := int32(date)

	bank := args[4]
	if !slices.Contains(validBanks, strings.ToLower(bank)) {
		return fmt.Errorf("Please enter valid bank argument from the following list: %v", validBanks)
	}

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Couldn't get current user ID: %v", err)
	}

	transaction, err := s.db.CreateTransaction(context.Background(), database.CreateTransactionParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Type:      transactionType,
		Amount:    amount,
		DueDate:   day,
		Bank:      bank,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("Couldn't create bill: %v", err)
	}

	fmt.Println("Bill created successfully:")
	printTransaction(transaction)
	return nil
}

func commandDeleteTransaction(s *state, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: delete <name>")
	}

	name := args[0]
	bill, err := s.db.GetTransactionByName(context.Background(), name)
	if err != nil {
		return fmt.Errorf("Failed to get transaction '%v': %v", name, err)
	}

	err = s.db.DeleteTransaction(context.Background(), bill.ID)
	if err != nil {
		return fmt.Errorf("Failed to delete transaction '%v': %v", name, err)
	}

	return nil
}

func printTransaction(t database.Transaction) {
	fmt.Printf(" * ID:         %v\n", t.ID)
	fmt.Printf(" * Name:       %v\n", t.Name)
	fmt.Printf(" * Type:       %v\n", t.Type)
	fmt.Printf(" * Amount:     €%v\n", t.Amount)
	if len(fmt.Sprintf("%v", t.DueDate)) == 1 {
		fmt.Printf(" * DueDate:    0%v\n", t.DueDate)
	} else {
		fmt.Printf(" * DueDate:    %v\n", t.DueDate)
	}
	fmt.Printf(" * Bank:       %v\n", t.Bank)
	fmt.Printf(" * UserID:     %v\n", t.UserID)
}

func printShortTransaction(t database.Transaction) {
	fmt.Printf(" * Name:      %v\n", t.Name)
	fmt.Printf(" * Amount:    €%v\n", t.Amount)
	fmt.Printf(" * Type:      %v\n", t.Type)
}
