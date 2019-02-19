// Package model contains the data level logic and entities
package model

import "fmt"

// TXNotFound represents a missing transaction
type TXNotFound struct {
	ID string
}

func (tx *TXNotFound) Error() string {
	return fmt.Sprintf("TX with ID %s not found", tx.ID)
}

// TXDatabaseEmpty represents missing transactions
type TXDatabaseEmpty struct{}

func (tx *TXDatabaseEmpty) Error() string {
	return fmt.Sprintf("No TXs found in the database")
}

// TXInvalid represents a transaction missing manadatory data
type TXInvalid struct {
	Reason string
}

func (tx *TXInvalid) Error() string {
	return fmt.Sprintf("Invalid or incomplete transaction")
}
