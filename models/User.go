package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID            uuid.UUID   `gorm:"primary_key;type:char(36)" json:"id"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     null.Time   `gorm:"index" json:"deleted_at,omitempty"`
	Name          string      `gorm:"size:100;not null;" json:"name"`
	Picture       null.String `gorm:"size:255;" json:"picture"`
	Email         string      `gorm:"size:255;not null;unique;" json:"email"`
	EmailVerified bool        `gorm:"default:0" json:"email_verified"`
	Password      null.String `gorm:"size:255;" json:"password"`
	PhoneNumber   null.String `gorm:"unique" json:"phone_number"`
	Address       null.String `gorm:"type:longtext" json:"address"`
	Role          string      `gorm:"type:enum('admin','buyer','seller');default:'buyer'" json:"role"`
	Identities    Identities  `gorm:"constraint:OnDelete:CASCADE;" json:"identities,omitempty"`
	UserLog       UserLog     `gorm:"constraint:OnDelete:CASCADE;" json:"user_log,omitempty"`
}

type Identities struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      null.Time `gorm:"index" json:"deleted_at,omitempty"`
	UserID         uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	Provider       string    `gorm:"size:255;not null" json:"provider"`
	UserIDProvider string    `gorm:"size:255;not null" json:"user_id_provider"`
}

type UserLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt null.Time `gorm:"index" json:"deleted_at,omitempty"`
	UserID    uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	LastLogin null.Time `json:"last_login"`
}
