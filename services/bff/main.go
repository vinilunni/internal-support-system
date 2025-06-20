package main

import (
	"log"
	"net/http"

	"internal-support-system/pkg/config"
	"internal-support-system/services/bff/handlers"
	"internal-support-system/services/bff/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	cfg := config.Load()

	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)

	userHandler := handlers.NewUserHandler()
	assetHandler := handlers.NewAssetHandler()
	ticketHandler := handlers.NewTicketHandler()
	assignmentHandler := handlers.NewAssignmentHandler()

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	api.HandleFunc("/assets", assetHandler.ListAssets).Methods("GET")
	api.HandleFunc("/assets", assetHandler.CreateAsset).Methods("POST")
	api.HandleFunc("/assets/{id}", assetHandler.GetAsset).Methods("GET")
	api.HandleFunc("/assets/{id}", assetHandler.UpdateAsset).Methods("PUT")
	api.HandleFunc("/assets/{id}", assetHandler.DeleteAsset).Methods("DELETE")
	api.HandleFunc("/assets/user/{userId}", assetHandler.GetAssetsByUser).Methods("GET")

	api.HandleFunc("/tickets", ticketHandler.ListTickets).Methods("GET")
	api.HandleFunc("/tickets", ticketHandler.CreateTicket).Methods("POST")
	api.HandleFunc("/tickets/{id}", ticketHandler.GetTicket).Methods("GET")
	api.HandleFunc("/tickets/{id}", ticketHandler.UpdateTicket).Methods("PUT")
	api.HandleFunc("/tickets/{id}/close", ticketHandler.CloseTicket).Methods("POST")
	api.HandleFunc("/tickets/{id}/comments", ticketHandler.AddComment).Methods("POST")
	api.HandleFunc("/tickets/{id}/comments", ticketHandler.GetComments).Methods("GET")

	api.HandleFunc("/assignments", assignmentHandler.ListAssignments).Methods("GET")
	api.HandleFunc("/assignments/assign", assignmentHandler.AssignAsset).Methods("POST")
	api.HandleFunc("/assignments/{id}/unassign", assignmentHandler.UnassignAsset).Methods("POST")
	api.HandleFunc("/assignments/user/{userId}", assignmentHandler.GetUserAssignments).Methods("GET")

	api.Use(authMiddleware.Authenticate)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Printf("BFF server listening on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+"8000", handler); err != nil {
		log.Fatal("Failed to start BFF server:", err)
	}
}
