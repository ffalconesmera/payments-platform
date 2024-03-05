package model

import "time"

type PayRefund struct {
	UUID          string     `gorm:"column:pay_ref_id; unique; not null" json:"-"`
	Code          string     `gorm:"column:pay_ref_code" json:"code,omitempty"`
	BankReference string     `gorm:"column:pay_ref_bank_reference" json:"bank_reference,omitempty"`
	Date          *time.Time `gorm:"column:pay_ref_date" json:"-"`
	BaseModel
}
