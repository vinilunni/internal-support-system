package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"internal-support-system/pkg/jwt"
)

type AuthMiddleware struct {
	jwtManager jwt.JWTManager
}

func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		jwtManager: jwt.NewJWTManager(jwtSecret, 24*time.Hour),
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := m.extractTokenFromHeader(r)
		if token == "" {
			http.Error(w, "No authorization token provided", http.StatusUnauthorized)
			return
		}

		claims, err := m.jwtManager.ValidateToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "user_email", claims.Email)
		ctx = context.WithValue(ctx, "user_roles", claims.Roles)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *AuthMiddleware) extractTokenFromHeader(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		return ""
	}

	parts := strings.Split(bearerToken, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}
