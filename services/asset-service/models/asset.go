package models

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	ID            string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Name          string    `json:"name" gorm:"not null"`
	Type          int32     `json:"type" gorm:"not null"`
	Brand         string    `json:"brand"`
	Model         string    `json:"model"`
	SerialNumber  string    `json:"serial_number" gorm:"uniqueIndex"`
	PurchaseDate  string    `json:"purchase_date"`
	PurchasePrice float64   `json:"purchase_price"`
	Status        int32     `json:"status" gorm:"default:0"`
	AssignedTo    string    `json:"assigned_to"`
	AssignedDate  string    `json:"assigned_date"`
	Location      string    `json:"location"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (a *Asset) BeforeCreate() error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	return nil
}
