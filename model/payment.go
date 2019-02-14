package model

// Value represents a financial value
type Value string

// Currency represents a financial denomination
type Currency string

// Amount represents an amount of certain denomination
type Amount struct {
	Amount   Value
	Currency Currency
}

// Account represents a funds account (such as a bank account)
type Account struct {
	AccountName       string
	AccountNumber     string
	AccountNumberCode string
	AccountType       int
	Address           string
	BankID            string `json:"bank_id"`
	BankIDCode        string
	Name              string
}

// ChargesInformation represents any charges associated with a transaction
type ChargesInformation struct {
	BearerCode              string
	ReceiverChargesAmount   Value
	ReceiverChargesCurrency Currency
	SenderCharges           []Amount
}

// FX represents the market exchange rate
type FX struct {
	ContractReference string
	ExchangeRate      string
	OriginalAmount    Value
	OriginalCurrency  Currency
}

// Attributes represents the transaction meta data
type Attributes struct {
	Amount               Value
	Currency             Currency
	EndToEndReference    string
	NumericReference     string
	PaymentID            string `json:"payment_id"`
	PaymentPurpose       string
	PaymentScheme        string
	PaymentType          string
	ProcessingDate       string
	Reference            string
	SchemePaymentSubType string
	SchemePaymentType    string
	BeneficiaryParty     Account
	ChargesInformation   ChargesInformation
	DebtorParty          Account
	SponsorParty         Account
	Fx                   FX
}

// TX represents one single transaction - typically a payment
type TX struct {
	Type           string
	ID             string `json:"id"`
	Version        int
	OrganisationID string `json:"organisation_id"`
	// Attributes     *Attributes `json:",omitempty"`
	Attributes Attributes
}

// Data contains an array of transactions
type Data struct {
	TXs []TX `json:"data"`
}
