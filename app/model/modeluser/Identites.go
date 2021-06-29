package modeluser

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

type Identities struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  null.Time `gorm:"index" json:"deleted_at,omitempty"`
	UserID     uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	Provider   string    `gorm:"size:255;not null" json:"provider"`
	ProviderID string    `gorm:"size:255;not null" json:"provider_id"`
}
