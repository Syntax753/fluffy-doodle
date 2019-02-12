package model

// TX represents one single transaction - typically a payment
type TX struct {
	Type    string
	ID      string `json:"id"`
	Version int
}
