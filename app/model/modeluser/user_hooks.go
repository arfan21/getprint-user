package modeluser

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	u.Identities.UserID = u.ID
	u.UserLog.UserID = u.ID

	if u.Identities.Provider == "" {
		u.Identities.Provider = "getprint"
		u.Identities.ProviderID = u.ID.String()
	}

	err = nil
	return
}