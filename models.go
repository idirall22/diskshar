package user

import "time"

// User model
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"Email"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}