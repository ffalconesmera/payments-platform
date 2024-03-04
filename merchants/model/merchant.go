package model

type PayMerchant struct {
	UUID         string `gorm:"column:pay_merch_id; unique; not null" json:"-"`
	MerchantCode string `gorm:"column:pay_merch_code; unique; not null" json:"merchant_code"`
	Name         string `gorm:"column:pay_merch_name" json:"name"`
	BaseModel
}
