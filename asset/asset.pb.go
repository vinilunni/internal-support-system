// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: asset/asset.proto

package asset

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AssetStatus int32

const (
	AssetStatus_AVAILABLE      AssetStatus = 0
	AssetStatus_ASSIGNED       AssetStatus = 1
	AssetStatus_IN_MAINTENANCE AssetStatus = 2
	AssetStatus_RETIRED        AssetStatus = 3
)

// Enum value maps for AssetStatus.
var (
	AssetStatus_name = map[int32]string{
		0: "AVAILABLE",
		1: "ASSIGNED",
		2: "IN_MAINTENANCE",
		3: "RETIRED",
	}
	AssetStatus_value = map[string]int32{
		"AVAILABLE":      0,
		"ASSIGNED":       1,
		"IN_MAINTENANCE": 2,
		"RETIRED":        3,
	}
)

func (x AssetStatus) Enum() *AssetStatus {
	p := new(AssetStatus)
	*p = x
	return p
}

func (x AssetStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_asset_asset_proto_enumTypes[0].Descriptor()
}

func (AssetStatus) Type() protoreflect.EnumType {
	return &file_asset_asset_proto_enumTypes[0]
}

func (x AssetStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetStatus.Descriptor instead.
func (AssetStatus) EnumDescriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{0}
}

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

// Enum value maps for AssetType.
var (
	AssetType_name = map[int32]string{
		0: "LAPTOP",
		1: "DESKTOP",
		2: "MONITOR",
		3: "PHONE",
		4: "TABLET",
		5: "PERIPHERAL",
		6: "SOFTWARE_LICENSE",
		7: "OTHER",
	}
	AssetType_value = map[string]int32{
		"LAPTOP":           0,
		"DESKTOP":          1,
		"MONITOR":          2,
		"PHONE":            3,
		"TABLET":           4,
		"PERIPHERAL":       5,
		"SOFTWARE_LICENSE": 6,
		"OTHER":            7,
	}
)

func (x AssetType) Enum() *AssetType {
	p := new(AssetType)
	*p = x
	return p
}

func (x AssetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetType) Descriptor() protoreflect.EnumDescriptor {
	return file_asset_asset_proto_enumTypes[1].Descriptor()
}

func (AssetType) Type() protoreflect.EnumType {
	return &file_asset_asset_proto_enumTypes[1]
}

func (x AssetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetType.Descriptor instead.
func (AssetType) EnumDescriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{1}
}

type Asset struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type          AssetType              `protobuf:"varint,3,opt,name=type,proto3,enum=asset.AssetType" json:"type,omitempty"`
	Brand         string                 `protobuf:"bytes,4,opt,name=brand,proto3" json:"brand,omitempty"`
	Model         string                 `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
	SerialNumber  string                 `protobuf:"bytes,6,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	PurchaseDate  string                 `protobuf:"bytes,7,opt,name=purchase_date,json=purchaseDate,proto3" json:"purchase_date,omitempty"`
	PurchasePrice float64                `protobuf:"fixed64,8,opt,name=purchase_price,json=purchasePrice,proto3" json:"purchase_price,omitempty"`
	Status        AssetStatus            `protobuf:"varint,9,opt,name=status,proto3,enum=asset.AssetStatus" json:"status,omitempty"`
	AssignedTo    string                 `protobuf:"bytes,10,opt,name=assigned_to,json=assignedTo,proto3" json:"assigned_to,omitempty"`
	AssignedDate  string                 `protobuf:"bytes,11,opt,name=assigned_date,json=assignedDate,proto3" json:"assigned_date,omitempty"`
	Location      string                 `protobuf:"bytes,12,opt,name=location,proto3" json:"location,omitempty"`
	Notes         string                 `protobuf:"bytes,13,opt,name=notes,proto3" json:"notes,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Asset) Reset() {
	*x = Asset{}
	mi := &file_asset_asset_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{0}
}

func (x *Asset) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Asset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Asset) GetType() AssetType {
	if x != nil {
		return x.Type
	}
	return AssetType_LAPTOP
}

func (x *Asset) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Asset) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Asset) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *Asset) GetPurchaseDate() string {
	if x != nil {
		return x.PurchaseDate
	}
	return ""
}

func (x *Asset) GetPurchasePrice() float64 {
	if x != nil {
		return x.PurchasePrice
	}
	return 0
}

func (x *Asset) GetStatus() AssetStatus {
	if x != nil {
		return x.Status
	}
	return AssetStatus_AVAILABLE
}

func (x *Asset) GetAssignedTo() string {
	if x != nil {
		return x.AssignedTo
	}
	return ""
}

func (x *Asset) GetAssignedDate() string {
	if x != nil {
		return x.AssignedDate
	}
	return ""
}

func (x *Asset) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Asset) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Asset) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Asset) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateAssetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          AssetType              `protobuf:"varint,2,opt,name=type,proto3,enum=asset.AssetType" json:"type,omitempty"`
	Brand         string                 `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
	Model         string                 `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	SerialNumber  string                 `protobuf:"bytes,5,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	PurchaseDate  string                 `protobuf:"bytes,6,opt,name=purchase_date,json=purchaseDate,proto3" json:"purchase_date,omitempty"`
	PurchasePrice float64                `protobuf:"fixed64,7,opt,name=purchase_price,json=purchasePrice,proto3" json:"purchase_price,omitempty"`
	Location      string                 `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	Notes         string                 `protobuf:"bytes,9,opt,name=notes,proto3" json:"notes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAssetRequest) Reset() {
	*x = CreateAssetRequest{}
	mi := &file_asset_asset_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssetRequest) ProtoMessage() {}

