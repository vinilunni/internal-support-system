package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Picture   string    `json:"picture"`
	Domain    string    `json:"domain" gorm:"not null"`
	Active    bool      `json:"active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RefreshToken struct {
	ID        string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	UserID    string    `json:"user_id" gorm:"not null"`
	Token     string    `json:"token" gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate() error {
	if u.ID == "" {
		u.ID = generateID()
	}
	return nil
}

func (r *RefreshToken) BeforeCreate() error {
	if r.ID == "" {
		r.ID = generateID()
	}
	return nil
}

func generateID() string {
	return "usr_" + time.Now().Format("20060102150405")
}
