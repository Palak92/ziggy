// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: api/crypto.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetCoinHistoryRequest_SortDirection int32

const (
	GetCoinHistoryRequest_SORT_DIRECTION_UNSPECIFIED GetCoinHistoryRequest_SortDirection = 0
	GetCoinHistoryRequest_DESCENDING                 GetCoinHistoryRequest_SortDirection = 1
	GetCoinHistoryRequest_ASCENDING                  GetCoinHistoryRequest_SortDirection = 2
)

// Enum value maps for GetCoinHistoryRequest_SortDirection.
var (
	GetCoinHistoryRequest_SortDirection_name = map[int32]string{
		0: "SORT_DIRECTION_UNSPECIFIED",
		1: "DESCENDING",
		2: "ASCENDING",
	}
	GetCoinHistoryRequest_SortDirection_value = map[string]int32{
		"SORT_DIRECTION_UNSPECIFIED": 0,
		"DESCENDING":                 1,
		"ASCENDING":                  2,
	}
)

func (x GetCoinHistoryRequest_SortDirection) Enum() *GetCoinHistoryRequest_SortDirection {
	p := new(GetCoinHistoryRequest_SortDirection)
	*p = x
	return p
}

func (x GetCoinHistoryRequest_SortDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetCoinHistoryRequest_SortDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_api_crypto_proto_enumTypes[0].Descriptor()
}

func (GetCoinHistoryRequest_SortDirection) Type() protoreflect.EnumType {
	return &file_api_crypto_proto_enumTypes[0]
}

func (x GetCoinHistoryRequest_SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetCoinHistoryRequest_SortDirection.Descriptor instead.
func (GetCoinHistoryRequest_SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{3, 0}
}

type ListCoinsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinNames []string `protobuf:"bytes,1,rep,name=coin_names,json=coinNames,proto3" json:"coin_names,omitempty"`
}

func (x *ListCoinsRequest) Reset() {
	*x = ListCoinsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCoinsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoinsRequest) ProtoMessage() {}

func (x *ListCoinsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoinsRequest.ProtoReflect.Descriptor instead.
func (*ListCoinsRequest) Descriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{0}
}

func (x *ListCoinsRequest) GetCoinNames() []string {
	if x != nil {
		return x.CoinNames
	}
	return nil
}

type ListCoinsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coins []*Coin `protobuf:"bytes,1,rep,name=coins,proto3" json:"coins,omitempty"`
}

func (x *ListCoinsResponse) Reset() {
	*x = ListCoinsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCoinsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoinsResponse) ProtoMessage() {}

func (x *ListCoinsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoinsResponse.ProtoReflect.Descriptor instead.
func (*ListCoinsResponse) Descriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{1}
}

func (x *ListCoinsResponse) GetCoins() []*Coin {
	if x != nil {
		return x.Coins
	}
	return nil
}

type Coin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol     string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Tracked    bool   `protobuf:"varint,4,opt,name=tracked,proto3" json:"tracked,omitempty"`
	Price      string `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	LastSynced string `protobuf:"bytes,7,opt,name=last_synced,json=lastSynced,proto3" json:"last_synced,omitempty"`
}

func (x *Coin) Reset() {
	*x = Coin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coin.ProtoReflect.Descriptor instead.
func (*Coin) Descriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{2}
}

func (x *Coin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Coin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Coin) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Coin) GetTracked() bool {
	if x != nil {
		return x.Tracked
	}
	return false
}

func (x *Coin) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Coin) GetLastSynced() string {
	if x != nil {
		return x.LastSynced
	}
	return ""
}

type GetCoinHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinId               string                              `protobuf:"bytes,1,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	HistoricalDataPoints int32                               `protobuf:"varint,2,opt,name=historical_data_points,json=historicalDataPoints,proto3" json:"historical_data_points,omitempty"`
	SortDirection        GetCoinHistoryRequest_SortDirection `protobuf:"varint,3,opt,name=sort_direction,json=sortDirection,proto3,enum=crypto.GetCoinHistoryRequest_SortDirection" json:"sort_direction,omitempty"`
}

func (x *GetCoinHistoryRequest) Reset() {
	*x = GetCoinHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinHistoryRequest) ProtoMessage() {}

func (x *GetCoinHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoinHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetCoinHistoryRequest) Descriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{3}
}

func (x *GetCoinHistoryRequest) GetCoinId() string {
	if x != nil {
		return x.CoinId
	}
	return ""
}

func (x *GetCoinHistoryRequest) GetHistoricalDataPoints() int32 {
	if x != nil {
		return x.HistoricalDataPoints
	}
	return 0
}

