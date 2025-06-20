package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Email      string    `json:"email" gorm:"uniqueIndex;not null"`
	Name       string    `json:"name" gorm:"not null"`
	Department string    `json:"department"`
	Role       string    `json:"role"`
	ManagerID  string    `json:"manager_id"`
	Active     bool      `json:"active" gorm:"default:true"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Manager    *User     `json:"manager,omitempty" gorm:"foreignKey:ManagerID"`
}

func (u *User) BeforeCreate() error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}
