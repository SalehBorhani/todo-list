package request

type UserRegister struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
