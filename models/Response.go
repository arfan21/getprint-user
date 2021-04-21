package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserResoponse struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	Name          string    `json:"name"`
	Picture       string    `json:"picture"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	PhoneNumber   string    `json:"phone_number"`
	Address       string    `json:"address"`
	Role          string    `json:"role"`
	Provider      string    `json:"provider"`
	LastLogin     time.Time `json:"last_login"`
}

type UserLoginResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Role  string    `json:"role"`
}
