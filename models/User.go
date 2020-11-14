package models

import (
	"time"

	validator "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gopkg.in/guregu/null.v4"
	"gopkg.in/guregu/null.v4/zero"
)

type User struct {
	ID          uint        `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DeletedAt   null.Time   `gorm:"index" json:"deleted_at,omitempty"`
	UserIDLine  null.String `gorm:"unique" json:"user_id_line"`
	Name        string      `gorm:"size:255;not null;" json:"name"`
	Picture     null.String `gorm:"size:255;" json:"picture"`
	Email       string      `gorm:"size:255;not null;unique;" json:"email"`
	Password    null.String `gorm:"size:255;" json:"password"`
	PhoneNumber zero.Int    `json:"phone_number"`
	Address     null.String `gorm:"type:longtext" json:"address"`
	Role        string      `gorm:"type:enum('admin','buyer','seller');default:'buyer'" json:"role"`
}

func (u User) Validate() error {
	return validator.ValidateStruct(&u,
		validator.Field(&u.Name, validator.Required),
		validator.Field(&u.Email, validator.Required, is.Email),
		validator.Field(&u.Password, validator.Length(8, 20)),
	)
}

var (
	ErrorEmailRegistered = "email already taken"
	ErrorEmailNotFound   = "email not registered"
)

type UserRepository interface {
	Create(user *User) error
	Get(users *[]User) error
	GetByID(id uint, user *User) error
	Update(user *User) error
}

type UserService interface {
	Create(user *User) error
	Get(users *[]User) error
	GetByID(id uint, user *User) error
	Update(user *User) error
}
