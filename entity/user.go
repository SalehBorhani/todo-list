package entity

type User struct {
	ID          uint8  `json:"id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
