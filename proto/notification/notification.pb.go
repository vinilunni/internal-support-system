package notification

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type NotificationType int32

const (
	NotificationType_EMAIL  NotificationType = 0
	NotificationType_SLACK  NotificationType = 1
	NotificationType_IN_APP NotificationType = 2
	NotificationType_SMS    NotificationType = 3
)

type NotificationPriority int32

const (
	NotificationPriority_LOW    NotificationPriority = 0
	NotificationPriority_MEDIUM NotificationPriority = 1
	NotificationPriority_HIGH   NotificationPriority = 2
	NotificationPriority_URGENT NotificationPriority = 3
)

type Notification struct {
	Id        string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string               `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title     string               `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Message   string               `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Type      NotificationType     `protobuf:"varint,5,opt,name=type,proto3,enum=notification.NotificationType" json:"type,omitempty"`
	Priority  NotificationPriority `protobuf:"varint,6,opt,name=priority,proto3,enum=notification.NotificationPriority" json:"priority,omitempty"`
	Read      bool                 `protobuf:"varint,7,opt,name=read,proto3" json:"read,omitempty"`
	Metadata  string               `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	CreatedAt string               `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string               `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Notification) Reset()         {}
func (x *Notification) String() string { return "Notification" }
func (*Notification) ProtoMessage()    {}

type SendNotificationRequest struct {
	UserId   string               `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title    string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Message  string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Type     NotificationType     `protobuf:"varint,4,opt,name=type,proto3,enum=notification.NotificationType" json:"type,omitempty"`
	Priority NotificationPriority `protobuf:"varint,5,opt,name=priority,proto3,enum=notification.NotificationPriority" json:"priority,omitempty"`
	Metadata string               `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *SendNotificationRequest) Reset()         {}
func (x *SendNotificationRequest) String() string { return "SendNotificationRequest" }
func (*SendNotificationRequest) ProtoMessage()    {}

type SendNotificationResponse struct {
	Notification *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *SendNotificationResponse) Reset()         {}
func (x *SendNotificationResponse) String() string { return "SendNotificationResponse" }
func (*SendNotificationResponse) ProtoMessage()    {}

type GetNotificationsRequest struct {
	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	ReadOnly bool   `protobuf:"varint,4,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
}

func (x *GetNotificationsRequest) Reset()         {}
func (x *GetNotificationsRequest) String() string { return "GetNotificationsRequest" }
func (*GetNotificationsRequest) ProtoMessage()    {}

type GetNotificationsResponse struct {
	Notifications []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	Total         int32           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page          int32           `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32           `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetNotificationsResponse) Reset()         {}
func (x *GetNotificationsResponse) String() string { return "GetNotificationsResponse" }
func (*GetNotificationsResponse) ProtoMessage()    {}

type MarkAsReadRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MarkAsReadRequest) Reset()         {}
func (x *MarkAsReadRequest) String() string { return "MarkAsReadRequest" }
func (*MarkAsReadRequest) ProtoMessage()    {}

type MarkAsReadResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *MarkAsReadResponse) Reset()         {}
func (x *MarkAsReadResponse) String() string { return "MarkAsReadResponse" }
func (*MarkAsReadResponse) ProtoMessage()    {}

type MarkAllAsReadRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *MarkAllAsReadRequest) Reset()         {}
func (x *MarkAllAsReadRequest) String() string { return "MarkAllAsReadRequest" }
func (*MarkAllAsReadRequest) ProtoMessage()    {}

type MarkAllAsReadResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *MarkAllAsReadResponse) Reset()         {}
func (x *MarkAllAsReadResponse) String() string { return "MarkAllAsReadResponse" }
func (*MarkAllAsReadResponse) ProtoMessage()    {}

type DeleteNotificationRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteNotificationRequest) Reset()         {}
func (x *DeleteNotificationRequest) String() string { return "DeleteNotificationRequest" }
func (*DeleteNotificationRequest) ProtoMessage()    {}

type DeleteNotificationResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteNotificationResponse) Reset()         {}
func (x *DeleteNotificationResponse) String() string { return "DeleteNotificationResponse" }
func (*DeleteNotificationResponse) ProtoMessage()    {}

type SendBulkNotificationRequest struct {
	UserIds  []string             `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	Title    string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Message  string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Type     NotificationType     `protobuf:"varint,4,opt,name=type,proto3,enum=notification.NotificationType" json:"type,omitempty"`
	Priority NotificationPriority `protobuf:"varint,5,opt,name=priority,proto3,enum=notification.NotificationPriority" json:"priority,omitempty"`
	Metadata string               `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *SendBulkNotificationRequest) Reset()         {}
func (x *SendBulkNotificationRequest) String() string { return "SendBulkNotificationRequest" }
func (*SendBulkNotificationRequest) ProtoMessage()    {}

type SendBulkNotificationResponse struct {
	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Count   int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SendBulkNotificationResponse) Reset()         {}
func (x *SendBulkNotificationResponse) String() string { return "SendBulkNotificationResponse" }
func (*SendBulkNotificationResponse) ProtoMessage()    {}

type CreateTemplateRequest struct {
	Name    string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Subject string           `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body    string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Type    NotificationType `protobuf:"varint,4,opt,name=type,proto3,enum=notification.NotificationType" json:"type,omitempty"`
}

