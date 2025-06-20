package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"internal-support-system/pkg/config"
	"internal-support-system/pkg/jwt"
	authpb "internal-support-system/proto/auth"
	"internal-support-system/services/auth-service/handlers"
	"internal-support-system/services/auth-service/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type server struct {
	authpb.UnimplementedAuthServiceServer
	authService *service.AuthService
}

func (s *server) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	return s.authService.ValidateToken(ctx, req)
}

func (s *server) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	return s.authService.RefreshToken(ctx, req)
}

func (s *server) RevokeToken(ctx context.Context, req *authpb.RevokeTokenRequest) (*authpb.RevokeTokenResponse, error) {
	return s.authService.RevokeToken(ctx, req)
}

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	jwtManager := jwt.NewJWTManager(cfg.JWTSecret, 24*time.Hour)
	authService := service.NewAuthService(db, jwtManager, cfg)

	go startGRPCServer(cfg, authService)

	startHTTPServer(cfg, authService)
}

func startGRPCServer(cfg *config.Config, authService *service.AuthService) {
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &server{authService: authService})

	log.Printf("gRPC server listening on port %s", cfg.GRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC:", err)
	}
}

func startHTTPServer(cfg *config.Config, authService *service.AuthService) {
	r := mux.NewRouter()

	authHandlers := handlers.NewAuthHandlers(authService)

	r.HandleFunc("/auth/google", authHandlers.GoogleAuth).Methods("GET")
	r.HandleFunc("/auth/google/callback", authHandlers.GoogleCallback).Methods("GET")
	r.HandleFunc("/auth/refresh", authHandlers.RefreshToken).Methods("POST")
	r.HandleFunc("/auth/logout", authHandlers.Logout).Methods("POST")
	r.HandleFunc("/auth/me", authHandlers.GetMe).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Printf("HTTP server listening on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, handler); err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}
