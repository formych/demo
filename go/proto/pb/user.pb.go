// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: proto3/user.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string            `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone    string            `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Addr     string            `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	NickInfo []*NickInfo       `protobuf:"bytes,5,rep,name=nickInfo,proto3" json:"nickInfo,omitempty"`
	Likes    map[string]string `protobuf:"bytes,6,rep,name=likes,proto3" json:"likes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto3_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_proto3_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_proto3_user_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Info) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Info) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Info) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Info) GetNickInfo() []*NickInfo {
	if x != nil {
		return x.NickInfo
	}
	return nil
}

func (x *Info) GetLikes() map[string]string {
	if x != nil {
		return x.Likes
	}
	return nil
}

type NickInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Time int64  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *NickInfo) Reset() {
	*x = NickInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto3_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NickInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NickInfo) ProtoMessage() {}

func (x *NickInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto3_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NickInfo.ProtoReflect.Descriptor instead.
func (*NickInfo) Descriptor() ([]byte, []int) {
	return file_proto3_user_proto_rawDescGZIP(), []int{1}
}

func (x *NickInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NickInfo) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Info *Info `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto3_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto3_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto3_user_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_proto3_user_proto protoreflect.FileDescriptor

var file_proto3_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xe9, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x12, 0x28, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x69, 0x63, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29,
	0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x62, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x4c, 0x69, 0x6b,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x32, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x07, 0x5a,
	0x05, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto3_user_proto_rawDescOnce sync.Once
	file_proto3_user_proto_rawDescData = file_proto3_user_proto_rawDesc
)

func file_proto3_user_proto_rawDescGZIP() []byte {
	file_proto3_user_proto_rawDescOnce.Do(func() {
		file_proto3_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto3_user_proto_rawDescData)
	})
	return file_proto3_user_proto_rawDescData
}

var file_proto3_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto3_user_proto_goTypes = []interface{}{
	(*Info)(nil),     // 0: pb.Info
	(*NickInfo)(nil), // 1: pb.NickInfo
	(*User)(nil),     // 2: pb.User
	nil,              // 3: pb.Info.LikesEntry
}
var file_proto3_user_proto_depIdxs = []int32{
	1, // 0: pb.Info.nickInfo:type_name -> pb.NickInfo
	3, // 1: pb.Info.likes:type_name -> pb.Info.LikesEntry
	0, // 2: pb.User.info:type_name -> pb.Info
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto3_user_proto_init() }
func file_proto3_user_proto_init() {
	if File_proto3_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto3_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_proto3_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NickInfo); i {
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
		file_proto3_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_proto3_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto3_user_proto_goTypes,
		DependencyIndexes: file_proto3_user_proto_depIdxs,
		MessageInfos:      file_proto3_user_proto_msgTypes,
	}.Build()
	File_proto3_user_proto = out.File
	file_proto3_user_proto_rawDesc = nil
	file_proto3_user_proto_goTypes = nil
	file_proto3_user_proto_depIdxs = nil
}
