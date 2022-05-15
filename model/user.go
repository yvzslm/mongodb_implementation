package model

import "time"

type User struct {
	FullName    string    `json:"fullName" bson:"FullName"`
	Identity    string    `json:"identity" bson:"Identity"`
	PhoneNumber string    `json:"phoneNumber" bson:"PhoneNumber"`
	Email       string    `json:"email" bson:"Email"`
	Gender      string    `json:"gender" bson:"Gender"`
	IsActive    bool      `json:"isActive" bson:"IsActive"`
	IsDeleted   bool      `json:"isDeleted" bson:"IsDeleted"`
	CreatedDate time.Time `json:"createdDate" bson:"CreatedDate"`
}

func NewUser(fullName, identity, phoneNumber, email, gender string) *User {
	return &User{
		FullName:    fullName,
		Identity:    identity,
		PhoneNumber: phoneNumber,
		Email:       email,
		Gender:      gender,
		IsActive:    true,
		IsDeleted:   false,
		CreatedDate: time.Now(),
	}
}
