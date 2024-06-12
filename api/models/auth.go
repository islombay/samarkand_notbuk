package models

type Login struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type LoginClient struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type VerifyLogin struct {
	RequestID string `json:"request_id" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

type ProfileLogin struct {
	RequestID string `json:"request_id" binding:"required"`
	FirstName string `json:"first_name" binding:"required,max=40"`
	LastName  string `json:"last_name" binding:"required,max=40"`
}
