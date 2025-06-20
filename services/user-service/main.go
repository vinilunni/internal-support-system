package main

import (
	"context"
	"log"
	"net"

	"internal-support-system/pkg/config"
	userpb "internal-support-system/proto/user"
	"internal-support-system/services/user-service/service"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type server struct {
	userpb.UnimplementedUserServiceServer
	userService *service.UserService
}

func (s *server) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	return s.userService.CreateUser(ctx, req)
}

func (s *server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return s.userService.GetUser(ctx, req)
}

func (s *server) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	return s.userService.UpdateUser(ctx, req)
}

func (s *server) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	return s.userService.DeleteUser(ctx, req)
}

func (s *server) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	return s.userService.ListUsers(ctx, req)
}

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	userService := service.NewUserService(db)

	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{userService: userService})

	log.Println("User service listening on port 9091")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
