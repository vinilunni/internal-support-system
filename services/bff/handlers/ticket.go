package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type TicketHandler struct {
}

func NewTicketHandler() *TicketHandler {
	return &TicketHandler{}
}

func (h *TicketHandler) ListTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "List tickets endpoint",
		"status":  "placeholder",
	})
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Create ticket endpoint",
		"status":  "placeholder",
	})
}

func (h *TicketHandler) GetTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "Get ticket endpoint",
		"ticket_id": ticketID,
		"status":    "placeholder",
	})
}

func (h *TicketHandler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "Update ticket endpoint",
		"ticket_id": ticketID,
		"status":    "placeholder",
	})
}

func (h *TicketHandler) CloseTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "Close ticket endpoint",
		"ticket_id": ticketID,
		"status":    "placeholder",
	})
}

func (h *TicketHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "Add comment endpoint",
		"ticket_id": ticketID,
		"status":    "placeholder",
	})
}

func (h *TicketHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "Get comments endpoint",
		"ticket_id": ticketID,
		"status":    "placeholder",
	})
}
