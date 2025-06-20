package service

import (
	"context"
	"time"

	userpb "internal-support-system/proto/user"
	"internal-support-system/services/user-service/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	db.AutoMigrate(&models.User{})
	return &UserService{db: db}
}

func (s *UserService) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user := &models.User{
		Email:      req.Email,
		Name:       req.Name,
		Department: req.Department,
		Role:       req.Role,
		ManagerID:  req.ManagerId,
		Active:     true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: s.modelToProto(user),
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	var user models.User
	if err := s.db.First(&user, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		User: s.modelToProto(&user),
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	var user models.User
	if err := s.db.First(&user, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}

	updates := map[string]interface{}{
		"name":       req.Name,
		"department": req.Department,
		"role":       req.Role,
		"manager_id": req.ManagerId,
		"active":     req.Active,
		"updated_at": time.Now(),
	}

	if err := s.db.Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: s.modelToProto(&user),
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	if err := s.db.Delete(&models.User{}, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}

	return &userpb.DeleteUserResponse{
		Success: true,
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	var users []models.User
	var total int64

	query := s.db.Model(&models.User{})

	if req.Department != "" {
		query = query.Where("department = ?", req.Department)
	}

	if req.ActiveOnly {
		query = query.Where("active = ?", true)
	}

	query.Count(&total)

	offset := int(req.Page * req.PageSize)
	if err := query.Offset(offset).Limit(int(req.PageSize)).Find(&users).Error; err != nil {
		return nil, err
	}

	protoUsers := make([]*userpb.User, len(users))
	for i, user := range users {
		protoUsers[i] = s.modelToProto(&user)
	}

	return &userpb.ListUsersResponse{
		Users:    protoUsers,
		Total:    int32(total),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *UserService) modelToProto(user *models.User) *userpb.User {
	return &userpb.User{
		Id:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		Department: user.Department,
		Role:       user.Role,
		ManagerId:  user.ManagerID,
		CreatedAt:  user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  user.UpdatedAt.Format(time.RFC3339),
		Active:     user.Active,
	}
}
