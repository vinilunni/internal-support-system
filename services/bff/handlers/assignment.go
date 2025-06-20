package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AssignmentHandler struct {
}

func NewAssignmentHandler() *AssignmentHandler {
	return &AssignmentHandler{}
}

func (h *AssignmentHandler) ListAssignments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "List assignments endpoint",
		"status":  "placeholder",
	})
}

func (h *AssignmentHandler) AssignAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Assign asset endpoint",
		"status":  "placeholder",
	})
}

func (h *AssignmentHandler) UnassignAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assignmentID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":       "Unassign asset endpoint",
		"assignment_id": assignmentID,
		"status":        "placeholder",
	})
}

func (h *AssignmentHandler) GetUserAssignments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Get user assignments endpoint",
		"user_id": userID,
		"status":  "placeholder",
	})
}
