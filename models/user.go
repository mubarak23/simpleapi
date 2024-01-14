package model 

import "time"

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"email"`
	CreatedAt time.time `json:"created_at"`
	UpdatedAt time.time `json:"updated_at"`
}

type Measurement struct {
	Id int `json:"id"`
	userId int `json:"user_id"`
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
	BodyFat float64 `json:"body_fat"`
	CreatedAt time.time `json:"created_at"`
}