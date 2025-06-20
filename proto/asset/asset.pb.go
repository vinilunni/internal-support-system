package asset

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Asset Status enumeration
type AssetStatus int32

const (
	AssetStatus_AVAILABLE      AssetStatus = 0
	AssetStatus_ASSIGNED       AssetStatus = 1
	AssetStatus_IN_MAINTENANCE AssetStatus = 2
	AssetStatus_RETIRED        AssetStatus = 3
)

// Asset Type enumeration
type AssetType int32

const (
	AssetType_LAPTOP           AssetType = 0
	AssetType_DESKTOP          AssetType = 1
	AssetType_MONITOR          AssetType = 2
	AssetType_PHONE            AssetType = 3
	AssetType_TABLET           AssetType = 4
	AssetType_PERIPHERAL       AssetType = 5
	AssetType_SOFTWARE_LICENSE AssetType = 6
	AssetType_OTHER            AssetType = 7
)

// Asset message
type Asset struct {
	Id            string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type          AssetType   `protobuf:"varint,3,opt,name=type,proto3,enum=asset.AssetType" json:"type,omitempty"`
	Brand         string      `protobuf:"bytes,4,opt,name=brand,proto3" json:"brand,omitempty"`
	Model         string      `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
	SerialNumber  string      `protobuf:"bytes,6,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	PurchaseDate  string      `protobuf:"bytes,7,opt,name=purchase_date,json=purchaseDate,proto3" json:"purchase_date,omitempty"`
	PurchasePrice float64     `protobuf:"fixed64,8,opt,name=purchase_price,json=purchasePrice,proto3" json:"purchase_price,omitempty"`
	Status        AssetStatus `protobuf:"varint,9,opt,name=status,proto3,enum=asset.AssetStatus" json:"status,omitempty"`
	AssignedTo    string      `protobuf:"bytes,10,opt,name=assigned_to,json=assignedTo,proto3" json:"assigned_to,omitempty"`
	AssignedDate  string      `protobuf:"bytes,11,opt,name=assigned_date,json=assignedDate,proto3" json:"assigned_date,omitempty"`
	Location      string      `protobuf:"bytes,12,opt,name=location,proto3" json:"location,omitempty"`
	Notes         string      `protobuf:"bytes,13,opt,name=notes,proto3" json:"notes,omitempty"`
	CreatedAt     string      `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string      `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Asset) Reset()         {}
func (x *Asset) String() string { return "Asset" }
func (*Asset) ProtoMessage()    {}

// Request/Response messages
type CreateAssetRequest struct {
	Name          string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          AssetType `protobuf:"varint,2,opt,name=type,proto3,enum=asset.AssetType" json:"type,omitempty"`
	Brand         string    `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
	Model         string    `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	SerialNumber  string    `protobuf:"bytes,5,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	PurchaseDate  string    `protobuf:"bytes,6,opt,name=purchase_date,json=purchaseDate,proto3" json:"purchase_date,omitempty"`
	PurchasePrice float64   `protobuf:"fixed64,7,opt,name=purchase_price,json=purchasePrice,proto3" json:"purchase_price,omitempty"`
	Location      string    `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	Notes         string    `protobuf:"bytes,9,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *CreateAssetRequest) Reset()         {}
func (x *CreateAssetRequest) String() string { return "CreateAssetRequest" }
func (*CreateAssetRequest) ProtoMessage()    {}

type CreateAssetResponse struct {
	Asset *Asset `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (x *CreateAssetResponse) Reset()         {}
func (x *CreateAssetResponse) String() string { return "CreateAssetResponse" }
func (*CreateAssetResponse) ProtoMessage()    {}

type GetAssetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAssetRequest) Reset()         {}
func (x *GetAssetRequest) String() string { return "GetAssetRequest" }
func (*GetAssetRequest) ProtoMessage()    {}

type GetAssetResponse struct {
	Asset *Asset `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (x *GetAssetResponse) Reset()         {}
func (x *GetAssetResponse) String() string { return "GetAssetResponse" }
func (*GetAssetResponse) ProtoMessage()    {}

type UpdateAssetRequest struct {
	Id            string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type          AssetType   `protobuf:"varint,3,opt,name=type,proto3,enum=asset.AssetType" json:"type,omitempty"`
	Brand         string      `protobuf:"bytes,4,opt,name=brand,proto3" json:"brand,omitempty"`
	Model         string      `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
	SerialNumber  string      `protobuf:"bytes,6,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	PurchaseDate  string      `protobuf:"bytes,7,opt,name=purchase_date,json=purchaseDate,proto3" json:"purchase_date,omitempty"`
	PurchasePrice float64     `protobuf:"fixed64,8,opt,name=purchase_price,json=purchasePrice,proto3" json:"purchase_price,omitempty"`
	Status        AssetStatus `protobuf:"varint,9,opt,name=status,proto3,enum=asset.AssetStatus" json:"status,omitempty"`
	Location      string      `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
	Notes         string      `protobuf:"bytes,11,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *UpdateAssetRequest) Reset()         {}
func (x *UpdateAssetRequest) String() string { return "UpdateAssetRequest" }
func (*UpdateAssetRequest) ProtoMessage()    {}

type UpdateAssetResponse struct {
	Asset *Asset `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (x *UpdateAssetResponse) Reset()         {}
func (x *UpdateAssetResponse) String() string { return "UpdateAssetResponse" }
func (*UpdateAssetResponse) ProtoMessage()    {}

type DeleteAssetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAssetRequest) Reset()         {}
func (x *DeleteAssetRequest) String() string { return "DeleteAssetRequest" }
func (*DeleteAssetRequest) ProtoMessage()    {}

type DeleteAssetResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteAssetResponse) Reset()         {}
func (x *DeleteAssetResponse) String() string { return "DeleteAssetResponse" }
func (*DeleteAssetResponse) ProtoMessage()    {}

type ListAssetsRequest struct {
	Page     int32       `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32       `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Type     AssetType   `protobuf:"varint,3,opt,name=type,proto3,enum=asset.AssetType" json:"type,omitempty"`
	Status   AssetStatus `protobuf:"varint,4,opt,name=status,proto3,enum=asset.AssetStatus" json:"status,omitempty"`
	Location string      `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *ListAssetsRequest) Reset()         {}
func (x *ListAssetsRequest) String() string { return "ListAssetsRequest" }
func (*ListAssetsRequest) ProtoMessage()    {}

type ListAssetsResponse struct {
	Assets   []*Asset `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	Total    int32    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page     int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListAssetsResponse) Reset()         {}
func (x *ListAssetsResponse) String() string { return "ListAssetsResponse" }
func (*ListAssetsResponse) ProtoMessage()    {}

type GetAssetsByUserRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetAssetsByUserRequest) Reset()         {}
func (x *GetAssetsByUserRequest) String() string { return "GetAssetsByUserRequest" }
func (*GetAssetsByUserRequest) ProtoMessage()    {}

type GetAssetsByUserResponse struct {
	Assets []*Asset `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
}

func (x *GetAssetsByUserResponse) Reset()         {}
func (x *GetAssetsByUserResponse) String() string { return "GetAssetsByUserResponse" }
func (*GetAssetsByUserResponse) ProtoMessage()    {}

// AssetService client interface
type AssetServiceClient interface {
	CreateAsset(ctx context.Context, in *CreateAssetRequest, opts ...grpc.CallOption) (*CreateAssetResponse, error)
	GetAsset(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*GetAssetResponse, error)
	UpdateAsset(ctx context.Context, in *UpdateAssetRequest, opts ...grpc.CallOption) (*UpdateAssetResponse, error)
	DeleteAsset(ctx context.Context, in *DeleteAssetRequest, opts ...grpc.CallOption) (*DeleteAssetResponse, error)
	ListAssets(ctx context.Context, in *ListAssetsRequest, opts ...grpc.CallOption) (*ListAssetsResponse, error)
	GetAssetsByUser(ctx context.Context, in *GetAssetsByUserRequest, opts ...grpc.CallOption) (*GetAssetsByUserResponse, error)
}

// AssetService server interface
type AssetServiceServer interface {
	CreateAsset(context.Context, *CreateAssetRequest) (*CreateAssetResponse, error)
	GetAsset(context.Context, *GetAssetRequest) (*GetAssetResponse, error)
	UpdateAsset(context.Context, *UpdateAssetRequest) (*UpdateAssetResponse, error)
	DeleteAsset(context.Context, *DeleteAssetRequest) (*DeleteAssetResponse, error)
	ListAssets(context.Context, *ListAssetsRequest) (*ListAssetsResponse, error)
	GetAssetsByUser(context.Context, *GetAssetsByUserRequest) (*GetAssetsByUserResponse, error)
}

// Unimplemented server
type UnimplementedAssetServiceServer struct{}

func (UnimplementedAssetServiceServer) CreateAsset(context.Context, *CreateAssetRequest) (*CreateAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAsset not implemented")
}

func (UnimplementedAssetServiceServer) GetAsset(context.Context, *GetAssetRequest) (*GetAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAsset not implemented")
}

func (UnimplementedAssetServiceServer) UpdateAsset(context.Context, *UpdateAssetRequest) (*UpdateAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAsset not implemented")
}

func (UnimplementedAssetServiceServer) DeleteAsset(context.Context, *DeleteAssetRequest) (*DeleteAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAsset not implemented")
}

func (UnimplementedAssetServiceServer) ListAssets(context.Context, *ListAssetsRequest) (*ListAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAssets not implemented")
}

func (UnimplementedAssetServiceServer) GetAssetsByUser(context.Context, *GetAssetsByUserRequest) (*GetAssetsByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetsByUser not implemented")
}

// RegisterAssetServiceServer registers the service
func RegisterAssetServiceServer(s grpc.ServiceRegistrar, srv AssetServiceServer) {
	// Registration implementation would go here
}

// NewAssetServiceClient creates a new client
func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return &assetServiceClient{cc}
}

type assetServiceClient struct {
	cc grpc.ClientConnInterface
}

func (c *assetServiceClient) CreateAsset(ctx context.Context, in *CreateAssetRequest, opts ...grpc.CallOption) (*CreateAssetResponse, error) {
	out := new(CreateAssetResponse)
	return out, nil
}

func (c *assetServiceClient) GetAsset(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*GetAssetResponse, error) {
	out := new(GetAssetResponse)
	return out, nil
}

func (c *assetServiceClient) UpdateAsset(ctx context.Context, in *UpdateAssetRequest, opts ...grpc.CallOption) (*UpdateAssetResponse, error) {
	out := new(UpdateAssetResponse)
	return out, nil
}

func (c *assetServiceClient) DeleteAsset(ctx context.Context, in *DeleteAssetRequest, opts ...grpc.CallOption) (*DeleteAssetResponse, error) {
	out := new(DeleteAssetResponse)
	return out, nil
}

func (c *assetServiceClient) ListAssets(ctx context.Context, in *ListAssetsRequest, opts ...grpc.CallOption) (*ListAssetsResponse, error) {
	out := new(ListAssetsResponse)
	return out, nil
}

func (c *assetServiceClient) GetAssetsByUser(ctx context.Context, in *GetAssetsByUserRequest, opts ...grpc.CallOption) (*GetAssetsByUserResponse, error) {
	out := new(GetAssetsByUserResponse)
	return out, nil
}
