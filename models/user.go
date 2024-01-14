package model 

import "time"

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Measurements struct {
	Id int `json:"id"`
	UserID int `json:"user_id"`
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
	BodyFat float64 `json:"body_fat"`
	CreatedAt time.Time `json:"created_at"`
}