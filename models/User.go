package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gopkg.in/guregu/null.v4/zero"
)

type User struct {
	ID            uint        `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	DeletedAt     null.Time   `gorm:"index" json:"deleted_at,omitempty"`
	UserIDLine    null.String `gorm:"unique" json:"user_id_line"`
	Name          string      `gorm:"size:100;not null;" json:"name"`
	Picture       null.String `gorm:"size:255;" json:"picture"`
	Email         string      `gorm:"size:255;not null;unique;" json:"email"`
	EmailVerified bool        `gorm:"default:false" json:"email_verified"`
	Password      null.String `gorm:"size:255;" json:"password"`
	PhoneNumber   zero.Int    `gorm:"unique" json:"phone_number"`
	Address       null.String `gorm:"type:longtext" json:"address"`
	LastLogin     time.Time   `json:"last_login"`
	Role          string      `gorm:"type:enum('admin','buyer','seller');default:'buyer'" json:"role"`
}

type Identities struct {
	Provider   string `json:"provider"`
	UserID     string `json:"user_id"`
	Connection string `json:"connection"`
	IsSocial   bool   `json:"is_social"`
}

type UserRepository interface {
	Create(user *User) error
	Get(users *[]User) error
	GetByID(id uint, user *User) error
	GetByEmail(user *User) error
	Update(user *User) error
}

type UserService interface {
	Create(user *User) error
	Get(users *[]User) error
	GetByID(id uint, user *User) error
	Update(user *User) error
	Login(user *User) error
}
