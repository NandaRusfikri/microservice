// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/order/order.proto

package order

import (
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

type ModelProtoOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId uint64 `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Amount    int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ModelProtoOrder) Reset() {
	*x = ModelProtoOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelProtoOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelProtoOrder) ProtoMessage() {}

func (x *ModelProtoOrder) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelProtoOrder.ProtoReflect.Descriptor instead.
func (*ModelProtoOrder) Descriptor() ([]byte, []int) {
	return file_proto_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *ModelProtoOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ModelProtoOrder) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ModelProtoOrder) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ModelProtoOrder) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ResponseEntityOrderList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*ModelProtoOrder `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ResponseEntityOrderList) Reset() {
	*x = ResponseEntityOrderList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseEntityOrderList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseEntityOrderList) ProtoMessage() {}

func (x *ResponseEntityOrderList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseEntityOrderList.ProtoReflect.Descriptor instead.
func (*ResponseEntityOrderList) Descriptor() ([]byte, []int) {
	return file_proto_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseEntityOrderList) GetList() []*ModelProtoOrder {
	if x != nil {
		return x.List
	}
	return nil
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId int64 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_order_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateOrderRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

var File_proto_order_order_proto protoreflect.FileDescriptor

var file_proto_order_order_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x22, 0x71, 0x0a, 0x0f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x32, 0x5f, 0x0a, 0x16, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x52, 0x50, 0x43, 0x12, 0x45, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x50, 0x43, 0x12, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x0e, 0x5a, 0x0c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_order_order_proto_rawDescOnce sync.Once
	file_proto_order_order_proto_rawDescData = file_proto_order_order_proto_rawDesc
)

func file_proto_order_order_proto_rawDescGZIP() []byte {
	file_proto_order_order_proto_rawDescOnce.Do(func() {
		file_proto_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_order_order_proto_rawDescData)
	})
	return file_proto_order_order_proto_rawDescData
}

var file_proto_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_order_order_proto_goTypes = []interface{}{
	(*ModelProtoOrder)(nil),         // 0: models.ModelProtoOrder
	(*ResponseEntityOrderList)(nil), // 1: models.ResponseEntityOrderList
	(*CreateOrderRequest)(nil),      // 2: models.CreateOrderRequest
}
var file_proto_order_order_proto_depIdxs = []int32{
	0, // 0: models.ResponseEntityOrderList.list:type_name -> models.ModelProtoOrder
	2, // 1: models.ServiceOrderHandlerRPC.CreateOrderRPC:input_type -> models.CreateOrderRequest
	0, // 2: models.ServiceOrderHandlerRPC.CreateOrderRPC:output_type -> models.ModelProtoOrder
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_order_order_proto_init() }
func file_proto_order_order_proto_init() {
	if File_proto_order_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_order_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelProtoOrder); i {
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
		file_proto_order_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseEntityOrderList); i {
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
		file_proto_order_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
			RawDescriptor: file_proto_order_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_order_proto_goTypes,
		DependencyIndexes: file_proto_order_order_proto_depIdxs,
		MessageInfos:      file_proto_order_order_proto_msgTypes,
	}.Build()
	File_proto_order_order_proto = out.File
	file_proto_order_order_proto_rawDesc = nil
	file_proto_order_order_proto_goTypes = nil
	file_proto_order_order_proto_depIdxs = nil
}