func (x *CreateAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssetRequest.ProtoReflect.Descriptor instead.
func (*CreateAssetRequest) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAssetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAssetRequest) GetType() AssetType {
	if x != nil {
		return x.Type
	}
	return AssetType_LAPTOP
}

func (x *CreateAssetRequest) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CreateAssetRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CreateAssetRequest) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *CreateAssetRequest) GetPurchaseDate() string {
	if x != nil {
		return x.PurchaseDate
	}
	return ""
}

func (x *CreateAssetRequest) GetPurchasePrice() float64 {
	if x != nil {
		return x.PurchasePrice
	}
	return 0
}

func (x *CreateAssetRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateAssetRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type CreateAssetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Asset         *Asset                 `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAssetResponse) Reset() {
	*x = CreateAssetResponse{}
	mi := &file_asset_asset_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAssetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssetResponse) ProtoMessage() {}

func (x *CreateAssetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssetResponse.ProtoReflect.Descriptor instead.
func (*CreateAssetResponse) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAssetResponse) GetAsset() *Asset {
	if x != nil {
		return x.Asset
	}
	return nil
}

type GetAssetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssetRequest) Reset() {
	*x = GetAssetRequest{}
	mi := &file_asset_asset_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetRequest) ProtoMessage() {}

func (x *GetAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetRequest.ProtoReflect.Descriptor instead.
func (*GetAssetRequest) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{3}
}

func (x *GetAssetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAssetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Asset         *Asset                 `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssetResponse) Reset() {
	*x = GetAssetResponse{}
	mi := &file_asset_asset_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetResponse) ProtoMessage() {}

func (x *GetAssetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetResponse.ProtoReflect.Descriptor instead.
func (*GetAssetResponse) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{4}
}

func (x *GetAssetResponse) GetAsset() *Asset {
	if x != nil {
		return x.Asset
	}
	return nil
}

type UpdateAssetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type          AssetType              `protobuf:"varint,3,opt,name=type,proto3,enum=asset.AssetType" json:"type,omitempty"`
	Brand         string                 `protobuf:"bytes,4,opt,name=brand,proto3" json:"brand,omitempty"`
	Model         string                 `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
	SerialNumber  string                 `protobuf:"bytes,6,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	PurchaseDate  string                 `protobuf:"bytes,7,opt,name=purchase_date,json=purchaseDate,proto3" json:"purchase_date,omitempty"`
	PurchasePrice float64                `protobuf:"fixed64,8,opt,name=purchase_price,json=purchasePrice,proto3" json:"purchase_price,omitempty"`
	Status        AssetStatus            `protobuf:"varint,9,opt,name=status,proto3,enum=asset.AssetStatus" json:"status,omitempty"`
	Location      string                 `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
	Notes         string                 `protobuf:"bytes,11,opt,name=notes,proto3" json:"notes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAssetRequest) Reset() {
	*x = UpdateAssetRequest{}
	mi := &file_asset_asset_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssetRequest) ProtoMessage() {}

