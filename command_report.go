package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"Billsly/internal/database"
)

func commandReport(s *state, args ...string) error {
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

	initial := strings.ToUpper(string(s.cfg.CurrentUserName[0]))
	name := initial + string(s.cfg.CurrentUserName[1:])

	report := fmt.Sprintf("%v's Budget Report:\n", name)
	report += "------------------------\n\n"

	formatted, err := formatReport(s, inBills, outBills)
	if err != nil {
		return err
	}

	report += formatted

	fileName := fmt.Sprintf("%vbudget.txt", s.cfg.CurrentUserName)

	err = os.WriteFile(fileName, []byte(report), 0644)
	if err != nil {
		return fmt.Errorf("Failed to write to file '%v': %v", fileName, err)
	}

	fmt.Println(report)
	fmt.Printf("\n --- Report written to %v --- \n", fileName)

	return nil
}

func formatReport(s *state, income, expenditures []database.Transaction) (string, error) {
	incomeAmount := 0.00
	report := "Income:\n"
	for _, bill := range income {
		incomeAmount += bill.Amount
		report += fmt.Sprintf("\nName:       %v\nAmount:     %v\nDue:        %v\nBank:       %v\n", bill.TName, bill.Amount, bill.DueDate, bill.Bank)
	}

	expenditureAmount := 0.00

	report += "\nOutcome:\n"
	for _, bill := range expenditures {
		expenditureAmount += bill.Amount
		report += fmt.Sprintf("\nName:       %v\nAmount:     %v\nDue:        %v\nBank:       %v\n", bill.TName, bill.Amount, bill.DueDate, bill.Bank)
	}

	report += fmt.Sprintf("\nMonthly Allowance:  €%v\n", (incomeAmount - expenditureAmount))

	banks, err := s.db.GetAllBanks(context.Background())
	if err != nil {
		return "", fmt.Errorf("Failed to get banks list: %v", err)
	}

	for _, bank := range banks {
		cost := 0.00
		for _, bill := range expenditures {
			if bill.Bank == bank {
				cost += bill.Amount
			}
		}
		report += fmt.Sprintf("\n%v:  €%v", bank, cost)
	}

	return report, nil
}
