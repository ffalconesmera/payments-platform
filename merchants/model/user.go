package model

type PayUser struct {
	UUID         string `gorm:"column:pay_usr_id; unique; not null" json:"-"`
	Username     string `gorm:"column:pay_usr_username; unique; not null" json:"username"`
	Email        string `gorm:"column:pay_usr_email" json:"email"`
	Password     string `gorm:"column:pay_usr_password" json:"password,omitempty"`
	MerchantUUID string `gorm:"column:pay_usr_merchant_id;not null" json:"-"`
	BaseModel
}