func (x *UpdateAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAssetRequest.ProtoReflect.Descriptor instead.
func (*UpdateAssetRequest) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAssetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAssetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateAssetRequest) GetType() AssetType {
	if x != nil {
		return x.Type
	}
	return AssetType_LAPTOP
}

func (x *UpdateAssetRequest) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *UpdateAssetRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *UpdateAssetRequest) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *UpdateAssetRequest) GetPurchaseDate() string {
	if x != nil {
		return x.PurchaseDate
	}
	return ""
}

func (x *UpdateAssetRequest) GetPurchasePrice() float64 {
	if x != nil {
		return x.PurchasePrice
	}
	return 0
}

func (x *UpdateAssetRequest) GetStatus() AssetStatus {
	if x != nil {
		return x.Status
	}
	return AssetStatus_AVAILABLE
}

func (x *UpdateAssetRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UpdateAssetRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type UpdateAssetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Asset         *Asset                 `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAssetResponse) Reset() {
	*x = UpdateAssetResponse{}
	mi := &file_asset_asset_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAssetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssetResponse) ProtoMessage() {}

func (x *UpdateAssetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAssetResponse.ProtoReflect.Descriptor instead.
func (*UpdateAssetResponse) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAssetResponse) GetAsset() *Asset {
	if x != nil {
		return x.Asset
	}
	return nil
}

type DeleteAssetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAssetRequest) Reset() {
	*x = DeleteAssetRequest{}
	mi := &file_asset_asset_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAssetRequest) ProtoMessage() {}

func (x *DeleteAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAssetRequest.ProtoReflect.Descriptor instead.
func (*DeleteAssetRequest) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAssetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAssetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAssetResponse) Reset() {
	*x = DeleteAssetResponse{}
	mi := &file_asset_asset_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAssetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAssetResponse) ProtoMessage() {}

func (x *DeleteAssetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAssetResponse.ProtoReflect.Descriptor instead.
func (*DeleteAssetResponse) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAssetResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListAssetsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Type          AssetType              `protobuf:"varint,3,opt,name=type,proto3,enum=asset.AssetType" json:"type,omitempty"`
	Status        AssetStatus            `protobuf:"varint,4,opt,name=status,proto3,enum=asset.AssetStatus" json:"status,omitempty"`
	Location      string                 `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAssetsRequest) Reset() {
	*x = ListAssetsRequest{}
	mi := &file_asset_asset_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsRequest) ProtoMessage() {}

func (x *ListAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsRequest.ProtoReflect.Descriptor instead.
func (*ListAssetsRequest) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{9}
}

func (x *ListAssetsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListAssetsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAssetsRequest) GetType() AssetType {
	if x != nil {
		return x.Type
	}
	return AssetType_LAPTOP
}

func (x *ListAssetsRequest) GetStatus() AssetStatus {
	if x != nil {
		return x.Status
	}
	return AssetStatus_AVAILABLE
}

func (x *ListAssetsRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type ListAssetsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Assets        []*Asset               `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page          int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAssetsResponse) Reset() {
	*x = ListAssetsResponse{}
	mi := &file_asset_asset_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsResponse) ProtoMessage() {}

func (x *ListAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsResponse.ProtoReflect.Descriptor instead.
func (*ListAssetsResponse) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{10}
}

func (x *ListAssetsResponse) GetAssets() []*Asset {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *ListAssetsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListAssetsResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListAssetsResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetAssetsByUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssetsByUserRequest) Reset() {
	*x = GetAssetsByUserRequest{}
	mi := &file_asset_asset_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssetsByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetsByUserRequest) ProtoMessage() {}

func (x *GetAssetsByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetsByUserRequest.ProtoReflect.Descriptor instead.
func (*GetAssetsByUserRequest) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{11}
}

func (x *GetAssetsByUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetAssetsByUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Assets        []*Asset               `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssetsByUserResponse) Reset() {
	*x = GetAssetsByUserResponse{}
	mi := &file_asset_asset_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssetsByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetsByUserResponse) ProtoMessage() {}

