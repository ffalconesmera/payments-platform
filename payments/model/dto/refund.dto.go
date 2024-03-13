package dto

type Refund struct {
	Code          string `json:"code"`
	BankReference string `json:"reference"`
	Date          string `json:"date"`
}