func (x *CreateTemplateRequest) Reset()         {}
func (x *CreateTemplateRequest) String() string { return "CreateTemplateRequest" }
func (*CreateTemplateRequest) ProtoMessage()    {}

type CreateTemplateResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateTemplateResponse) Reset()         {}
func (x *CreateTemplateResponse) String() string { return "CreateTemplateResponse" }
func (*CreateTemplateResponse) ProtoMessage()    {}

type UpdateTemplateRequest struct {
	Id      string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Subject string           `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Body    string           `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Type    NotificationType `protobuf:"varint,5,opt,name=type,proto3,enum=notification.NotificationType" json:"type,omitempty"`
}

func (x *UpdateTemplateRequest) Reset()         {}
func (x *UpdateTemplateRequest) String() string { return "UpdateTemplateRequest" }
func (*UpdateTemplateRequest) ProtoMessage()    {}

type UpdateTemplateResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateTemplateResponse) Reset()         {}
func (x *UpdateTemplateResponse) String() string { return "UpdateTemplateResponse" }
func (*UpdateTemplateResponse) ProtoMessage()    {}

type NotificationServiceClient interface {
	SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*SendNotificationResponse, error)
	GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error)
	MarkAsRead(ctx context.Context, in *MarkAsReadRequest, opts ...grpc.CallOption) (*MarkAsReadResponse, error)
	MarkAllAsRead(ctx context.Context, in *MarkAllAsReadRequest, opts ...grpc.CallOption) (*MarkAllAsReadResponse, error)
	DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...grpc.CallOption) (*DeleteNotificationResponse, error)
	SendBulkNotification(ctx context.Context, in *SendBulkNotificationRequest, opts ...grpc.CallOption) (*SendBulkNotificationResponse, error)
	CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error)
	UpdateTemplate(ctx context.Context, in *UpdateTemplateRequest, opts ...grpc.CallOption) (*UpdateTemplateResponse, error)
}

type NotificationServiceServer interface {
	SendNotification(context.Context, *SendNotificationRequest) (*SendNotificationResponse, error)
	GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error)
	MarkAsRead(context.Context, *MarkAsReadRequest) (*MarkAsReadResponse, error)
	MarkAllAsRead(context.Context, *MarkAllAsReadRequest) (*MarkAllAsReadResponse, error)
	DeleteNotification(context.Context, *DeleteNotificationRequest) (*DeleteNotificationResponse, error)
	SendBulkNotification(context.Context, *SendBulkNotificationRequest) (*SendBulkNotificationResponse, error)
	CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error)
	UpdateTemplate(context.Context, *UpdateTemplateRequest) (*UpdateTemplateResponse, error)
}

type UnimplementedNotificationServiceServer struct{}

func (UnimplementedNotificationServiceServer) SendNotification(context.Context, *SendNotificationRequest) (*SendNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotification not implemented")
}

func (UnimplementedNotificationServiceServer) GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}

func (UnimplementedNotificationServiceServer) MarkAsRead(context.Context, *MarkAsReadRequest) (*MarkAsReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsRead not implemented")
}

func (UnimplementedNotificationServiceServer) MarkAllAsRead(context.Context, *MarkAllAsReadRequest) (*MarkAllAsReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAllAsRead not implemented")
}

func (UnimplementedNotificationServiceServer) DeleteNotification(context.Context, *DeleteNotificationRequest) (*DeleteNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotification not implemented")
}

func (UnimplementedNotificationServiceServer) SendBulkNotification(context.Context, *SendBulkNotificationRequest) (*SendBulkNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBulkNotification not implemented")
}

func (UnimplementedNotificationServiceServer) CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}

func (UnimplementedNotificationServiceServer) UpdateTemplate(context.Context, *UpdateTemplateRequest) (*UpdateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplate not implemented")
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func (c *notificationServiceClient) SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*SendNotificationResponse, error) {
	out := new(SendNotificationResponse)
	return out, nil
}

func (c *notificationServiceClient) GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error) {
	out := new(GetNotificationsResponse)
	return out, nil
}

func (c *notificationServiceClient) MarkAsRead(ctx context.Context, in *MarkAsReadRequest, opts ...grpc.CallOption) (*MarkAsReadResponse, error) {
	out := new(MarkAsReadResponse)
	return out, nil
}

func (c *notificationServiceClient) MarkAllAsRead(ctx context.Context, in *MarkAllAsReadRequest, opts ...grpc.CallOption) (*MarkAllAsReadResponse, error) {
	out := new(MarkAllAsReadResponse)
	return out, nil
}

func (c *notificationServiceClient) DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...grpc.CallOption) (*DeleteNotificationResponse, error) {
	out := new(DeleteNotificationResponse)
	return out, nil
}

func (c *notificationServiceClient) SendBulkNotification(ctx context.Context, in *SendBulkNotificationRequest, opts ...grpc.CallOption) (*SendBulkNotificationResponse, error) {
	out := new(SendBulkNotificationResponse)
	return out, nil
}

func (c *notificationServiceClient) CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error) {
	out := new(CreateTemplateResponse)
	return out, nil
}

func (c *notificationServiceClient) UpdateTemplate(ctx context.Context, in *UpdateTemplateRequest, opts ...grpc.CallOption) (*UpdateTemplateResponse, error) {
	out := new(UpdateTemplateResponse)
	return out, nil
}
