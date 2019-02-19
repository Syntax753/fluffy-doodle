package models

import "fmt"

// TXNotFound is custom error or
type TXNotFound struct {
	ID string
}

func (tx *TXNotFound) Error() string {
	return fmt.Sprintf("TX with ID %s not found", tx.ID)
}
