package service

import (
	"context"
	"time"

	assignmentpb "internal-support-system/proto/assignment"
	"internal-support-system/services/assignment-service/models"

	"gorm.io/gorm"
)

type AssignmentService struct {
	db *gorm.DB
}

func NewAssignmentService(db *gorm.DB) *AssignmentService {
	db.AutoMigrate(&models.Assignment{})
	return &AssignmentService{db: db}
}

func (s *AssignmentService) AssignAsset(ctx context.Context, req *assignmentpb.AssignAssetRequest) (*assignmentpb.AssignAssetResponse, error) {
	assignment := &models.Assignment{
		AssetID:            req.AssetId,
		UserID:             req.UserId,
		AssignedBy:         req.AssignedBy,
		AssignedDate:       req.AssignedDate,
		ExpectedReturnDate: req.ExpectedReturnDate,
		Status:             int32(assignmentpb.AssignmentStatus_ACTIVE),
		Notes:              req.Notes,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	if err := s.db.Create(assignment).Error; err != nil {
		return nil, err
	}

	return &assignmentpb.AssignAssetResponse{
		Assignment: s.modelToProto(assignment),
	}, nil
}

func (s *AssignmentService) UnassignAsset(ctx context.Context, req *assignmentpb.UnassignAssetRequest) (*assignmentpb.UnassignAssetResponse, error) {
	var assignment models.Assignment
	if err := s.db.First(&assignment, "id = ?", req.AssignmentId).Error; err != nil {
		return nil, err
	}

	updates := map[string]interface{}{
		"status":             int32(assignmentpb.AssignmentStatus_RETURNED),
		"actual_return_date": time.Now().Format(time.RFC3339),
		"notes":              req.Notes,
		"updated_at":         time.Now(),
	}

	if err := s.db.Model(&assignment).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &assignmentpb.UnassignAssetResponse{
		Assignment: s.modelToProto(&assignment),
	}, nil
}

func (s *AssignmentService) GetAssignment(ctx context.Context, req *assignmentpb.GetAssignmentRequest) (*assignmentpb.GetAssignmentResponse, error) {
	var assignment models.Assignment
	if err := s.db.First(&assignment, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}

	return &assignmentpb.GetAssignmentResponse{
		Assignment: s.modelToProto(&assignment),
	}, nil
}

func (s *AssignmentService) ListAssignments(ctx context.Context, req *assignmentpb.ListAssignmentsRequest) (*assignmentpb.ListAssignmentsResponse, error) {
	var assignments []models.Assignment
	var total int64

	query := s.db.Model(&models.Assignment{})

	if req.Status != assignmentpb.AssignmentStatus_ACTIVE {
		query = query.Where("status = ?", int32(req.Status))
	}

	if req.UserId != "" {
		query = query.Where("user_id = ?", req.UserId)
	}

	if req.AssetId != "" {
		query = query.Where("asset_id = ?", req.AssetId)
	}

	query.Count(&total)

	offset := int(req.Page * req.PageSize)
	if err := query.Offset(offset).Limit(int(req.PageSize)).Find(&assignments).Error; err != nil {
		return nil, err
	}

	protoAssignments := make([]*assignmentpb.Assignment, len(assignments))
	for i, assignment := range assignments {
		protoAssignments[i] = s.modelToProto(&assignment)
	}

	return &assignmentpb.ListAssignmentsResponse{
		Assignments: protoAssignments,
		Total:       int32(total),
		Page:        req.Page,
		PageSize:    req.PageSize,
	}, nil
}

func (s *AssignmentService) GetUserAssignments(ctx context.Context, req *assignmentpb.GetUserAssignmentsRequest) (*assignmentpb.GetUserAssignmentsResponse, error) {
	var assignments []models.Assignment
	query := s.db.Where("user_id = ?", req.UserId)

	if req.Status != assignmentpb.AssignmentStatus_ACTIVE {
		query = query.Where("status = ?", int32(req.Status))
	}

	if err := query.Find(&assignments).Error; err != nil {
		return nil, err
	}

	protoAssignments := make([]*assignmentpb.Assignment, len(assignments))
	for i, assignment := range assignments {
		protoAssignments[i] = s.modelToProto(&assignment)
	}

	return &assignmentpb.GetUserAssignmentsResponse{
		Assignments: protoAssignments,
	}, nil
}

func (s *AssignmentService) modelToProto(assignment *models.Assignment) *assignmentpb.Assignment {
	return &assignmentpb.Assignment{
		Id:                 assignment.ID,
		AssetId:            assignment.AssetID,
		UserId:             assignment.UserID,
		AssignedBy:         assignment.AssignedBy,
		AssignedDate:       assignment.AssignedDate,
		ExpectedReturnDate: assignment.ExpectedReturnDate,
		ActualReturnDate:   assignment.ActualReturnDate,
		Status:             assignmentpb.AssignmentStatus(assignment.Status),
		Notes:              assignment.Notes,
		CreatedAt:          assignment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          assignment.UpdatedAt.Format(time.RFC3339),
	}
}
