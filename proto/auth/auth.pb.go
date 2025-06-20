package auth

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type AuthServiceClient interface {
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	RevokeToken(ctx context.Context, in *RevokeTokenRequest, opts ...grpc.CallOption) (*RevokeTokenResponse, error)
}

type AuthServiceServer interface {
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	RevokeToken(context.Context, *RevokeTokenRequest) (*RevokeTokenResponse, error)
}

type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}

func (UnimplementedAuthServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}

func (UnimplementedAuthServiceServer) RevokeToken(context.Context, *RevokeTokenRequest) (*RevokeTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeToken not implemented")
}

type ValidateTokenRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidateTokenRequest) Reset()         {}
func (x *ValidateTokenRequest) String() string { return "ValidateTokenRequest" }
func (*ValidateTokenRequest) ProtoMessage()    {}

type ValidateTokenResponse struct {
	Valid  bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	UserId string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email  string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Domain string   `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	Roles  []string `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *ValidateTokenResponse) Reset()         {}
func (x *ValidateTokenResponse) String() string { return "ValidateTokenResponse" }
func (*ValidateTokenResponse) ProtoMessage()    {}

type RefreshTokenRequest struct {
	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshTokenRequest) Reset()         {}
func (x *RefreshTokenRequest) String() string { return "RefreshTokenRequest" }
func (*RefreshTokenRequest) ProtoMessage()    {}

type RefreshTokenResponse struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresIn    int64  `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *RefreshTokenResponse) Reset()         {}
func (x *RefreshTokenResponse) String() string { return "RefreshTokenResponse" }
func (*RefreshTokenResponse) ProtoMessage()    {}

type RevokeTokenRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RevokeTokenRequest) Reset()         {}
func (x *RevokeTokenRequest) String() string { return "RevokeTokenRequest" }
func (*RevokeTokenRequest) ProtoMessage()    {}

type RevokeTokenResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RevokeTokenResponse) Reset()         {}
func (x *RevokeTokenResponse) String() string { return "RevokeTokenResponse" }
func (*RevokeTokenResponse) ProtoMessage()    {}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func (c *authServiceClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	return out, nil
}

func (c *authServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	return out, nil
}

func (c *authServiceClient) RevokeToken(ctx context.Context, in *RevokeTokenRequest, opts ...grpc.CallOption) (*RevokeTokenResponse, error) {
	out := new(RevokeTokenResponse)
	return out, nil
}
