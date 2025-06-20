package main

import (
	"context"
	"log"
	"net"

	"internal-support-system/pkg/config"
	ticketpb "internal-support-system/proto/ticket"
	"internal-support-system/services/ticket-service/service"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type server struct {
	ticketpb.UnimplementedTicketServiceServer
	ticketService *service.TicketService
}

func (s *server) CreateTicket(ctx context.Context, req *ticketpb.CreateTicketRequest) (*ticketpb.CreateTicketResponse, error) {
	return s.ticketService.CreateTicket(ctx, req)
}

func (s *server) GetTicket(ctx context.Context, req *ticketpb.GetTicketRequest) (*ticketpb.GetTicketResponse, error) {
	return s.ticketService.GetTicket(ctx, req)
}

func (s *server) UpdateTicket(ctx context.Context, req *ticketpb.UpdateTicketRequest) (*ticketpb.UpdateTicketResponse, error) {
	return s.ticketService.UpdateTicket(ctx, req)
}

func (s *server) CloseTicket(ctx context.Context, req *ticketpb.CloseTicketRequest) (*ticketpb.CloseTicketResponse, error) {
	return s.ticketService.CloseTicket(ctx, req)
}

func (s *server) ListTickets(ctx context.Context, req *ticketpb.ListTicketsRequest) (*ticketpb.ListTicketsResponse, error) {
	return s.ticketService.ListTickets(ctx, req)
}

func (s *server) GetUserTickets(ctx context.Context, req *ticketpb.GetUserTicketsRequest) (*ticketpb.GetUserTicketsResponse, error) {
	return s.ticketService.GetUserTickets(ctx, req)
}

func (s *server) AddComment(ctx context.Context, req *ticketpb.AddCommentRequest) (*ticketpb.AddCommentResponse, error) {
	return s.ticketService.AddComment(ctx, req)
}

func (s *server) GetComments(ctx context.Context, req *ticketpb.GetCommentsRequest) (*ticketpb.GetCommentsResponse, error) {
	return s.ticketService.GetComments(ctx, req)
}

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	ticketService := service.NewTicketService(db)

	lis, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	s := grpc.NewServer()
	ticketpb.RegisterTicketServiceServer(s, &server{ticketService: ticketService})

	log.Println("Ticket service listening on port 9093")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
