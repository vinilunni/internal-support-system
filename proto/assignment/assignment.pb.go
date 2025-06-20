package assignment

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type AssignmentStatus int32

const (
	AssignmentStatus_ACTIVE   AssignmentStatus = 0
	AssignmentStatus_RETURNED AssignmentStatus = 1
	AssignmentStatus_OVERDUE  AssignmentStatus = 2
	AssignmentStatus_LOST     AssignmentStatus = 3
	AssignmentStatus_DAMAGED  AssignmentStatus = 4
)

type Assignment struct {
	Id                 string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AssetId            string           `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	UserId             string           `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AssignedDate       string           `protobuf:"bytes,4,opt,name=assigned_date,json=assignedDate,proto3" json:"assigned_date,omitempty"`
	ReturnDate         string           `protobuf:"bytes,5,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	ExpectedReturnDate string           `protobuf:"bytes,6,opt,name=expected_return_date,json=expectedReturnDate,proto3" json:"expected_return_date,omitempty"`
	ActualReturn       string           `protobuf:"bytes,7,opt,name=actual_return,json=actualReturn,proto3" json:"actual_return,omitempty"`
	ActualReturnDate   string           `protobuf:"bytes,8,opt,name=actual_return_date,json=actualReturnDate,proto3" json:"actual_return_date,omitempty"`
	Status             AssignmentStatus `protobuf:"varint,9,opt,name=status,proto3,enum=assignment.AssignmentStatus" json:"status,omitempty"`
	Notes              string           `protobuf:"bytes,10,opt,name=notes,proto3" json:"notes,omitempty"`
	AssignedBy         string           `protobuf:"bytes,11,opt,name=assigned_by,json=assignedBy,proto3" json:"assigned_by,omitempty"`
	CreatedAt          string           `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          string           `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Assignment) Reset()         {}
func (x *Assignment) String() string { return "Assignment" }
func (*Assignment) ProtoMessage()    {}

type CreateAssignmentRequest struct {
	AssetId    string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ReturnDate string `protobuf:"bytes,3,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	Notes      string `protobuf:"bytes,4,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *CreateAssignmentRequest) Reset()         {}
func (x *CreateAssignmentRequest) String() string { return "CreateAssignmentRequest" }
func (*CreateAssignmentRequest) ProtoMessage()    {}

type CreateAssignmentResponse struct {
	Assignment *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
}

func (x *CreateAssignmentResponse) Reset()         {}
func (x *CreateAssignmentResponse) String() string { return "CreateAssignmentResponse" }
func (*CreateAssignmentResponse) ProtoMessage()    {}

type GetAssignmentRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAssignmentRequest) Reset()         {}
func (x *GetAssignmentRequest) String() string { return "GetAssignmentRequest" }
func (*GetAssignmentRequest) ProtoMessage()    {}

type GetAssignmentResponse struct {
	Assignment *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
}

func (x *GetAssignmentResponse) Reset()         {}
func (x *GetAssignmentResponse) String() string { return "GetAssignmentResponse" }
func (*GetAssignmentResponse) ProtoMessage()    {}

type UpdateAssignmentRequest struct {
	Id           string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReturnDate   string           `protobuf:"bytes,2,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	ActualReturn string           `protobuf:"bytes,3,opt,name=actual_return,json=actualReturn,proto3" json:"actual_return,omitempty"`
	Status       AssignmentStatus `protobuf:"varint,4,opt,name=status,proto3,enum=assignment.AssignmentStatus" json:"status,omitempty"`
	Notes        string           `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *UpdateAssignmentRequest) Reset()         {}
func (x *UpdateAssignmentRequest) String() string { return "UpdateAssignmentRequest" }
func (*UpdateAssignmentRequest) ProtoMessage()    {}

type UpdateAssignmentResponse struct {
	Assignment *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
}

func (x *UpdateAssignmentResponse) Reset()         {}
func (x *UpdateAssignmentResponse) String() string { return "UpdateAssignmentResponse" }
func (*UpdateAssignmentResponse) ProtoMessage()    {}

type DeleteAssignmentRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAssignmentRequest) Reset()         {}
func (x *DeleteAssignmentRequest) String() string { return "DeleteAssignmentRequest" }
func (*DeleteAssignmentRequest) ProtoMessage()    {}

type DeleteAssignmentResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteAssignmentResponse) Reset()         {}
func (x *DeleteAssignmentResponse) String() string { return "DeleteAssignmentResponse" }
func (*DeleteAssignmentResponse) ProtoMessage()    {}

type ListAssignmentsRequest struct {
	Page     int32            `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32            `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	UserId   string           `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AssetId  string           `protobuf:"bytes,4,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Status   AssignmentStatus `protobuf:"varint,5,opt,name=status,proto3,enum=assignment.AssignmentStatus" json:"status,omitempty"`
}

func (x *ListAssignmentsRequest) Reset()         {}
func (x *ListAssignmentsRequest) String() string { return "ListAssignmentsRequest" }
func (*ListAssignmentsRequest) ProtoMessage()    {}

type ListAssignmentsResponse struct {
	Assignments []*Assignment `protobuf:"bytes,1,rep,name=assignments,proto3" json:"assignments,omitempty"`
	Total       int32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page        int32         `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize    int32         `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListAssignmentsResponse) Reset()         {}
func (x *ListAssignmentsResponse) String() string { return "ListAssignmentsResponse" }
func (*ListAssignmentsResponse) ProtoMessage()    {}

type ReturnAssetRequest struct {
	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Notes string `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *ReturnAssetRequest) Reset()         {}
func (x *ReturnAssetRequest) String() string { return "ReturnAssetRequest" }
func (*ReturnAssetRequest) ProtoMessage()    {}

type ReturnAssetResponse struct {
	Assignment *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
}

func (x *ReturnAssetResponse) Reset()         {}
func (x *ReturnAssetResponse) String() string { return "ReturnAssetResponse" }
func (*ReturnAssetResponse) ProtoMessage()    {}

type GetOverdueAssignmentsRequest struct {
	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetOverdueAssignmentsRequest) Reset()         {}
func (x *GetOverdueAssignmentsRequest) String() string { return "GetOverdueAssignmentsRequest" }
func (*GetOverdueAssignmentsRequest) ProtoMessage()    {}

type GetOverdueAssignmentsResponse struct {
	Assignments []*Assignment `protobuf:"bytes,1,rep,name=assignments,proto3" json:"assignments,omitempty"`
	Total       int32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page        int32         `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize    int32         `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetOverdueAssignmentsResponse) Reset()         {}
func (x *GetOverdueAssignmentsResponse) String() string { return "GetOverdueAssignmentsResponse" }
func (*GetOverdueAssignmentsResponse) ProtoMessage()    {}

type AssignAssetRequest struct {
	AssetId            string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	UserId             string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ReturnDate         string `protobuf:"bytes,3,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	ExpectedReturnDate string `protobuf:"bytes,4,opt,name=expected_return_date,json=expectedReturnDate,proto3" json:"expected_return_date,omitempty"`
	AssignedDate       string `protobuf:"bytes,5,opt,name=assigned_date,json=assignedDate,proto3" json:"assigned_date,omitempty"`
	Notes              string `protobuf:"bytes,6,opt,name=notes,proto3" json:"notes,omitempty"`
	AssignedBy         string `protobuf:"bytes,7,opt,name=assigned_by,json=assignedBy,proto3" json:"assigned_by,omitempty"`
}

func (x *AssignAssetRequest) Reset()         {}
func (x *AssignAssetRequest) String() string { return "AssignAssetRequest" }
func (*AssignAssetRequest) ProtoMessage()    {}

type AssignAssetResponse struct {
	Assignment *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
}

func (x *AssignAssetResponse) Reset()         {}
func (x *AssignAssetResponse) String() string { return "AssignAssetResponse" }
func (*AssignAssetResponse) ProtoMessage()    {}

