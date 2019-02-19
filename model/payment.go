package model

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// Value represents a financial value
type Value string

// Currency represents a financial denomination
type Currency string

// Amount represents an amount of certain denomination
type Amount struct {
	Amount   Value    `json:"amount"`
	Currency Currency `json:"currency"`
}

// Account represents a funds account (such as a bank account)
type Account struct {
	AccountName       string `json:"account_name,omitempty"`
	AccountNumber     string `json:"account_number,omitempty"`
	AccountNumberCode string `json:"account_number_code,omitempty"`
	AccountType       int    `json:"account_type,omitempty"`
	Address           string `json:"address,omitempty"`
	BankID            string `json:"bank_id,omitempty"`
	BankIDCode        string `json:"bank_id_code,omitempty"`
	Name              string `json:"name,omitempty"`
}

// ChargesInformation represents any charges associated with a transaction
type ChargesInformation struct {
	BearerCode              string   `json:"bearer_code"`
	SenderCharges           []Amount `json:"sender_charges"`
	ReceiverChargesAmount   Value    `json:"receiver_charges_amount"`
	ReceiverChargesCurrency Currency `json:"receiver_charges_currency"`
}

// FX represents the market exchange rate
type FX struct {
	ContractReference string   `json:"contract_reference"`
	ExchangeRate      string   `json:"exchange_rate"`
	OriginalAmount    Value    `json:"original_amount"`
	OriginalCurrency  Currency `json:"original_currency"`
}

// Attributes represents the transaction meta data
type Attributes struct {
	Amount               Value               `json:"amount"`
	BeneficiaryParty     *Account            `json:"beneficiary_party"`
	ChargesInformation   *ChargesInformation `json:"charges_information"`
	Currency             Currency            `json:"currency"`
	DebtorParty          *Account            `json:"debtor_party"`
	EndToEndReference    string              `json:"end_to_end_reference"`
	Fx                   *FX                 `json:"fx"`
	NumericReference     string              `json:"numeric_reference"`
	PaymentID            string              `json:"payment_id"`
	PaymentPurpose       string              `json:"payment_purpose"`
	PaymentScheme        string              `json:"payment_scheme"`
	PaymentType          string              `json:"payment_type"`
	ProcessingDate       string              `json:"processing_date"`
	Reference            string              `json:"reference"`
	SchemePaymentSubType string              `json:"scheme_payment_sub_type"`
	SchemePaymentType    string              `json:"scheme_payment_type"`
	SponsorParty         *Account            `json:"sponsor_party"`
}

// TX represents one single transaction - typically a payment
type TX struct {
	Type           string      `json:"type"`
	ID             string      `json:"id"`
	Version        int         `json:"version"`
	OrganisationID string      `json:"organisation_id"`
	Attributes     *Attributes `json:"attributes"`
}

// Data contains an array of transactions
type Data struct {
	TXs []TX `json:"data"`
}

// GetAllTX is responsible for retrieving all the transactions
// TODO pagination
func (db *DB) GetAllTX() ([]*TX, error) {
	db.Lock()
	defer db.Unlock()

	if db.TXs == nil || len(db.TXs) == 0 {
		return nil, &TXDatabaseEmpty{}
	}

	m := make([]*TX, 0, len(db.TXs))

	for _, val := range db.TXs {
		m = append(m, val)
	}

	return m, nil
}

// GetTX is responsible for retrieving a transaction
func (db *DB) GetTX(ID string) (*TX, error) {
	db.Lock()
	defer db.Unlock()

	if v, ok := db.TXs[ID]; ok {
		return v, nil
	}

	return nil, &TXNotFound{ID}
}

// CreateTX is responsible for creating a transaction
func (db *DB) CreateTX(tx TX) (*TX, error) {
	// Arbirtrary value that should be mandatory
	if tx.Type != "Payment" {
		return nil, &TXInvalid{}
	}

	// Generate id if mising
	if tx.ID == "" {
		tx.ID = fmt.Sprintf("%s", uuid.Must(uuid.NewV4()))
	}

	db.Lock()
	defer db.Unlock()
	db.TXs[tx.ID] = &tx

	return &tx, nil
}

// UpdateTX is responsible for updating a transaction
func (db *DB) UpdateTX(tx TX) (*TX, error) {
	// Arbirtrary value that should be mandatory
	if tx.Type != "Payment" {
		return nil, &TXInvalid{}
	}

	// ID mandatory
	if tx.ID == "" {
		return nil, &TXInvalid{}
	}

	if _, ok := db.TXs[tx.ID]; !ok {
		return nil, &TXNotFound{tx.ID}
	}

	db.Lock()
	defer db.Unlock()
	db.TXs[tx.ID] = &tx

	return &tx, nil
}

// DeleteTX is responsible for deleting a transaction
func (db *DB) DeleteTX(ID string) error {
	if _, ok := db.TXs[ID]; !ok {
		return &TXNotFound{ID}
	}

	db.Lock()
	defer db.Unlock()
	delete(db.TXs, ID)

	return nil
}
