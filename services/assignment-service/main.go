package main

import (
	"context"
	"log"
	"net"

	"internal-support-system/pkg/config"
	assignmentpb "internal-support-system/proto/assignment"
	"internal-support-system/services/assignment-service/service"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type server struct {
	assignmentpb.UnimplementedAssignmentServiceServer
	assignmentService *service.AssignmentService
}

func (s *server) AssignAsset(ctx context.Context, req *assignmentpb.AssignAssetRequest) (*assignmentpb.AssignAssetResponse, error) {
	return s.assignmentService.AssignAsset(ctx, req)
}

func (s *server) UnassignAsset(ctx context.Context, req *assignmentpb.UnassignAssetRequest) (*assignmentpb.UnassignAssetResponse, error) {
	return s.assignmentService.UnassignAsset(ctx, req)
}

func (s *server) GetAssignment(ctx context.Context, req *assignmentpb.GetAssignmentRequest) (*assignmentpb.GetAssignmentResponse, error) {
	return s.assignmentService.GetAssignment(ctx, req)
}

func (s *server) ListAssignments(ctx context.Context, req *assignmentpb.ListAssignmentsRequest) (*assignmentpb.ListAssignmentsResponse, error) {
	return s.assignmentService.ListAssignments(ctx, req)
}

func (s *server) GetUserAssignments(ctx context.Context, req *assignmentpb.GetUserAssignmentsRequest) (*assignmentpb.GetUserAssignmentsResponse, error) {
	return s.assignmentService.GetUserAssignments(ctx, req)
}

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	assignmentService := service.NewAssignmentService(db)

	lis, err := net.Listen("tcp", ":9094")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	s := grpc.NewServer()
	assignmentpb.RegisterAssignmentServiceServer(s, &server{assignmentService: assignmentService})

	log.Println("Assignment service listening on port 9094")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
