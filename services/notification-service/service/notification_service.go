package service

import (
	"context"
	"fmt"
	"time"

	notificationpb "internal-support-system/proto/notification"

	"gorm.io/gorm"
)

type NotificationService struct {
	db *gorm.DB
}

type NotificationTemplate struct {
	ID        string            `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Name      string            `json:"name" gorm:"not null"`
	Subject   string            `json:"subject" gorm:"not null"`
	Body      string            `json:"body" gorm:"not null"`
	Type      int32             `json:"type" gorm:"not null"`
	Variables map[string]string `json:"variables" gorm:"serializer:json"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type Notification struct {
	ID        string            `json:"id" gorm:"primaryKey;type:varchar(36)"`
	UserID    string            `json:"user_id" gorm:"not null"`
	Title     string            `json:"title" gorm:"not null"`
	Message   string            `json:"message" gorm:"not null"`
	Type      int32             `json:"type" gorm:"not null"`
	Priority  int32             `json:"priority" gorm:"not null"`
	Read      bool              `json:"read" gorm:"default:false"`
	CreatedAt time.Time         `json:"created_at"`
	ReadAt    *time.Time        `json:"read_at"`
	Metadata  map[string]string `json:"metadata" gorm:"serializer:json"`
}

func NewNotificationService(db *gorm.DB) *NotificationService {
	db.AutoMigrate(&Notification{}, &NotificationTemplate{})
	return &NotificationService{db: db}
}

func (s *NotificationService) SendNotification(ctx context.Context, req *notificationpb.SendNotificationRequest) (*notificationpb.SendNotificationResponse, error) {
	fmt.Println("Sending notification...")
	return &notificationpb.SendNotificationResponse{
		Notification: &notificationpb.Notification{
			Id:      "notification-id",
			Title:   req.Title,
			Message: req.Message,
		},
	}, nil
}

func (s *NotificationService) GetNotifications(ctx context.Context, req *notificationpb.GetNotificationsRequest) (*notificationpb.GetNotificationsResponse, error) {
	fmt.Println("Getting notifications...")
	return &notificationpb.GetNotificationsResponse{
		Notifications: []*notificationpb.Notification{},
		Total:         0,
	}, nil
}

func (s *NotificationService) MarkAsRead(ctx context.Context, req *notificationpb.MarkAsReadRequest) (*notificationpb.MarkAsReadResponse, error) {
	fmt.Println("Marking notification as read...")
	return &notificationpb.MarkAsReadResponse{
		Success: true,
	}, nil
}

func (s *NotificationService) CreateTemplate(ctx context.Context, req *notificationpb.CreateTemplateRequest) (*notificationpb.CreateTemplateResponse, error) {
	fmt.Println("Creating notification template...")
	return &notificationpb.CreateTemplateResponse{
		Id: "template-id",
	}, nil
}

func (s *NotificationService) UpdateTemplate(ctx context.Context, req *notificationpb.UpdateTemplateRequest) (*notificationpb.UpdateTemplateResponse, error) {
	fmt.Println("Updating notification template...")
	return &notificationpb.UpdateTemplateResponse{
		Success: true,
	}, nil
}
