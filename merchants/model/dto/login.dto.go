package dto

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Message  string `json:"message"`
}
