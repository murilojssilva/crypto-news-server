package models

import "time"

type User struct {
	ID int `json:"id"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Password_confirmation string `json:"password_confirmation"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}