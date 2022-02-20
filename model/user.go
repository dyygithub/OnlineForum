package model

type UserAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Userinfo struct {
	User_id string `json:"user_id"`
	Avatar string `json:"avatar"`
	Nickname string `json:"nickname"`
	Introduction string `json:"introduction"`
	Phone string `json:"phone"`
	Qq string `json:"qq"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Birthday string `json:"birthday"`
	username string `json:"username"`
}
