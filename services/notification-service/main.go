package main

import (
	"context"
	"log"
	"net"

	"internal-support-system/pkg/config"
	notificationpb "internal-support-system/proto/notification"
	"internal-support-system/services/notification-service/service"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type server struct {
	notificationpb.UnimplementedNotificationServiceServer
	notificationService *service.NotificationService
}

func (s *server) SendNotification(ctx context.Context, req *notificationpb.SendNotificationRequest) (*notificationpb.SendNotificationResponse, error) {
	return s.notificationService.SendNotification(ctx, req)
}

func (s *server) GetNotifications(ctx context.Context, req *notificationpb.GetNotificationsRequest) (*notificationpb.GetNotificationsResponse, error) {
	return s.notificationService.GetNotifications(ctx, req)
}

func (s *server) MarkAsRead(ctx context.Context, req *notificationpb.MarkAsReadRequest) (*notificationpb.MarkAsReadResponse, error) {
	return s.notificationService.MarkAsRead(ctx, req)
}

func (s *server) CreateTemplate(ctx context.Context, req *notificationpb.CreateTemplateRequest) (*notificationpb.CreateTemplateResponse, error) {
	return s.notificationService.CreateTemplate(ctx, req)
}

func (s *server) UpdateTemplate(ctx context.Context, req *notificationpb.UpdateTemplateRequest) (*notificationpb.UpdateTemplateResponse, error) {
	return s.notificationService.UpdateTemplate(ctx, req)
}

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	notificationService := service.NewNotificationService(db)

	lis, err := net.Listen("tcp", ":9095")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	s := grpc.NewServer()
	notificationpb.RegisterNotificationServiceServer(s, &server{notificationService: notificationService})

	log.Println("Notification service listening on port 9095")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
