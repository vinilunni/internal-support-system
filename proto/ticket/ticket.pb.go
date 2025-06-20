package ticket

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type TicketStatus int32

const (
	TicketStatus_OPEN        TicketStatus = 0
	TicketStatus_IN_PROGRESS TicketStatus = 1
	TicketStatus_PENDING     TicketStatus = 2
	TicketStatus_RESOLVED    TicketStatus = 3
	TicketStatus_CLOSED      TicketStatus = 4
)

type TicketPriority int32

const (
	TicketPriority_LOW    TicketPriority = 0
	TicketPriority_MEDIUM TicketPriority = 1
	TicketPriority_HIGH   TicketPriority = 2
	TicketPriority_URGENT TicketPriority = 3
)

type TicketCategory int32

const (
	TicketCategory_HARDWARE TicketCategory = 0
	TicketCategory_SOFTWARE TicketCategory = 1
	TicketCategory_NETWORK  TicketCategory = 2
	TicketCategory_ACCESS   TicketCategory = 3
	TicketCategory_GENERAL  TicketCategory = 4
)

type Ticket struct {
	Id          string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Category    TicketCategory `protobuf:"varint,4,opt,name=category,proto3,enum=ticket.TicketCategory" json:"category,omitempty"`
	Priority    TicketPriority `protobuf:"varint,5,opt,name=priority,proto3,enum=ticket.TicketPriority" json:"priority,omitempty"`
	Status      TicketStatus   `protobuf:"varint,6,opt,name=status,proto3,enum=ticket.TicketStatus" json:"status,omitempty"`
	RequesterId string         `protobuf:"bytes,7,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	AssigneeId  string         `protobuf:"bytes,8,opt,name=assignee_id,json=assigneeId,proto3" json:"assignee_id,omitempty"`
	CreatedAt   string         `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string         `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ResolvedAt  string         `protobuf:"bytes,11,opt,name=resolved_at,json=resolvedAt,proto3" json:"resolved_at,omitempty"`
}

func (x *Ticket) Reset()         {}
func (x *Ticket) String() string { return "Ticket" }
func (*Ticket) ProtoMessage()    {}

type TicketComment struct {
	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TicketId   string `protobuf:"bytes,2,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	UserId     string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content    string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Internal   bool   `protobuf:"varint,5,opt,name=internal,proto3" json:"internal,omitempty"`
	IsInternal bool   `protobuf:"varint,6,opt,name=is_internal,json=isInternal,proto3" json:"is_internal,omitempty"`
	CreatedAt  string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *TicketComment) Reset()         {}
func (x *TicketComment) String() string { return "TicketComment" }
func (*TicketComment) ProtoMessage()    {}

type CreateTicketRequest struct {
	Title       string         `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Category    TicketCategory `protobuf:"varint,3,opt,name=category,proto3,enum=ticket.TicketCategory" json:"category,omitempty"`
	Priority    TicketPriority `protobuf:"varint,4,opt,name=priority,proto3,enum=ticket.TicketPriority" json:"priority,omitempty"`
	RequesterId string         `protobuf:"bytes,5,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
}

func (x *CreateTicketRequest) Reset()         {}
func (x *CreateTicketRequest) String() string { return "CreateTicketRequest" }
func (*CreateTicketRequest) ProtoMessage()    {}

type CreateTicketResponse struct {
	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *CreateTicketResponse) Reset()         {}
func (x *CreateTicketResponse) String() string { return "CreateTicketResponse" }
func (*CreateTicketResponse) ProtoMessage()    {}

type GetTicketRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTicketRequest) Reset()         {}
func (x *GetTicketRequest) String() string { return "GetTicketRequest" }
func (*GetTicketRequest) ProtoMessage()    {}

type GetTicketResponse struct {
	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *GetTicketResponse) Reset()         {}
func (x *GetTicketResponse) String() string { return "GetTicketResponse" }
func (*GetTicketResponse) ProtoMessage()    {}

type UpdateTicketRequest struct {
	Id          string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Category    TicketCategory `protobuf:"varint,4,opt,name=category,proto3,enum=ticket.TicketCategory" json:"category,omitempty"`
	Priority    TicketPriority `protobuf:"varint,5,opt,name=priority,proto3,enum=ticket.TicketPriority" json:"priority,omitempty"`
	Status      TicketStatus   `protobuf:"varint,6,opt,name=status,proto3,enum=ticket.TicketStatus" json:"status,omitempty"`
	AssigneeId  string         `protobuf:"bytes,7,opt,name=assignee_id,json=assigneeId,proto3" json:"assignee_id,omitempty"`
}

