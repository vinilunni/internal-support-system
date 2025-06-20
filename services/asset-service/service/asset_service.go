package service

import (
	"context"
	assetpb "internal-support-system/proto/asset"
	"internal-support-system/services/asset-service/models"
	"time"

	"gorm.io/gorm"
)

type AssetService struct {
	db *gorm.DB
}

func NewAssetService(db *gorm.DB) *AssetService {
	db.AutoMigrate(&models.Asset{})
	return &AssetService{db: db}
}

func (s *AssetService) CreateAsset(ctx context.Context, req *assetpb.CreateAssetRequest) (*assetpb.CreateAssetResponse, error) {
	asset := &models.Asset{
		Name:          req.Name,
		Type:          int32(req.Type),
		Brand:         req.Brand,
		Model:         req.Model,
		SerialNumber:  req.SerialNumber,
		PurchaseDate:  req.PurchaseDate,
		PurchasePrice: req.PurchasePrice,
		Status:        int32(assetpb.AssetStatus_AVAILABLE),
		Location:      req.Location,
		Notes:         req.Notes,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.db.Create(asset).Error; err != nil {
		return nil, err
	}

	return &assetpb.CreateAssetResponse{
		Asset: s.modelToProto(asset),
	}, nil
}

func (s *AssetService) GetAsset(ctx context.Context, req *assetpb.GetAssetRequest) (*assetpb.GetAssetResponse, error) {
	var asset models.Asset
	if err := s.db.First(&asset, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}

	return &assetpb.GetAssetResponse{
		Asset: s.modelToProto(&asset),
	}, nil
}

func (s *AssetService) UpdateAsset(ctx context.Context, req *assetpb.UpdateAssetRequest) (*assetpb.UpdateAssetResponse, error) {
	var asset models.Asset
	if err := s.db.First(&asset, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}

	updates := map[string]interface{}{
		"name":           req.Name,
		"type":           int32(req.Type),
		"brand":          req.Brand,
		"model":          req.Model,
		"serial_number":  req.SerialNumber,
		"purchase_date":  req.PurchaseDate,
		"purchase_price": req.PurchasePrice,
		"status":         int32(req.Status),
		"location":       req.Location,
		"notes":          req.Notes,
		"updated_at":     time.Now(),
	}

	if err := s.db.Model(&asset).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &assetpb.UpdateAssetResponse{
		Asset: s.modelToProto(&asset),
	}, nil
}

func (s *AssetService) DeleteAsset(ctx context.Context, req *assetpb.DeleteAssetRequest) (*assetpb.DeleteAssetResponse, error) {
	if err := s.db.Delete(&models.Asset{}, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}

	return &assetpb.DeleteAssetResponse{
		Success: true,
	}, nil
}

func (s *AssetService) ListAssets(ctx context.Context, req *assetpb.ListAssetsRequest) (*assetpb.ListAssetsResponse, error) {
	var assets []models.Asset
	var total int64

	query := s.db.Model(&models.Asset{})

	if req.Type != assetpb.AssetType_LAPTOP {
		query = query.Where("type = ?", int32(req.Type))
	}

	if req.Status != assetpb.AssetStatus_AVAILABLE {
		query = query.Where("status = ?", int32(req.Status))
	}

	if req.Location != "" {
		query = query.Where("location = ?", req.Location)
	}

	query.Count(&total)

	offset := int(req.Page * req.PageSize)
	if err := query.Offset(offset).Limit(int(req.PageSize)).Find(&assets).Error; err != nil {
		return nil, err
	}

	protoAssets := make([]*assetpb.Asset, len(assets))
	for i, asset := range assets {
		protoAssets[i] = s.modelToProto(&asset)
	}

	return &assetpb.ListAssetsResponse{
		Assets:   protoAssets,
		Total:    int32(total),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *AssetService) GetAssetsByUser(ctx context.Context, req *assetpb.GetAssetsByUserRequest) (*assetpb.GetAssetsByUserResponse, error) {
	var assets []models.Asset
	if err := s.db.Where("assigned_to = ?", req.UserId).Find(&assets).Error; err != nil {
		return nil, err
	}

	protoAssets := make([]*assetpb.Asset, len(assets))
	for i, asset := range assets {
		protoAssets[i] = s.modelToProto(&asset)
	}

	return &assetpb.GetAssetsByUserResponse{
		Assets: protoAssets,
	}, nil
}

func (s *AssetService) modelToProto(asset *models.Asset) *assetpb.Asset {
	return &assetpb.Asset{
		Id:            asset.ID,
		Name:          asset.Name,
		Type:          assetpb.AssetType(asset.Type),
		Brand:         asset.Brand,
		Model:         asset.Model,
		SerialNumber:  asset.SerialNumber,
		PurchaseDate:  asset.PurchaseDate,
		PurchasePrice: asset.PurchasePrice,
		Status:        assetpb.AssetStatus(asset.Status),
		AssignedTo:    asset.AssignedTo,
		AssignedDate:  asset.AssignedDate,
		Location:      asset.Location,
		Notes:         asset.Notes,
		CreatedAt:     asset.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     asset.UpdatedAt.Format(time.RFC3339),
	}
}
