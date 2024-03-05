package dto

type Customer struct {
	UUID    string `json:"-"`
	DNI     string `json:"dni"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
