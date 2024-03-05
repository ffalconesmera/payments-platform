package model

import "time"

type TransactionStatus string

var (
	TransactionStatusOK        TransactionStatus = "ok"
	TransactionStatusSucceeded TransactionStatus = "succeeded"
	TransactionStatusFailure   TransactionStatus = "failure"
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusRefunded  TransactionStatus = "refunded"
)

type PayTransaction struct {
	UUID                     string            `gorm:"column:pay_tx_id; unique; not null" json:"-"`
	PaymentCode              string            `gorm:"column:pay_tx_code; unique; not null" json:"payment_code"`
	Amount                   float64           `gorm:"column:pay_tx_amount" json:"amount"`
	Description              string            `gorm:"column:pay_tx_description" json:"description"`
	Currency                 string            `gorm:"column:pay_tx_currency" json:"currency"`
	Status                   TransactionStatus `gorm:"column:pay_tx_status" json:"status"`
	ExpirationProcess        *time.Time        `gorm:"column:pay_tx_expiration_process" json:"-"`
	NaturalExpirationProcess string            `gorm:"column:pay_tx_natural_expiration_process" json:"natural_expiration_process"`
	FailureReason            *string           `gorm:"column:pay_tx_failure_reason" json:"failure_reason,omitempty"`
	BankReference            *string           `gorm:"column:pay_tx_bank_reference;" json:"bank_reference,omitempty"`
	BankName                 *string           `gorm:"column:pay_tx_bank_name;" json:"bank_name,omitempty"`
	MerchantCode             string            `gorm:"column:pay_tx_merchant_code" json:"-"`
	CustomerUUID             string            `gorm:"column:pay_tx_customer_id;" json:"-"`
	RefundUUID               *string           `gorm:"column:pay_tx_refund_id" json:"-"`
	BaseModel
}
