package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AssetHandler struct {
}

func NewAssetHandler() *AssetHandler {
	return &AssetHandler{}
}

func (h *AssetHandler) ListAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "List assets endpoint",
		"status":  "placeholder",
	})
}

func (h *AssetHandler) CreateAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Create asset endpoint",
		"status":  "placeholder",
	})
}

func (h *AssetHandler) GetAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Get asset endpoint",
		"asset_id": assetID,
		"status":   "placeholder",
	})
}

func (h *AssetHandler) UpdateAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Update asset endpoint",
		"asset_id": assetID,
		"status":   "placeholder",
	})
}

func (h *AssetHandler) DeleteAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Delete asset endpoint",
		"asset_id": assetID,
		"status":   "placeholder",
	})
}

func (h *AssetHandler) GetAssetsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Get assets by user endpoint",
		"user_id": userID,
		"status":  "placeholder",
	})
}
