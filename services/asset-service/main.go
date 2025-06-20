package main

import (
	"context"
	"log"
	"net"

	"internal-support-system/pkg/config"
	assetpb "internal-support-system/proto/asset"
	"internal-support-system/services/asset-service/service"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type server struct {
	assetpb.UnimplementedAssetServiceServer
	assetService *service.AssetService
}

func (s *server) CreateAsset(ctx context.Context, req *assetpb.CreateAssetRequest) (*assetpb.CreateAssetResponse, error) {
	return s.assetService.CreateAsset(ctx, req)
}

func (s *server) GetAsset(ctx context.Context, req *assetpb.GetAssetRequest) (*assetpb.GetAssetResponse, error) {
	return s.assetService.GetAsset(ctx, req)
}

func (s *server) UpdateAsset(ctx context.Context, req *assetpb.UpdateAssetRequest) (*assetpb.UpdateAssetResponse, error) {
	return s.assetService.UpdateAsset(ctx, req)
}

func (s *server) DeleteAsset(ctx context.Context, req *assetpb.DeleteAssetRequest) (*assetpb.DeleteAssetResponse, error) {
	return s.assetService.DeleteAsset(ctx, req)
}

func (s *server) ListAssets(ctx context.Context, req *assetpb.ListAssetsRequest) (*assetpb.ListAssetsResponse, error) {
	return s.assetService.ListAssets(ctx, req)
}

func (s *server) GetAssetsByUser(ctx context.Context, req *assetpb.GetAssetsByUserRequest) (*assetpb.GetAssetsByUserResponse, error) {
	return s.assetService.GetAssetsByUser(ctx, req)
}

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	assetService := service.NewAssetService(db)

	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	s := grpc.NewServer()
	assetpb.RegisterAssetServiceServer(s, &server{assetService: assetService})

	log.Println("Asset service listening on port 9092")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
