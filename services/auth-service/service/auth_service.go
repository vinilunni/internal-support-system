package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"internal-support-system/pkg/config"
	"internal-support-system/pkg/jwt"
	authpb "internal-support-system/proto/auth"
	"internal-support-system/services/auth-service/models"

	"gorm.io/gorm"
)

type AuthService struct {
	db         *gorm.DB
	jwtManager jwt.JWTManager
	config     *config.Config
}

func NewAuthService(db *gorm.DB, jwtManager jwt.JWTManager, cfg *config.Config) *AuthService {
	db.AutoMigrate(&models.User{}, &models.RefreshToken{})

	return &AuthService{
		db:         db,
		jwtManager: jwtManager,
		config:     cfg,
	}
}

func (s *AuthService) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	claims, err := s.jwtManager.ValidateToken(req.Token)
	if err != nil {
		return &authpb.ValidateTokenResponse{Valid: false}, nil
	}

	return &authpb.ValidateTokenResponse{
		Valid:  true,
		UserId: claims.UserID,
		Email:  claims.Email,
		Domain: claims.Domain,
		Roles:  claims.Roles,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	var refreshToken models.RefreshToken
	if err := s.db.Where("token = ? AND expires_at > ?", req.RefreshToken, time.Now()).First(&refreshToken).Error; err != nil {
		return nil, errors.New("invalid refresh token")
	}

	var user models.User
	if err := s.db.First(&user, refreshToken.UserID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	accessToken, err := s.jwtManager.GenerateToken(user.ID, user.Email, user.Domain, []string{"user"})
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.createRefreshToken(&user)
	if err != nil {
		return nil, err
	}

	s.db.Delete(&refreshToken)

	return &authpb.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken.Token,
		ExpiresIn:    int64(24 * time.Hour.Seconds()),
	}, nil
}

func (s *AuthService) RevokeToken(ctx context.Context, req *authpb.RevokeTokenRequest) (*authpb.RevokeTokenResponse, error) {
	claims, err := s.jwtManager.ValidateToken(req.Token)
	if err != nil {
		return &authpb.RevokeTokenResponse{Success: false}, nil
	}

	s.db.Where("user_id = ?", claims.UserID).Delete(&models.RefreshToken{})

	return &authpb.RevokeTokenResponse{Success: true}, nil
}

func (s *AuthService) ProcessOAuthUser(email, name, picture string) (*models.User, string, string, error) {
	domain := s.extractDomain(email)
	if !s.isDomainAllowed(domain) {
		return nil, "", "", errors.New("domain not allowed")
	}

	var user models.User
	result := s.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			user = models.User{
				Email:     email,
				Name:      name,
				Picture:   picture,
				Domain:    domain,
				Active:    true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := s.db.Create(&user).Error; err != nil {
				return nil, "", "", err
			}
		} else {
			return nil, "", "", result.Error
		}
	}

	accessToken, err := s.jwtManager.GenerateToken(user.ID, user.Email, user.Domain, []string{"user"})
	if err != nil {
		return nil, "", "", err
	}

	refreshToken, err := s.createRefreshToken(&user)
	if err != nil {
		return nil, "", "", err
	}

	return &user, accessToken, refreshToken.Token, nil
}

func (s *AuthService) GetUserByToken(token string) (*models.User, error) {
	claims, err := s.jwtManager.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := s.db.First(&user, "id = ?", claims.UserID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) createRefreshToken(user *models.User) (*models.RefreshToken, error) {
	refreshToken := &models.RefreshToken{
		UserID:    user.ID,
		Token:     generateRandomToken(),
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		CreatedAt: time.Now(),
	}

	if err := s.db.Create(refreshToken).Error; err != nil {
		return nil, err
	}

	return refreshToken, nil
}

func (s *AuthService) extractDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func (s *AuthService) isDomainAllowed(domain string) bool {
	for _, allowedDomain := range s.config.AllowedDomains {
		if domain == allowedDomain {
			return true
		}
	}
	return false
}

func generateRandomToken() string {
	return "refresh_token_" + time.Now().Format("20060102150405")
}
