package models

import (
	"time"

	"github.com/google/uuid"
)

type Assignment struct {
	ID                 string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	AssetID            string    `json:"asset_id" gorm:"not null"`
	UserID             string    `json:"user_id" gorm:"not null"`
	AssignedBy         string    `json:"assigned_by" gorm:"not null"`
	AssignedDate       string    `json:"assigned_date" gorm:"not null"`
	ExpectedReturnDate string    `json:"expected_return_date"`
	ActualReturnDate   string    `json:"actual_return_date"`
	Status             int32     `json:"status" gorm:"default:0"`
	Notes              string    `json:"notes"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (a *Assignment) BeforeCreate() error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	return nil
}
