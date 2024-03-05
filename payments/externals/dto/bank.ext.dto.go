package dto

type BankPayment struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
}

type BankRefund struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
}

type BankCardInput struct {
	CardNumber       string `json:"card_number"`
	PaymentReference string `json:"payment_reference"`
}

type RefundInput struct {
	RefundCase       string `json:"refund_case"`
	PaymentReference string `json:"payment_reference"`
}

type JSONBankPayment struct {
	Status      string      `json:"status"`
	BankPayment BankPayment `json:"data"`
}

type JSONBankRedfund struct {
	Status      string          `json:"status"`
	BankPayment JSONBankPayment `json:"data"`
}
