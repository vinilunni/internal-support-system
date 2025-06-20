package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"internal-support-system/services/auth-service/service"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHandlers struct {
	authService *service.AuthService
	oauthConfig *oauth2.Config
}

type GoogleUserInfo struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func NewAuthHandlers(authService *service.AuthService) *AuthHandlers {
	oauthConfig := &oauth2.Config{
		ClientID:     "your-google-client-id",
		ClientSecret: "your-google-client-secret",
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     google.Endpoint,
	}

	return &AuthHandlers{
		authService: authService,
		oauthConfig: oauthConfig,
	}
}

func (h *AuthHandlers) GoogleAuth(w http.ResponseWriter, r *http.Request) {
	state := "random-state-string"
	url := h.oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *AuthHandlers) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "No code in request", http.StatusBadRequest)
		return
	}

	token, err := h.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	client := h.oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
		return
	}

	user, accessToken, refreshToken, err := h.authService.ProcessOAuthUser(
		userInfo.Email,
		userInfo.Name,
		userInfo.Picture,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"token_type":    "Bearer",
		"expires_in":    86400,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandlers) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Refresh token functionality to be implemented",
	})
}

func (h *AuthHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	token := h.extractTokenFromHeader(r)
	if token == "" {
		http.Error(w, "No token provided", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{
		"success": true,
	})
}

func (h *AuthHandlers) GetMe(w http.ResponseWriter, r *http.Request) {
	token := h.extractTokenFromHeader(r)
	if token == "" {
		http.Error(w, "No token provided", http.StatusUnauthorized)
		return
	}

	user, err := h.authService.GetUserByToken(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandlers) extractTokenFromHeader(r *http.Request) string {
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
