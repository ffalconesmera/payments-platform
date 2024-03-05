package dto

type Merchant struct {
	MerchantCode string `json:"merchant_code"`
	Name         string `json:"name"`
}

type JSONMerchant struct {
	Status   string   `json:"status"`
	Merchant Merchant `json:"data"`
}
