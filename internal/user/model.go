package user

import "time"

type Role string

const (
	RoleAdmin   Role = "ADMIN"
	RoleDoctor  Role = "DOCTOR"
	RolePatient Role = "PATIENT"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	FullName     string
	Roles        []Role
	Enabled      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"fullName"`
	Roles     []Role    `json:"roles"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"createdAt"`
}
