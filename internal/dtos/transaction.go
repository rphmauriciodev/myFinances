package dtos

import "fmt"

type Transaction struct {
	Name       string `json:"name"`
	Merchant   string `json:"merchant"`
	CardOrPass string `json:"cardOrPass"`
	Amount     string `json:"amount"`
	Date       string `json:"date"`
}

func (t *Transaction) Validate() error {
	// Implement validation logic for the transaction fields here
	// For example, you might want to check if the required fields are not empty,
	// if the amount is a valid number, or if the date is in a valid format.

	if t.Name == "" {
		return fmt.Errorf("transaction name is required")
	}
	if t.Merchant == "" {
		return fmt.Errorf("merchant is required")
	}
	if t.CardOrPass == "" {
		return fmt.Errorf("card or pass is required")
	}
	if t.Amount == "" {
		return fmt.Errorf("amount is required")
	}
	if t.Date == "" {
		return fmt.Errorf("date is required")
	}

	return nil
}