type UnassignAssetRequest struct {
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AssignmentId string `protobuf:"bytes,2,opt,name=assignment_id,json=assignmentId,proto3" json:"assignment_id,omitempty"`
	Notes        string `protobuf:"bytes,3,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *UnassignAssetRequest) Reset()         {}
func (x *UnassignAssetRequest) String() string { return "UnassignAssetRequest" }
func (*UnassignAssetRequest) ProtoMessage()    {}

type UnassignAssetResponse struct {
	Assignment *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
}

func (x *UnassignAssetResponse) Reset()         {}
func (x *UnassignAssetResponse) String() string { return "UnassignAssetResponse" }
func (*UnassignAssetResponse) ProtoMessage()    {}

type GetUserAssignmentsRequest struct {
	UserId   string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page     int32            `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32            `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Status   AssignmentStatus `protobuf:"varint,4,opt,name=status,proto3,enum=assignment.AssignmentStatus" json:"status,omitempty"`
}

func (x *GetUserAssignmentsRequest) Reset()         {}
func (x *GetUserAssignmentsRequest) String() string { return "GetUserAssignmentsRequest" }
func (*GetUserAssignmentsRequest) ProtoMessage()    {}

type GetUserAssignmentsResponse struct {
	Assignments []*Assignment `protobuf:"bytes,1,rep,name=assignments,proto3" json:"assignments,omitempty"`
	Total       int32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page        int32         `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize    int32         `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetUserAssignmentsResponse) Reset()         {}
func (x *GetUserAssignmentsResponse) String() string { return "GetUserAssignmentsResponse" }
func (*GetUserAssignmentsResponse) ProtoMessage()    {}

type AssignmentServiceClient interface {
	CreateAssignment(ctx context.Context, in *CreateAssignmentRequest, opts ...grpc.CallOption) (*CreateAssignmentResponse, error)
	GetAssignment(ctx context.Context, in *GetAssignmentRequest, opts ...grpc.CallOption) (*GetAssignmentResponse, error)
	UpdateAssignment(ctx context.Context, in *UpdateAssignmentRequest, opts ...grpc.CallOption) (*UpdateAssignmentResponse, error)
	DeleteAssignment(ctx context.Context, in *DeleteAssignmentRequest, opts ...grpc.CallOption) (*DeleteAssignmentResponse, error)
	ListAssignments(ctx context.Context, in *ListAssignmentsRequest, opts ...grpc.CallOption) (*ListAssignmentsResponse, error)
	ReturnAsset(ctx context.Context, in *ReturnAssetRequest, opts ...grpc.CallOption) (*ReturnAssetResponse, error)
	GetOverdueAssignments(ctx context.Context, in *GetOverdueAssignmentsRequest, opts ...grpc.CallOption) (*GetOverdueAssignmentsResponse, error)
	AssignAsset(ctx context.Context, in *AssignAssetRequest, opts ...grpc.CallOption) (*AssignAssetResponse, error)
	UnassignAsset(ctx context.Context, in *UnassignAssetRequest, opts ...grpc.CallOption) (*UnassignAssetResponse, error)
	GetUserAssignments(ctx context.Context, in *GetUserAssignmentsRequest, opts ...grpc.CallOption) (*GetUserAssignmentsResponse, error)
}

type AssignmentServiceServer interface {
	CreateAssignment(context.Context, *CreateAssignmentRequest) (*CreateAssignmentResponse, error)
	GetAssignment(context.Context, *GetAssignmentRequest) (*GetAssignmentResponse, error)
	UpdateAssignment(context.Context, *UpdateAssignmentRequest) (*UpdateAssignmentResponse, error)
	DeleteAssignment(context.Context, *DeleteAssignmentRequest) (*DeleteAssignmentResponse, error)
	ListAssignments(context.Context, *ListAssignmentsRequest) (*ListAssignmentsResponse, error)
	ReturnAsset(context.Context, *ReturnAssetRequest) (*ReturnAssetResponse, error)
	GetOverdueAssignments(context.Context, *GetOverdueAssignmentsRequest) (*GetOverdueAssignmentsResponse, error)
	AssignAsset(context.Context, *AssignAssetRequest) (*AssignAssetResponse, error)
	UnassignAsset(context.Context, *UnassignAssetRequest) (*UnassignAssetResponse, error)
	GetUserAssignments(context.Context, *GetUserAssignmentsRequest) (*GetUserAssignmentsResponse, error)
}

type UnimplementedAssignmentServiceServer struct{}

func (UnimplementedAssignmentServiceServer) CreateAssignment(context.Context, *CreateAssignmentRequest) (*CreateAssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAssignment not implemented")
}

func (UnimplementedAssignmentServiceServer) GetAssignment(context.Context, *GetAssignmentRequest) (*GetAssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssignment not implemented")
}

func (UnimplementedAssignmentServiceServer) UpdateAssignment(context.Context, *UpdateAssignmentRequest) (*UpdateAssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAssignment not implemented")
}

func (UnimplementedAssignmentServiceServer) DeleteAssignment(context.Context, *DeleteAssignmentRequest) (*DeleteAssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAssignment not implemented")
}

func (UnimplementedAssignmentServiceServer) ListAssignments(context.Context, *ListAssignmentsRequest) (*ListAssignmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAssignments not implemented")
}

func (UnimplementedAssignmentServiceServer) ReturnAsset(context.Context, *ReturnAssetRequest) (*ReturnAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnAsset not implemented")
}

func (UnimplementedAssignmentServiceServer) GetOverdueAssignments(context.Context, *GetOverdueAssignmentsRequest) (*GetOverdueAssignmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverdueAssignments not implemented")
}

func (UnimplementedAssignmentServiceServer) AssignAsset(context.Context, *AssignAssetRequest) (*AssignAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignAsset not implemented")
}

func (UnimplementedAssignmentServiceServer) UnassignAsset(context.Context, *UnassignAssetRequest) (*UnassignAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnassignAsset not implemented")
}

func (UnimplementedAssignmentServiceServer) GetUserAssignments(context.Context, *GetUserAssignmentsRequest) (*GetUserAssignmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAssignments not implemented")
}

func RegisterAssignmentServiceServer(s grpc.ServiceRegistrar, srv AssignmentServiceServer) {}

func NewAssignmentServiceClient(cc grpc.ClientConnInterface) AssignmentServiceClient {
	return &assignmentServiceClient{cc}
}

type assignmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func (c *assignmentServiceClient) CreateAssignment(ctx context.Context, in *CreateAssignmentRequest, opts ...grpc.CallOption) (*CreateAssignmentResponse, error) {
	out := new(CreateAssignmentResponse)
	return out, nil
}

func (c *assignmentServiceClient) GetAssignment(ctx context.Context, in *GetAssignmentRequest, opts ...grpc.CallOption) (*GetAssignmentResponse, error) {
	out := new(GetAssignmentResponse)
	return out, nil
}

func (c *assignmentServiceClient) UpdateAssignment(ctx context.Context, in *UpdateAssignmentRequest, opts ...grpc.CallOption) (*UpdateAssignmentResponse, error) {
	out := new(UpdateAssignmentResponse)
	return out, nil
}

func (c *assignmentServiceClient) DeleteAssignment(ctx context.Context, in *DeleteAssignmentRequest, opts ...grpc.CallOption) (*DeleteAssignmentResponse, error) {
	out := new(DeleteAssignmentResponse)
	return out, nil
}

func (c *assignmentServiceClient) ListAssignments(ctx context.Context, in *ListAssignmentsRequest, opts ...grpc.CallOption) (*ListAssignmentsResponse, error) {
	out := new(ListAssignmentsResponse)
	return out, nil
}

func (c *assignmentServiceClient) ReturnAsset(ctx context.Context, in *ReturnAssetRequest, opts ...grpc.CallOption) (*ReturnAssetResponse, error) {
	out := new(ReturnAssetResponse)
	return out, nil
}

func (c *assignmentServiceClient) GetOverdueAssignments(ctx context.Context, in *GetOverdueAssignmentsRequest, opts ...grpc.CallOption) (*GetOverdueAssignmentsResponse, error) {
	out := new(GetOverdueAssignmentsResponse)
	return out, nil
}

func (c *assignmentServiceClient) AssignAsset(ctx context.Context, in *AssignAssetRequest, opts ...grpc.CallOption) (*AssignAssetResponse, error) {
	out := new(AssignAssetResponse)
	return out, nil
}

func (c *assignmentServiceClient) UnassignAsset(ctx context.Context, in *UnassignAssetRequest, opts ...grpc.CallOption) (*UnassignAssetResponse, error) {
	out := new(UnassignAssetResponse)
	return out, nil
}

func (c *assignmentServiceClient) GetUserAssignments(ctx context.Context, in *GetUserAssignmentsRequest, opts ...grpc.CallOption) (*GetUserAssignmentsResponse, error) {
	out := new(GetUserAssignmentsResponse)
	return out, nil
}
