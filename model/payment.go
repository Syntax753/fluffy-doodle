package model

// Amount is a financial value along with its currency
type Amount struct {
	Amount   float64
	Currency string
}

// Account represents a paying account (such as a bank account)
type Account struct {
	AccountName       string
	AccountNumber     string
	AccountNumberCode string
	AccountType       int // This will no doubt be a FK to a separate table

}

// TX represents one single transaction - typically a payment
type TX struct {
	Type    string
	ID      string `json:"id"`
	Version int
}