func (x *UpdateTicketRequest) Reset()         {}
func (x *UpdateTicketRequest) String() string { return "UpdateTicketRequest" }
func (*UpdateTicketRequest) ProtoMessage()    {}

type UpdateTicketResponse struct {
	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *UpdateTicketResponse) Reset()         {}
func (x *UpdateTicketResponse) String() string { return "UpdateTicketResponse" }
func (*UpdateTicketResponse) ProtoMessage()    {}

type DeleteTicketRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTicketRequest) Reset()         {}
func (x *DeleteTicketRequest) String() string { return "DeleteTicketRequest" }
func (*DeleteTicketRequest) ProtoMessage()    {}

type DeleteTicketResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteTicketResponse) Reset()         {}
func (x *DeleteTicketResponse) String() string { return "DeleteTicketResponse" }
func (*DeleteTicketResponse) ProtoMessage()    {}

type ListTicketsRequest struct {
	Page        int32          `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize    int32          `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Category    TicketCategory `protobuf:"varint,3,opt,name=category,proto3,enum=ticket.TicketCategory" json:"category,omitempty"`
	Priority    TicketPriority `protobuf:"varint,4,opt,name=priority,proto3,enum=ticket.TicketPriority" json:"priority,omitempty"`
	Status      TicketStatus   `protobuf:"varint,5,opt,name=status,proto3,enum=ticket.TicketStatus" json:"status,omitempty"`
	RequesterId string         `protobuf:"bytes,6,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	AssigneeId  string         `protobuf:"bytes,7,opt,name=assignee_id,json=assigneeId,proto3" json:"assignee_id,omitempty"`
}

func (x *ListTicketsRequest) Reset()         {}
func (x *ListTicketsRequest) String() string { return "ListTicketsRequest" }
func (*ListTicketsRequest) ProtoMessage()    {}

type ListTicketsResponse struct {
	Tickets  []*Ticket `protobuf:"bytes,1,rep,name=tickets,proto3" json:"tickets,omitempty"`
	Total    int32     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page     int32     `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32     `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListTicketsResponse) Reset()         {}
func (x *ListTicketsResponse) String() string { return "ListTicketsResponse" }
func (*ListTicketsResponse) ProtoMessage()    {}

type AddCommentRequest struct {
	TicketId   string `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content    string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	IsInternal bool   `protobuf:"varint,4,opt,name=is_internal,json=isInternal,proto3" json:"is_internal,omitempty"`
}

func (x *AddCommentRequest) Reset()         {}
func (x *AddCommentRequest) String() string { return "AddCommentRequest" }
func (*AddCommentRequest) ProtoMessage()    {}

type AddCommentResponse struct {
	Comment *TicketComment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *AddCommentResponse) Reset()         {}
func (x *AddCommentResponse) String() string { return "AddCommentResponse" }
func (*AddCommentResponse) ProtoMessage()    {}

type GetCommentsRequest struct {
	TicketId        string `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	IncludeInternal bool   `protobuf:"varint,2,opt,name=include_internal,json=includeInternal,proto3" json:"include_internal,omitempty"`
}

func (x *GetCommentsRequest) Reset()         {}
func (x *GetCommentsRequest) String() string { return "GetCommentsRequest" }
func (*GetCommentsRequest) ProtoMessage()    {}

type GetCommentsResponse struct {
	Comments []*TicketComment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *GetCommentsResponse) Reset()         {}
func (x *GetCommentsResponse) String() string { return "GetCommentsResponse" }
func (*GetCommentsResponse) ProtoMessage()    {}

type CloseTicketRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CloseTicketRequest) Reset()         {}
func (x *CloseTicketRequest) String() string { return "CloseTicketRequest" }
func (*CloseTicketRequest) ProtoMessage()    {}

type CloseTicketResponse struct {
	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *CloseTicketResponse) Reset()         {}
func (x *CloseTicketResponse) String() string { return "CloseTicketResponse" }
func (*CloseTicketResponse) ProtoMessage()    {}

type GetUserTicketsRequest struct {
	UserId   string       `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page     int32        `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32        `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Status   TicketStatus `protobuf:"varint,4,opt,name=status,proto3,enum=ticket.TicketStatus" json:"status,omitempty"`
}

