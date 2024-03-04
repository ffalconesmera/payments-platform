package dto

type User struct {
	UUID     string `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
