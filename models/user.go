package model 

import "time"

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
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

// type Nostrevent struct {
// 	Id int `json:"id"`
// 	PubKey string `json:"pubkey"`
// 	Kind string `json:"kind"`
// 	Ptags string[] `json:ptags`
// 	Etags string[] `json:etags`
// 	Expiration time.Time `json:"expiration"`
// 	Gtags string[] `json:"gtags"`
// 	Content string `json:"content"`
// }

// id SERIAL PRIMARY KEY,
// pubkey TEXT NOT NULL,
// created_at TIMESTAMP NOT NULL,
// kind TEXT,
// ptags TEXT[],
// etags TEXT[],
// dtag TEXT,
// expiration TIMESTAMP,
// gtags TEXT[],
// raw BYTEA
