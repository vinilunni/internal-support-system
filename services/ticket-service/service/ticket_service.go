package service

import (
	"context"
	"time"

	ticketpb "internal-support-system/proto/ticket"
	"internal-support-system/services/ticket-service/models"

	"gorm.io/gorm"
)

type TicketService struct {
	db *gorm.DB
}

func NewTicketService(db *gorm.DB) *TicketService {
	db.AutoMigrate(&models.Ticket{}, &models.TicketComment{})
	return &TicketService{db: db}
}

func (s *TicketService) CreateTicket(ctx context.Context, req *ticketpb.CreateTicketRequest) (*ticketpb.CreateTicketResponse, error) {
	ticket := &models.Ticket{
		Title:       req.Title,
		Description: req.Description,
		Category:    int32(req.Category),
		Priority:    int32(req.Priority),
		Status:      int32(ticketpb.TicketStatus_OPEN),
		RequesterID: req.RequesterId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.db.Create(ticket).Error; err != nil {
		return nil, err
	}

	return &ticketpb.CreateTicketResponse{
		Ticket: s.modelToProto(ticket),
	}, nil
}

func (s *TicketService) GetTicket(ctx context.Context, req *ticketpb.GetTicketRequest) (*ticketpb.GetTicketResponse, error) {
	var ticket models.Ticket
	if err := s.db.First(&ticket, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}

	return &ticketpb.GetTicketResponse{
		Ticket: s.modelToProto(&ticket),
	}, nil
}

func (s *TicketService) UpdateTicket(ctx context.Context, req *ticketpb.UpdateTicketRequest) (*ticketpb.UpdateTicketResponse, error) {
	var ticket models.Ticket
	if err := s.db.First(&ticket, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}

	updates := map[string]interface{}{
		"title":       req.Title,
		"description": req.Description,
		"category":    int32(req.Category),
		"priority":    int32(req.Priority),
		"status":      int32(req.Status),
		"assignee_id": req.AssigneeId,
		"updated_at":  time.Now(),
	}

	if err := s.db.Model(&ticket).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &ticketpb.UpdateTicketResponse{
		Ticket: s.modelToProto(&ticket),
	}, nil
}

func (s *TicketService) CloseTicket(ctx context.Context, req *ticketpb.CloseTicketRequest) (*ticketpb.CloseTicketResponse, error) {
	var ticket models.Ticket
	if err := s.db.First(&ticket, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}

	updates := map[string]interface{}{
		"status":      int32(ticketpb.TicketStatus_CLOSED),
		"resolved_at": time.Now().Format(time.RFC3339),
		"updated_at":  time.Now(),
	}

	if err := s.db.Model(&ticket).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &ticketpb.CloseTicketResponse{
		Ticket: s.modelToProto(&ticket),
	}, nil
}

func (s *TicketService) ListTickets(ctx context.Context, req *ticketpb.ListTicketsRequest) (*ticketpb.ListTicketsResponse, error) {
	var tickets []models.Ticket
	var total int64

	query := s.db.Model(&models.Ticket{})

	if req.Status != ticketpb.TicketStatus_OPEN {
		query = query.Where("status = ?", int32(req.Status))
	}

	if req.Priority != ticketpb.TicketPriority_LOW {
		query = query.Where("priority = ?", int32(req.Priority))
	}

	if req.Category != ticketpb.TicketCategory_HARDWARE {
		query = query.Where("category = ?", int32(req.Category))
	}

	if req.AssigneeId != "" {
		query = query.Where("assignee_id = ?", req.AssigneeId)
	}

	if req.RequesterId != "" {
		query = query.Where("requester_id = ?", req.RequesterId)
	}

	query.Count(&total)

	offset := int(req.Page * req.PageSize)
	if err := query.Offset(offset).Limit(int(req.PageSize)).Find(&tickets).Error; err != nil {
		return nil, err
	}

	protoTickets := make([]*ticketpb.Ticket, len(tickets))
	for i, ticket := range tickets {
		protoTickets[i] = s.modelToProto(&ticket)
	}

	return &ticketpb.ListTicketsResponse{
		Tickets:  protoTickets,
		Total:    int32(total),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *TicketService) GetUserTickets(ctx context.Context, req *ticketpb.GetUserTicketsRequest) (*ticketpb.GetUserTicketsResponse, error) {
	var tickets []models.Ticket
	query := s.db.Where("requester_id = ?", req.UserId)

	if req.Status != ticketpb.TicketStatus_OPEN {
		query = query.Where("status = ?", int32(req.Status))
	}

	if err := query.Find(&tickets).Error; err != nil {
		return nil, err
	}

	protoTickets := make([]*ticketpb.Ticket, len(tickets))
	for i, ticket := range tickets {
		protoTickets[i] = s.modelToProto(&ticket)
	}

	return &ticketpb.GetUserTicketsResponse{
		Tickets: protoTickets,
	}, nil
}

func (s *TicketService) AddComment(ctx context.Context, req *ticketpb.AddCommentRequest) (*ticketpb.AddCommentResponse, error) {
	comment := &models.TicketComment{
		TicketID:   req.TicketId,
		UserID:     req.UserId,
		Content:    req.Content,
		IsInternal: req.IsInternal,
		CreatedAt:  time.Now(),
	}

	if err := s.db.Create(comment).Error; err != nil {
		return nil, err
	}

	return &ticketpb.AddCommentResponse{
		Comment: s.commentToProto(comment),
	}, nil
}

func (s *TicketService) GetComments(ctx context.Context, req *ticketpb.GetCommentsRequest) (*ticketpb.GetCommentsResponse, error) {
	var comments []models.TicketComment
	query := s.db.Where("ticket_id = ?", req.TicketId)

	if !req.IncludeInternal {
		query = query.Where("is_internal = ?", false)
	}

	if err := query.Order("created_at ASC").Find(&comments).Error; err != nil {
		return nil, err
	}

	protoComments := make([]*ticketpb.TicketComment, len(comments))
	for i, comment := range comments {
		protoComments[i] = s.commentToProto(&comment)
	}

	return &ticketpb.GetCommentsResponse{
		Comments: protoComments,
	}, nil
}

func (s *TicketService) modelToProto(ticket *models.Ticket) *ticketpb.Ticket {
	return &ticketpb.Ticket{
		Id:          ticket.ID,
		Title:       ticket.Title,
		Description: ticket.Description,
		Category:    ticketpb.TicketCategory(ticket.Category),
		Priority:    ticketpb.TicketPriority(ticket.Priority),
		Status:      ticketpb.TicketStatus(ticket.Status),
		RequesterId: ticket.RequesterID,
		AssigneeId:  ticket.AssigneeID,
		CreatedAt:   ticket.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   ticket.UpdatedAt.Format(time.RFC3339),
		ResolvedAt:  ticket.ResolvedAt,
	}
}

func (s *TicketService) commentToProto(comment *models.TicketComment) *ticketpb.TicketComment {
	return &ticketpb.TicketComment{
		Id:         comment.ID,
		TicketId:   comment.TicketID,
		UserId:     comment.UserID,
		Content:    comment.Content,
		CreatedAt:  comment.CreatedAt.Format(time.RFC3339),
		IsInternal: comment.IsInternal,
	}
}