func (x *GetAssetsByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asset_asset_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetsByUserResponse.ProtoReflect.Descriptor instead.
func (*GetAssetsByUserResponse) Descriptor() ([]byte, []int) {
	return file_asset_asset_proto_rawDescGZIP(), []int{12}
}

func (x *GetAssetsByUserResponse) GetAssets() []*Asset {
	if x != nil {
		return x.Assets
	}
	return nil
}

var File_asset_asset_proto protoreflect.FileDescriptor

const file_asset_asset_proto_rawDesc = "" +
	"\n" +
	"\x11asset/asset.proto\x12\x05asset\"\xd0\x03\n" +
	"\x05Asset\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12$\n" +
	"\x04type\x18\x03 \x01(\x0e2\x10.asset.AssetTypeR\x04type\x12\x14\n" +
	"\x05brand\x18\x04 \x01(\tR\x05brand\x12\x14\n" +
	"\x05model\x18\x05 \x01(\tR\x05model\x12#\n" +
	"\rserial_number\x18\x06 \x01(\tR\fserialNumber\x12#\n" +
	"\rpurchase_date\x18\a \x01(\tR\fpurchaseDate\x12%\n" +
	"\x0epurchase_price\x18\b \x01(\x01R\rpurchasePrice\x12*\n" +
	"\x06status\x18\t \x01(\x0e2\x12.asset.AssetStatusR\x06status\x12\x1f\n" +
	"\vassigned_to\x18\n" +
	" \x01(\tR\n" +
	"assignedTo\x12#\n" +
	"\rassigned_date\x18\v \x01(\tR\fassignedDate\x12\x1a\n" +
	"\blocation\x18\f \x01(\tR\blocation\x12\x14\n" +
	"\x05notes\x18\r \x01(\tR\x05notes\x12\x1d\n" +
	"\n" +
	"created_at\x18\x0e \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x0f \x01(\tR\tupdatedAt\"\x9d\x02\n" +
	"\x12CreateAssetRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12$\n" +
	"\x04type\x18\x02 \x01(\x0e2\x10.asset.AssetTypeR\x04type\x12\x14\n" +
	"\x05brand\x18\x03 \x01(\tR\x05brand\x12\x14\n" +
	"\x05model\x18\x04 \x01(\tR\x05model\x12#\n" +
	"\rserial_number\x18\x05 \x01(\tR\fserialNumber\x12#\n" +
	"\rpurchase_date\x18\x06 \x01(\tR\fpurchaseDate\x12%\n" +
	"\x0epurchase_price\x18\a \x01(\x01R\rpurchasePrice\x12\x1a\n" +
	"\blocation\x18\b \x01(\tR\blocation\x12\x14\n" +
	"\x05notes\x18\t \x01(\tR\x05notes\"9\n" +
	"\x13CreateAssetResponse\x12\"\n" +
	"\x05asset\x18\x01 \x01(\v2\f.asset.AssetR\x05asset\"!\n" +
	"\x0fGetAssetRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"6\n" +
	"\x10GetAssetResponse\x12\"\n" +
	"\x05asset\x18\x01 \x01(\v2\f.asset.AssetR\x05asset\"\xd9\x02\n" +
	"\x12UpdateAssetRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12$\n" +
	"\x04type\x18\x03 \x01(\x0e2\x10.asset.AssetTypeR\x04type\x12\x14\n" +
	"\x05brand\x18\x04 \x01(\tR\x05brand\x12\x14\n" +
	"\x05model\x18\x05 \x01(\tR\x05model\x12#\n" +
	"\rserial_number\x18\x06 \x01(\tR\fserialNumber\x12#\n" +
	"\rpurchase_date\x18\a \x01(\tR\fpurchaseDate\x12%\n" +
	"\x0epurchase_price\x18\b \x01(\x01R\rpurchasePrice\x12*\n" +
	"\x06status\x18\t \x01(\x0e2\x12.asset.AssetStatusR\x06status\x12\x1a\n" +
	"\blocation\x18\n" +
	" \x01(\tR\blocation\x12\x14\n" +
	"\x05notes\x18\v \x01(\tR\x05notes\"9\n" +
	"\x13UpdateAssetResponse\x12\"\n" +
	"\x05asset\x18\x01 \x01(\v2\f.asset.AssetR\x05asset\"$\n" +
	"\x12DeleteAssetRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"/\n" +
	"\x13DeleteAssetResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\xb2\x01\n" +
	"\x11ListAssetsRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x05R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x05R\bpageSize\x12$\n" +
	"\x04type\x18\x03 \x01(\x0e2\x10.asset.AssetTypeR\x04type\x12*\n" +
	"\x06status\x18\x04 \x01(\x0e2\x12.asset.AssetStatusR\x06status\x12\x1a\n" +
	"\blocation\x18\x05 \x01(\tR\blocation\"\x81\x01\n" +
	"\x12ListAssetsResponse\x12$\n" +
	"\x06assets\x18\x01 \x03(\v2\f.asset.AssetR\x06assets\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x05R\x05total\x12\x12\n" +
	"\x04page\x18\x03 \x01(\x05R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x04 \x01(\x05R\bpageSize\"1\n" +
	"\x16GetAssetsByUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"?\n" +
	"\x17GetAssetsByUserResponse\x12$\n" +
	"\x06assets\x18\x01 \x03(\v2\f.asset.AssetR\x06assets*K\n" +
	"\vAssetStatus\x12\r\n" +
	"\tAVAILABLE\x10\x00\x12\f\n" +
	"\bASSIGNED\x10\x01\x12\x12\n" +
	"\x0eIN_MAINTENANCE\x10\x02\x12\v\n" +
	"\aRETIRED\x10\x03*y\n" +
	"\tAssetType\x12\n" +
	"\n" +
	"\x06LAPTOP\x10\x00\x12\v\n" +
	"\aDESKTOP\x10\x01\x12\v\n" +
	"\aMONITOR\x10\x02\x12\t\n" +
	"\x05PHONE\x10\x03\x12\n" +
	"\n" +
	"\x06TABLET\x10\x04\x12\x0e\n" +
	"\n" +
	"PERIPHERAL\x10\x05\x12\x14\n" +
	"\x10SOFTWARE_LICENSE\x10\x06\x12\t\n" +
	"\x05OTHER\x10\a2\xb2\x03\n" +
	"\fAssetService\x12D\n" +
	"\vCreateAsset\x12\x19.asset.CreateAssetRequest\x1a\x1a.asset.CreateAssetResponse\x12;\n" +
	"\bGetAsset\x12\x16.asset.GetAssetRequest\x1a\x17.asset.GetAssetResponse\x12D\n" +
	"\vUpdateAsset\x12\x19.asset.UpdateAssetRequest\x1a\x1a.asset.UpdateAssetResponse\x12D\n" +
	"\vDeleteAsset\x12\x19.asset.DeleteAssetRequest\x1a\x1a.asset.DeleteAssetResponse\x12A\n" +
	"\n" +
	"ListAssets\x12\x18.asset.ListAssetsRequest\x1a\x19.asset.ListAssetsResponse\x12P\n" +
	"\x0fGetAssetsByUser\x12\x1d.asset.GetAssetsByUserRequest\x1a\x1e.asset.GetAssetsByUserResponseB%Z#internal-support-system/proto/assetb\x06proto3"

var (
	file_asset_asset_proto_rawDescOnce sync.Once
	file_asset_asset_proto_rawDescData []byte
)

func file_asset_asset_proto_rawDescGZIP() []byte {
	file_asset_asset_proto_rawDescOnce.Do(func() {
		file_asset_asset_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_asset_asset_proto_rawDesc), len(file_asset_asset_proto_rawDesc)))
	})
	return file_asset_asset_proto_rawDescData
}

