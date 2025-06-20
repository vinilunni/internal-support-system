package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

type Claims struct {
	UserID string   `json:"user_id"`
	Email  string   `json:"email"`
	Domain string   `json:"domain"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

type JWTManager interface {
	GenerateToken(userID, email, domain string, roles []string) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
	RefreshToken(tokenString string) (string, error)
}

type jwtManager struct {
	secretKey     []byte
	tokenDuration time.Duration
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) JWTManager {
	return &jwtManager{
		secretKey:     []byte(secretKey),
		tokenDuration: tokenDuration,
	}
}

func (manager *jwtManager) GenerateToken(userID, email, domain string, roles []string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		Domain: domain,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(manager.secretKey)
}

func (manager *jwtManager) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, ErrInvalidToken
			}
			return manager.secretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func (manager *jwtManager) RefreshToken(tokenString string) (string, error) {
	claims, err := manager.ValidateToken(tokenString)
	if err != nil {
		if !errors.Is(err, ErrExpiredToken) {
			return "", err
		}
	}

	return manager.GenerateToken(claims.UserID, claims.Email, claims.Domain, claims.Roles)
}