func (x *GetCoinHistoryRequest) GetSortDirection() GetCoinHistoryRequest_SortDirection {
	if x != nil {
		return x.SortDirection
	}
	return GetCoinHistoryRequest_SORT_DIRECTION_UNSPECIFIED
}

type GetCoinHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataPoints []*DataPoint `protobuf:"bytes,1,rep,name=data_points,json=dataPoints,proto3" json:"data_points,omitempty"`
}

func (x *GetCoinHistoryResponse) Reset() {
	*x = GetCoinHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinHistoryResponse) ProtoMessage() {}

func (x *GetCoinHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoinHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetCoinHistoryResponse) Descriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{4}
}

func (x *GetCoinHistoryResponse) GetDataPoints() []*DataPoint {
	if x != nil {
		return x.DataPoints
	}
	return nil
}

type DataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Price    string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	SyncedOn string `protobuf:"bytes,4,opt,name=synced_on,json=syncedOn,proto3" json:"synced_on,omitempty"`
}

func (x *DataPoint) Reset() {
	*x = DataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPoint) ProtoMessage() {}

func (x *DataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPoint.ProtoReflect.Descriptor instead.
func (*DataPoint) Descriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{5}
}

func (x *DataPoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataPoint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataPoint) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *DataPoint) GetSyncedOn() string {
	if x != nil {
		return x.SyncedOn
	}
	return ""
}

type TrackCoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinId         string `protobuf:"bytes,1,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	EnableTracking bool   `protobuf:"varint,2,opt,name=enable_tracking,json=enableTracking,proto3" json:"enable_tracking,omitempty"`
}

func (x *TrackCoinRequest) Reset() {
	*x = TrackCoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackCoinRequest) ProtoMessage() {}

func (x *TrackCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackCoinRequest.ProtoReflect.Descriptor instead.
func (*TrackCoinRequest) Descriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{6}
}

func (x *TrackCoinRequest) GetCoinId() string {
	if x != nil {
		return x.CoinId
	}
	return ""
}

func (x *TrackCoinRequest) GetEnableTracking() bool {
	if x != nil {
		return x.EnableTracking
	}
	return false
}

type TrackCoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TrackCoinResponse) Reset() {
	*x = TrackCoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackCoinResponse) ProtoMessage() {}

func (x *TrackCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackCoinResponse.ProtoReflect.Descriptor instead.
func (*TrackCoinResponse) Descriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{7}
}

func (x *TrackCoinResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_crypto_proto protoreflect.FileDescriptor

var file_api_crypto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x37, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x05, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x69, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x22, 0x8a, 0x02, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x34, 0x0a, 0x16, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x14, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x52, 0x0a, 0x0e, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x72,
	0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x6f, 0x72, 0x74,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x0d, 0x53, 0x6f, 0x72,
	0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x4f,
	0x52, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45,
	0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53,
	0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x22, 0x4c, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x62, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0x54, 0x0a, 0x10, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x22, 0x2d, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xe3, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x42, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x18, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x1d, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69,
	0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x69, 0x6e, 0x12,
	0x18, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_crypto_proto_rawDescOnce sync.Once
	file_api_crypto_proto_rawDescData = file_api_crypto_proto_rawDesc
)

func file_api_crypto_proto_rawDescGZIP() []byte {
	file_api_crypto_proto_rawDescOnce.Do(func() {
		file_api_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_crypto_proto_rawDescData)
	})
	return file_api_crypto_proto_rawDescData
}

var file_api_crypto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_crypto_proto_goTypes = []interface{}{
	(GetCoinHistoryRequest_SortDirection)(0), // 0: crypto.GetCoinHistoryRequest.SortDirection
	(*ListCoinsRequest)(nil),                 // 1: crypto.ListCoinsRequest
	(*ListCoinsResponse)(nil),                // 2: crypto.ListCoinsResponse
	(*Coin)(nil),                             // 3: crypto.Coin
	(*GetCoinHistoryRequest)(nil),            // 4: crypto.GetCoinHistoryRequest
	(*GetCoinHistoryResponse)(nil),           // 5: crypto.GetCoinHistoryResponse
	(*DataPoint)(nil),                        // 6: crypto.DataPoint
	(*TrackCoinRequest)(nil),                 // 7: crypto.TrackCoinRequest
	(*TrackCoinResponse)(nil),                // 8: crypto.TrackCoinResponse
}
var file_api_crypto_proto_depIdxs = []int32{
	3, // 0: crypto.ListCoinsResponse.coins:type_name -> crypto.Coin
	0, // 1: crypto.GetCoinHistoryRequest.sort_direction:type_name -> crypto.GetCoinHistoryRequest.SortDirection
	6, // 2: crypto.GetCoinHistoryResponse.data_points:type_name -> crypto.DataPoint
	1, // 3: crypto.Crypto.ListCoins:input_type -> crypto.ListCoinsRequest
	4, // 4: crypto.Crypto.GetCoinHistory:input_type -> crypto.GetCoinHistoryRequest
	7, // 5: crypto.Crypto.TrackCoin:input_type -> crypto.TrackCoinRequest
	2, // 6: crypto.Crypto.ListCoins:output_type -> crypto.ListCoinsResponse
	5, // 7: crypto.Crypto.GetCoinHistory:output_type -> crypto.GetCoinHistoryResponse
	8, // 8: crypto.Crypto.TrackCoin:output_type -> crypto.TrackCoinResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_crypto_proto_init() }
func file_api_crypto_proto_init() {
	if File_api_crypto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_crypto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCoinsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCoinsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoinHistoryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoinHistoryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackCoinRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackCoinResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_crypto_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_crypto_proto_goTypes,
		DependencyIndexes: file_api_crypto_proto_depIdxs,
		EnumInfos:         file_api_crypto_proto_enumTypes,
		MessageInfos:      file_api_crypto_proto_msgTypes,
	}.Build()
	File_api_crypto_proto = out.File
	file_api_crypto_proto_rawDesc = nil
	file_api_crypto_proto_goTypes = nil
	file_api_crypto_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CryptoClient is the client API for Crypto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CryptoClient interface {
	ListCoins(ctx context.Context, in *ListCoinsRequest, opts ...grpc.CallOption) (*ListCoinsResponse, error)
	GetCoinHistory(ctx context.Context, in *GetCoinHistoryRequest, opts ...grpc.CallOption) (*GetCoinHistoryResponse, error)
	TrackCoin(ctx context.Context, in *TrackCoinRequest, opts ...grpc.CallOption) (*TrackCoinResponse, error)
}

type cryptoClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoClient(cc grpc.ClientConnInterface) CryptoClient {
	return &cryptoClient{cc}
}

func (c *cryptoClient) ListCoins(ctx context.Context, in *ListCoinsRequest, opts ...grpc.CallOption) (*ListCoinsResponse, error) {
	out := new(ListCoinsResponse)
	err := c.cc.Invoke(ctx, "/crypto.Crypto/ListCoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoClient) GetCoinHistory(ctx context.Context, in *GetCoinHistoryRequest, opts ...grpc.CallOption) (*GetCoinHistoryResponse, error) {
	out := new(GetCoinHistoryResponse)
	err := c.cc.Invoke(ctx, "/crypto.Crypto/GetCoinHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoClient) TrackCoin(ctx context.Context, in *TrackCoinRequest, opts ...grpc.CallOption) (*TrackCoinResponse, error) {
	out := new(TrackCoinResponse)
	err := c.cc.Invoke(ctx, "/crypto.Crypto/TrackCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServer is the server API for Crypto service.
type CryptoServer interface {
	ListCoins(context.Context, *ListCoinsRequest) (*ListCoinsResponse, error)
	GetCoinHistory(context.Context, *GetCoinHistoryRequest) (*GetCoinHistoryResponse, error)
	TrackCoin(context.Context, *TrackCoinRequest) (*TrackCoinResponse, error)
}

// UnimplementedCryptoServer can be embedded to have forward compatible implementations.
type UnimplementedCryptoServer struct {
}

func (*UnimplementedCryptoServer) ListCoins(context.Context, *ListCoinsRequest) (*ListCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCoins not implemented")
}
func (*UnimplementedCryptoServer) GetCoinHistory(context.Context, *GetCoinHistoryRequest) (*GetCoinHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinHistory not implemented")
}
func (*UnimplementedCryptoServer) TrackCoin(context.Context, *TrackCoinRequest) (*TrackCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackCoin not implemented")
}

func RegisterCryptoServer(s *grpc.Server, srv CryptoServer) {
	s.RegisterService(&_Crypto_serviceDesc, srv)
}

func _Crypto_ListCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServer).ListCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.Crypto/ListCoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServer).ListCoins(ctx, req.(*ListCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crypto_GetCoinHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServer).GetCoinHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.Crypto/GetCoinHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServer).GetCoinHistory(ctx, req.(*GetCoinHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crypto_TrackCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServer).TrackCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.Crypto/TrackCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServer).TrackCoin(ctx, req.(*TrackCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crypto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crypto.Crypto",
	HandlerType: (*CryptoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCoins",
			Handler:    _Crypto_ListCoins_Handler,
		},
		{
			MethodName: "GetCoinHistory",
			Handler:    _Crypto_GetCoinHistory_Handler,
		},
		{
			MethodName: "TrackCoin",
			Handler:    _Crypto_TrackCoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/crypto.proto",
}