var file_asset_asset_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_asset_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_asset_asset_proto_goTypes = []any{
	(AssetStatus)(0),                // 0: asset.AssetStatus
	(AssetType)(0),                  // 1: asset.AssetType
	(*Asset)(nil),                   // 2: asset.Asset
	(*CreateAssetRequest)(nil),      // 3: asset.CreateAssetRequest
	(*CreateAssetResponse)(nil),     // 4: asset.CreateAssetResponse
	(*GetAssetRequest)(nil),         // 5: asset.GetAssetRequest
	(*GetAssetResponse)(nil),        // 6: asset.GetAssetResponse
	(*UpdateAssetRequest)(nil),      // 7: asset.UpdateAssetRequest
	(*UpdateAssetResponse)(nil),     // 8: asset.UpdateAssetResponse
	(*DeleteAssetRequest)(nil),      // 9: asset.DeleteAssetRequest
	(*DeleteAssetResponse)(nil),     // 10: asset.DeleteAssetResponse
	(*ListAssetsRequest)(nil),       // 11: asset.ListAssetsRequest
	(*ListAssetsResponse)(nil),      // 12: asset.ListAssetsResponse
	(*GetAssetsByUserRequest)(nil),  // 13: asset.GetAssetsByUserRequest
	(*GetAssetsByUserResponse)(nil), // 14: asset.GetAssetsByUserResponse
}
var file_asset_asset_proto_depIdxs = []int32{
	1,  // 0: asset.Asset.type:type_name -> asset.AssetType
	0,  // 1: asset.Asset.status:type_name -> asset.AssetStatus
	1,  // 2: asset.CreateAssetRequest.type:type_name -> asset.AssetType
	2,  // 3: asset.CreateAssetResponse.asset:type_name -> asset.Asset
	2,  // 4: asset.GetAssetResponse.asset:type_name -> asset.Asset
	1,  // 5: asset.UpdateAssetRequest.type:type_name -> asset.AssetType
	0,  // 6: asset.UpdateAssetRequest.status:type_name -> asset.AssetStatus
	2,  // 7: asset.UpdateAssetResponse.asset:type_name -> asset.Asset
	1,  // 8: asset.ListAssetsRequest.type:type_name -> asset.AssetType
	0,  // 9: asset.ListAssetsRequest.status:type_name -> asset.AssetStatus
	2,  // 10: asset.ListAssetsResponse.assets:type_name -> asset.Asset
	2,  // 11: asset.GetAssetsByUserResponse.assets:type_name -> asset.Asset
	3,  // 12: asset.AssetService.CreateAsset:input_type -> asset.CreateAssetRequest
	5,  // 13: asset.AssetService.GetAsset:input_type -> asset.GetAssetRequest
	7,  // 14: asset.AssetService.UpdateAsset:input_type -> asset.UpdateAssetRequest
	9,  // 15: asset.AssetService.DeleteAsset:input_type -> asset.DeleteAssetRequest
	11, // 16: asset.AssetService.ListAssets:input_type -> asset.ListAssetsRequest
	13, // 17: asset.AssetService.GetAssetsByUser:input_type -> asset.GetAssetsByUserRequest
	4,  // 18: asset.AssetService.CreateAsset:output_type -> asset.CreateAssetResponse
	6,  // 19: asset.AssetService.GetAsset:output_type -> asset.GetAssetResponse
	8,  // 20: asset.AssetService.UpdateAsset:output_type -> asset.UpdateAssetResponse
	10, // 21: asset.AssetService.DeleteAsset:output_type -> asset.DeleteAssetResponse
	12, // 22: asset.AssetService.ListAssets:output_type -> asset.ListAssetsResponse
	14, // 23: asset.AssetService.GetAssetsByUser:output_type -> asset.GetAssetsByUserResponse
	18, // [18:24] is the sub-list for method output_type
	12, // [12:18] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_asset_asset_proto_init() }
func file_asset_asset_proto_init() {
	if File_asset_asset_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_asset_asset_proto_rawDesc), len(file_asset_asset_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_asset_asset_proto_goTypes,
		DependencyIndexes: file_asset_asset_proto_depIdxs,
		EnumInfos:         file_asset_asset_proto_enumTypes,
		MessageInfos:      file_asset_asset_proto_msgTypes,
	}.Build()
	File_asset_asset_proto = out.File
	file_asset_asset_proto_goTypes = nil
	file_asset_asset_proto_depIdxs = nil
}
