package userAccount

import "time"

// RegisterForm model
type RegisterForm struct {
	Username string `json:"username"`
	Email    string `json:"Email"`
	Password string `json:"password"`
}

// LoginForm model
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ValidLoginForm model
type ValidLoginForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User model
type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"Email"`
	Password  string     `json:"password"`
	Avatar    string     `json:"avatar"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
