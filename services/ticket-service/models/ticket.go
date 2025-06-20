package models

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID          string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Category    int32     `json:"category" gorm:"not null"`
	Priority    int32     `json:"priority" gorm:"not null"`
	Status      int32     `json:"status" gorm:"default:0"`
	RequesterID string    `json:"requester_id" gorm:"not null"`
	AssigneeID  string    `json:"assignee_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ResolvedAt  string    `json:"resolved_at"`
}

type TicketComment struct {
	ID         string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	TicketID   string    `json:"ticket_id" gorm:"not null"`
	UserID     string    `json:"user_id" gorm:"not null"`
	Content    string    `json:"content" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	IsInternal bool      `json:"is_internal" gorm:"default:false"`
	Ticket     Ticket    `json:"ticket" gorm:"foreignKey:TicketID"`
}

func (t *Ticket) BeforeCreate() error {
	if t.ID == "" {
		t.ID = uuid.New().String()
	}
	return nil
}

func (tc *TicketComment) BeforeCreate() error {
	if tc.ID == "" {
		tc.ID = uuid.New().String()
	}
	return nil
}
