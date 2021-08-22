package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is the user model that is used for the API and database tables
type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	Email     string    `json:"email" gorm:"not null;unique;size:256"`
	Username  string    `json:"username" gorm:"not null;unique;size:256"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// BeforeCreate is run just before inserting a new user in the users table.
// It create a unique id for the user!
func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()

	return nil
}
