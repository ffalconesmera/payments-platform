package dto

type Merchant struct {
	UUID         string `json:"-"`
	MerchantCode string `json:"merchant_code"`
	Name         string `json:"name"`
	User         *User  `json:"user,omitempty"`
}
