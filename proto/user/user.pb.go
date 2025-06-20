package user

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type User struct {
	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email      string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	FirstName  string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Department string `protobuf:"bytes,6,opt,name=department,proto3" json:"department,omitempty"`
	Role       string `protobuf:"bytes,7,opt,name=role,proto3" json:"role,omitempty"`
	ManagerId  string `protobuf:"bytes,8,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
	Active     bool   `protobuf:"varint,9,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt  string `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset()         {}
func (x *User) String() string { return "User" }
func (*User) ProtoMessage()    {}

type CreateUserRequest struct {
	Email      string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FirstName  string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Department string `protobuf:"bytes,5,opt,name=department,proto3" json:"department,omitempty"`
	Role       string `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	ManagerId  string `protobuf:"bytes,7,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
}

func (x *CreateUserRequest) Reset()         {}
func (x *CreateUserRequest) String() string { return "CreateUserRequest" }
func (*CreateUserRequest) ProtoMessage()    {}

type CreateUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserResponse) Reset()         {}
func (x *CreateUserResponse) String() string { return "CreateUserResponse" }
func (*CreateUserResponse) ProtoMessage()    {}

type GetUserRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserRequest) Reset()         {}
func (x *GetUserRequest) String() string { return "GetUserRequest" }
func (*GetUserRequest) ProtoMessage()    {}

type GetUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset()         {}
func (x *GetUserResponse) String() string { return "GetUserResponse" }
func (*GetUserResponse) ProtoMessage()    {}

type GetUserByEmailRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetUserByEmailRequest) Reset()         {}
func (x *GetUserByEmailRequest) String() string { return "GetUserByEmailRequest" }
func (*GetUserByEmailRequest) ProtoMessage()    {}

type GetUserByEmailResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByEmailResponse) Reset()         {}
func (x *GetUserByEmailResponse) String() string { return "GetUserByEmailResponse" }
func (*GetUserByEmailResponse) ProtoMessage()    {}

type UpdateUserRequest struct {
	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FirstName  string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Department string `protobuf:"bytes,5,opt,name=department,proto3" json:"department,omitempty"`
	Role       string `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	ManagerId  string `protobuf:"bytes,7,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
	Active     bool   `protobuf:"varint,8,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *UpdateUserRequest) Reset()         {}
func (x *UpdateUserRequest) String() string { return "UpdateUserRequest" }
func (*UpdateUserRequest) ProtoMessage()    {}

type UpdateUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserResponse) Reset()         {}
func (x *UpdateUserResponse) String() string { return "UpdateUserResponse" }
func (*UpdateUserResponse) ProtoMessage()    {}

type DeleteUserRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset()         {}
func (x *DeleteUserRequest) String() string { return "DeleteUserRequest" }
func (*DeleteUserRequest) ProtoMessage()    {}

type DeleteUserResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteUserResponse) Reset()         {}
func (x *DeleteUserResponse) String() string { return "DeleteUserResponse" }
func (*DeleteUserResponse) ProtoMessage()    {}

type ListUsersRequest struct {
	Page       int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize   int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Department string `protobuf:"bytes,3,opt,name=department,proto3" json:"department,omitempty"`
	Role       string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	ActiveOnly bool   `protobuf:"varint,5,opt,name=active_only,json=activeOnly,proto3" json:"active_only,omitempty"`
}

func (x *ListUsersRequest) Reset()         {}
func (x *ListUsersRequest) String() string { return "ListUsersRequest" }
func (*ListUsersRequest) ProtoMessage()    {}

type ListUsersResponse struct {
	Users    []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Total    int32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page     int32   `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32   `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListUsersResponse) Reset()         {}
func (x *ListUsersResponse) String() string { return "ListUsersResponse" }
func (*ListUsersResponse) ProtoMessage()    {}

type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*GetUserByEmailResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetUserByEmail(context.Context, *GetUserByEmailRequest) (*GetUserByEmailResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
}

type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func (UnimplementedUserServiceServer) GetUserByEmail(context.Context, *GetUserByEmailRequest) (*GetUserByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}

func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func (UnimplementedUserServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	return out, nil
}

func (c *userServiceClient) GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*GetUserByEmailResponse, error) {
	out := new(GetUserByEmailResponse)
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	return out, nil
}
