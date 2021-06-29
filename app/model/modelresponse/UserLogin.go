package modelresponse

import uuid "github.com/satori/go.uuid"

type UserLogin struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Role  string    `json:"role"`
}
