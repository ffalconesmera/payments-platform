package dto

import "time"

type Payment struct {
	UUID                     string     `json:"-"`
	PaymentCode              string     `json:"payment_code"`
	Amount                   float64    `json:"amount"`
	Description              string     `json:"description"`
	Currency                 string     `json:"currency"`
	Status                   string     `json:"status"`
	ExpirationProcess        *time.Time `json:"-"`
	NaturalExpirationProcess string     `json:"natural_expiration_process"`
	FailureReason            *string    `json:"failure_reason,omitempty"`
	BankReference            *string    `json:"bank_reference,omitempty"`
	BankName                 *string    `json:"bank_name,omitempty"`
	MerchantCode             string     `json:"-"`
	CustomerUUID             string     `json:"-"`
	RefundUUID               *string    `json:"-"`
	Customer                 Customer   `json:"customer"`
	Refund                   *Refund    `json:"refund,omitempty"`
}
