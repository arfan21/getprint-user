package modelresponse

import (
	"time"

	"github.com/arfan21/getprint-user/app/model/modeluser"
	uuid "github.com/satori/go.uuid"
)

type User struct {
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
	ProviderID    string    `json:"provider_id"`
	LastLogin     time.Time `json:"last_login"`
}

func (res *User) Set(user modeluser.User) {
	res.ID = user.ID
	res.CreatedAt = user.CreatedAt
	res.Name = user.Name
	res.Picture = user.Picture.String
	res.Email = user.Email
	res.EmailVerified = user.EmailVerified
	res.PhoneNumber = user.PhoneNumber.String
	res.Address = user.Address.String
	res.Role = user.Role
	res.Provider = user.Identities.Provider
	res.ProviderID = user.Identities.ProviderID
	res.LastLogin = user.UserLog.LastLogin.Time
}