func (x *GetUserTicketsRequest) Reset()         {}
func (x *GetUserTicketsRequest) String() string { return "GetUserTicketsRequest" }
func (*GetUserTicketsRequest) ProtoMessage()    {}

type GetUserTicketsResponse struct {
	Tickets  []*Ticket `protobuf:"bytes,1,rep,name=tickets,proto3" json:"tickets,omitempty"`
	Total    int32     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page     int32     `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32     `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetUserTicketsResponse) Reset()         {}
func (x *GetUserTicketsResponse) String() string { return "GetUserTicketsResponse" }
func (*GetUserTicketsResponse) ProtoMessage()    {}

type TicketServiceClient interface {
	CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error)
	GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error)
	UpdateTicket(ctx context.Context, in *UpdateTicketRequest, opts ...grpc.CallOption) (*UpdateTicketResponse, error)
	DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error)
	ListTickets(ctx context.Context, in *ListTicketsRequest, opts ...grpc.CallOption) (*ListTicketsResponse, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	CloseTicket(ctx context.Context, in *CloseTicketRequest, opts ...grpc.CallOption) (*CloseTicketResponse, error)
	GetUserTickets(ctx context.Context, in *GetUserTicketsRequest, opts ...grpc.CallOption) (*GetUserTicketsResponse, error)
}

type TicketServiceServer interface {
	CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error)
	GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error)
	UpdateTicket(context.Context, *UpdateTicketRequest) (*UpdateTicketResponse, error)
	DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error)
	ListTickets(context.Context, *ListTicketsRequest) (*ListTicketsResponse, error)
	AddComment(context.Context, *AddCommentRequest) (*AddCommentResponse, error)
	GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error)
	CloseTicket(context.Context, *CloseTicketRequest) (*CloseTicketResponse, error)
	GetUserTickets(context.Context, *GetUserTicketsRequest) (*GetUserTicketsResponse, error)
}

type UnimplementedTicketServiceServer struct{}

func (UnimplementedTicketServiceServer) CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}

func (UnimplementedTicketServiceServer) GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}

func (UnimplementedTicketServiceServer) UpdateTicket(context.Context, *UpdateTicketRequest) (*UpdateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTicket not implemented")
}

func (UnimplementedTicketServiceServer) DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTicket not implemented")
}

func (UnimplementedTicketServiceServer) ListTickets(context.Context, *ListTicketsRequest) (*ListTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTickets not implemented")
}

func (UnimplementedTicketServiceServer) AddComment(context.Context, *AddCommentRequest) (*AddCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}

func (UnimplementedTicketServiceServer) GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}

func (UnimplementedTicketServiceServer) CloseTicket(context.Context, *CloseTicketRequest) (*CloseTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTicket not implemented")
}

func (UnimplementedTicketServiceServer) GetUserTickets(context.Context, *GetUserTicketsRequest) (*GetUserTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTickets not implemented")
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func (c *ticketServiceClient) CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error) {
	out := new(CreateTicketResponse)
	return out, nil
}

func (c *ticketServiceClient) GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error) {
	out := new(GetTicketResponse)
	return out, nil
}

func (c *ticketServiceClient) UpdateTicket(ctx context.Context, in *UpdateTicketRequest, opts ...grpc.CallOption) (*UpdateTicketResponse, error) {
	out := new(UpdateTicketResponse)
	return out, nil
}

func (c *ticketServiceClient) DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error) {
	out := new(DeleteTicketResponse)
	return out, nil
}

func (c *ticketServiceClient) ListTickets(ctx context.Context, in *ListTicketsRequest, opts ...grpc.CallOption) (*ListTicketsResponse, error) {
	out := new(ListTicketsResponse)
	return out, nil
}

func (c *ticketServiceClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error) {
	out := new(AddCommentResponse)
	return out, nil
}

func (c *ticketServiceClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	out := new(GetCommentsResponse)
	return out, nil
}

func (c *ticketServiceClient) CloseTicket(ctx context.Context, in *CloseTicketRequest, opts ...grpc.CallOption) (*CloseTicketResponse, error) {
	out := new(CloseTicketResponse)
	return out, nil
}

func (c *ticketServiceClient) GetUserTickets(ctx context.Context, in *GetUserTicketsRequest, opts ...grpc.CallOption) (*GetUserTicketsResponse, error) {
	out := new(GetUserTicketsResponse)
	return out, nil
}
