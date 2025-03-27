package model

type User struct {
	Id           uint   `json:"id"`
	UserId       string `json:"user_id"`
	Otp          string `json:"otp"`
	StartDateOtp string `json:"start_date_otp"`
}

type RequestUser struct {
	Username string `json:"username"`
}

type ResponseUser struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

type UserUri struct {
	ID uint `uri:"id" binding:"required,number"`
}
