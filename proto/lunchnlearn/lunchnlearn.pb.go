// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto/lunchnlearn/lunchnlearn.proto

package lunchnlearn

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

type Lunch_State int32

const (
	Lunch_STATE_UNSPECIFIED Lunch_State = 0
	Lunch_ORDERED           Lunch_State = 1
	Lunch_DELIVERED         Lunch_State = 2
	Lunch_CANCELED          Lunch_State = 3
)

// Enum value maps for Lunch_State.
var (
	Lunch_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ORDERED",
		2: "DELIVERED",
		3: "CANCELED",
	}
	Lunch_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ORDERED":           1,
		"DELIVERED":         2,
		"CANCELED":          3,
	}
)

func (x Lunch_State) Enum() *Lunch_State {
	p := new(Lunch_State)
	*p = x
	return p
}

func (x Lunch_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Lunch_State) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_lunchnlearn_lunchnlearn_proto_enumTypes[0].Descriptor()
}

func (Lunch_State) Type() protoreflect.EnumType {
	return &file_proto_lunchnlearn_lunchnlearn_proto_enumTypes[0]
}

func (x Lunch_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Lunch_State.Descriptor instead.
func (Lunch_State) EnumDescriptor() ([]byte, []int) {
	return file_proto_lunchnlearn_lunchnlearn_proto_rawDescGZIP(), []int{1, 0}
}

type LunchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// order - plain text
	Order string `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *LunchRequest) Reset() {
	*x = LunchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lunchnlearn_lunchnlearn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LunchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LunchRequest) ProtoMessage() {}

func (x *LunchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lunchnlearn_lunchnlearn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LunchRequest.ProtoReflect.Descriptor instead.
func (*LunchRequest) Descriptor() ([]byte, []int) {
	return file_proto_lunchnlearn_lunchnlearn_proto_rawDescGZIP(), []int{0}
}

func (x *LunchRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type Lunch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State Lunch_State `protobuf:"varint,1,opt,name=state,proto3,enum=lunchnlearn.Lunch_State" json:"state,omitempty"`
	Name  string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Lunch) Reset() {
	*x = Lunch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lunchnlearn_lunchnlearn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lunch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lunch) ProtoMessage() {}

func (x *Lunch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lunchnlearn_lunchnlearn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lunch.ProtoReflect.Descriptor instead.
func (*Lunch) Descriptor() ([]byte, []int) {
	return file_proto_lunchnlearn_lunchnlearn_proto_rawDescGZIP(), []int{1}
}

func (x *Lunch) GetState() Lunch_State {
	if x != nil {
		return x.State
	}
	return Lunch_STATE_UNSPECIFIED
}

func (x *Lunch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LunchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lunch *Lunch `protobuf:"bytes,1,opt,name=lunch,proto3" json:"lunch,omitempty"`
}

func (x *LunchResponse) Reset() {
	*x = LunchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lunchnlearn_lunchnlearn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LunchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LunchResponse) ProtoMessage() {}

func (x *LunchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lunchnlearn_lunchnlearn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LunchResponse.ProtoReflect.Descriptor instead.
func (*LunchResponse) Descriptor() ([]byte, []int) {
	return file_proto_lunchnlearn_lunchnlearn_proto_rawDescGZIP(), []int{2}
}

func (x *LunchResponse) GetLunch() *Lunch {
	if x != nil {
		return x.Lunch
	}
	return nil
}

var File_proto_lunchnlearn_lunchnlearn_proto protoreflect.FileDescriptor

var file_proto_lunchnlearn_lunchnlearn_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x6e, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x2f, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x6e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x6e, 0x6c, 0x65, 0x61,
	0x72, 0x6e, 0x22, 0x24, 0x0a, 0x0c, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x95, 0x01, 0x0a, 0x05, 0x4c, 0x75, 0x6e,
	0x63, 0x68, 0x12, 0x2e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x6e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2e,
	0x4c, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x03,
	0x22, 0x39, 0x0a, 0x0d, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x05, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x6e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2e, 0x4c,
	0x75, 0x6e, 0x63, 0x68, 0x52, 0x05, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x32, 0x59, 0x0a, 0x12, 0x4c,
	0x75, 0x6e, 0x63, 0x68, 0x6e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x43, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x19, 0x2e, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x6e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2e, 0x4c, 0x75,
	0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x75, 0x6e,
	0x63, 0x68, 0x6e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x75, 0x62, 0x79, 0x63, 0x6f, 0x6e, 0x76, 0x69, 0x63, 0x74,
	0x2f, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x6e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x6e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_lunchnlearn_lunchnlearn_proto_rawDescOnce sync.Once
	file_proto_lunchnlearn_lunchnlearn_proto_rawDescData = file_proto_lunchnlearn_lunchnlearn_proto_rawDesc
)

func file_proto_lunchnlearn_lunchnlearn_proto_rawDescGZIP() []byte {
	file_proto_lunchnlearn_lunchnlearn_proto_rawDescOnce.Do(func() {
		file_proto_lunchnlearn_lunchnlearn_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_lunchnlearn_lunchnlearn_proto_rawDescData)
	})
	return file_proto_lunchnlearn_lunchnlearn_proto_rawDescData
}

var file_proto_lunchnlearn_lunchnlearn_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_lunchnlearn_lunchnlearn_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_lunchnlearn_lunchnlearn_proto_goTypes = []interface{}{
	(Lunch_State)(0),      // 0: lunchnlearn.Lunch.State
	(*LunchRequest)(nil),  // 1: lunchnlearn.LunchRequest
	(*Lunch)(nil),         // 2: lunchnlearn.Lunch
	(*LunchResponse)(nil), // 3: lunchnlearn.LunchResponse
}
var file_proto_lunchnlearn_lunchnlearn_proto_depIdxs = []int32{
	0, // 0: lunchnlearn.Lunch.state:type_name -> lunchnlearn.Lunch.State
	2, // 1: lunchnlearn.LunchResponse.lunch:type_name -> lunchnlearn.Lunch
	1, // 2: lunchnlearn.LunchnlearnService.PlaceOrder:input_type -> lunchnlearn.LunchRequest
	3, // 3: lunchnlearn.LunchnlearnService.PlaceOrder:output_type -> lunchnlearn.LunchResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_lunchnlearn_lunchnlearn_proto_init() }
func file_proto_lunchnlearn_lunchnlearn_proto_init() {
	if File_proto_lunchnlearn_lunchnlearn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_lunchnlearn_lunchnlearn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LunchRequest); i {
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
		file_proto_lunchnlearn_lunchnlearn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lunch); i {
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
		file_proto_lunchnlearn_lunchnlearn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LunchResponse); i {
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
			RawDescriptor: file_proto_lunchnlearn_lunchnlearn_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_lunchnlearn_lunchnlearn_proto_goTypes,
		DependencyIndexes: file_proto_lunchnlearn_lunchnlearn_proto_depIdxs,
		EnumInfos:         file_proto_lunchnlearn_lunchnlearn_proto_enumTypes,
		MessageInfos:      file_proto_lunchnlearn_lunchnlearn_proto_msgTypes,
	}.Build()
	File_proto_lunchnlearn_lunchnlearn_proto = out.File
	file_proto_lunchnlearn_lunchnlearn_proto_rawDesc = nil
	file_proto_lunchnlearn_lunchnlearn_proto_goTypes = nil
	file_proto_lunchnlearn_lunchnlearn_proto_depIdxs = nil
}