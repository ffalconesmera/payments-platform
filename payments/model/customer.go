package model

type PayCustomer struct {
	UUID    string `gorm:"column:pay_cust_id; unique; not null" json:"-"`
	DNI     string `gorm:"column:pay_cust_dni" json:"dni"`
	Name    string `gorm:"column:pay_cust_name" json:"name"`
	Email   string `gorm:"column:pay_cust_email" json:"email"`
	Phone   string `gorm:"column:pay_cust_phone" json:"phone"`
	Address string `gorm:"column:pay_cust_address" json:"address"`
	BaseModel
}
