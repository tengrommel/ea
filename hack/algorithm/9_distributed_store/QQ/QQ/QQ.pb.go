// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        (unknown)
// source: FindQQPassword.proto

package QQ

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type QQRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QQnum  string `protobuf:"bytes,1,opt,name=QQnum,json=qQnum,proto3" json:"QQnum,omitempty"`
	QQname string `protobuf:"bytes,2,opt,name=QQname,json=qQname,proto3" json:"QQname,omitempty"`
}

func (x *QQRequest) Reset() {
	*x = QQRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QQ_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QQRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QQRequest) ProtoMessage() {}

func (x *QQRequest) ProtoReflect() protoreflect.Message {
	mi := &file_QQ_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QQRequest.ProtoReflect.Descriptor instead.
func (*QQRequest) Descriptor() ([]byte, []int) {
	return file_QQ_proto_rawDescGZIP(), []int{0}
}

func (x *QQRequest) GetQQnum() string {
	if x != nil {
		return x.QQnum
	}
	return ""
}

func (x *QQRequest) GetQQname() string {
	if x != nil {
		return x.QQname
	}
	return ""
}

type QQResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password   []string `protobuf:"bytes,1,rep,name=password,proto3" json:"password,omitempty"`
	QQpassword string   `protobuf:"bytes,2,opt,name=QQpassword,json=qQpassword,proto3" json:"QQpassword,omitempty"`
}

func (x *QQResponse) Reset() {
	*x = QQResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QQ_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QQResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QQResponse) ProtoMessage() {}

func (x *QQResponse) ProtoReflect() protoreflect.Message {
	mi := &file_QQ_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QQResponse.ProtoReflect.Descriptor instead.
func (*QQResponse) Descriptor() ([]byte, []int) {
	return file_QQ_proto_rawDescGZIP(), []int{1}
}

func (x *QQResponse) GetPassword() []string {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *QQResponse) GetQQpassword() string {
	if x != nil {
		return x.QQpassword
	}
	return ""
}

var File_QQ_proto protoreflect.FileDescriptor

var file_QQ_proto_rawDesc = []byte{
	0x0a, 0x08, 0x51, 0x51, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x51, 0x51, 0x22, 0x39,
	0x0a, 0x09, 0x51, 0x51, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x51,
	0x51, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x51, 0x6e, 0x75,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x51, 0x51, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x71, 0x51, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0a, 0x51, 0x51, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x51, 0x51, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x51, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x32, 0x36, 0x0a, 0x02, 0x51, 0x51, 0x12, 0x30, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x51, 0x51, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x0d, 0x2e, 0x51, 0x51, 0x2e,
	0x51, 0x51, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x51, 0x51, 0x2e, 0x51,
	0x51, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x51,
	0x51, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_QQ_proto_rawDescOnce sync.Once
	file_QQ_proto_rawDescData = file_QQ_proto_rawDesc
)

func file_QQ_proto_rawDescGZIP() []byte {
	file_QQ_proto_rawDescOnce.Do(func() {
		file_QQ_proto_rawDescData = protoimpl.X.CompressGZIP(file_QQ_proto_rawDescData)
	})
	return file_QQ_proto_rawDescData
}

var file_QQ_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_QQ_proto_goTypes = []interface{}{
	(*QQRequest)(nil),  // 0: QQ.QQRequest
	(*QQResponse)(nil), // 1: QQ.QQResponse
}
var file_QQ_proto_depIdxs = []int32{
	0, // 0: QQ.QQ.GetQQPassWord:input_type -> QQ.QQRequest
	1, // 1: QQ.QQ.GetQQPassWord:output_type -> QQ.QQResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_QQ_proto_init() }
func file_QQ_proto_init() {
	if File_QQ_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_QQ_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QQRequest); i {
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
		file_QQ_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QQResponse); i {
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
			RawDescriptor: file_QQ_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_QQ_proto_goTypes,
		DependencyIndexes: file_QQ_proto_depIdxs,
		MessageInfos:      file_QQ_proto_msgTypes,
	}.Build()
	File_QQ_proto = out.File
	file_QQ_proto_rawDesc = nil
	file_QQ_proto_goTypes = nil
	file_QQ_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QQClient is the client API for QQ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QQClient interface {
	GetQQPassWord(ctx context.Context, in *QQRequest, opts ...grpc.CallOption) (*QQResponse, error)
}

type qQClient struct {
	cc grpc.ClientConnInterface
}

func NewQQClient(cc grpc.ClientConnInterface) QQClient {
	return &qQClient{cc}
}

func (c *qQClient) GetQQPassWord(ctx context.Context, in *QQRequest, opts ...grpc.CallOption) (*QQResponse, error) {
	out := new(QQResponse)
	err := c.cc.Invoke(ctx, "/QQ.QQ/GetQQPassWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QQServer is the server API for QQ service.
type QQServer interface {
	GetQQPassWord(context.Context, *QQRequest) (*QQResponse, error)
}

// UnimplementedQQServer can be embedded to have forward compatible implementations.
type UnimplementedQQServer struct {
}

func (*UnimplementedQQServer) GetQQPassWord(context.Context, *QQRequest) (*QQResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQQPassWord not implemented")
}

func RegisterQQServer(s *grpc.Server, srv QQServer) {
	s.RegisterService(&_QQ_serviceDesc, srv)
}

func _QQ_GetQQPassWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QQRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QQServer).GetQQPassWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QQ.QQ/GetQQPassWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QQServer).GetQQPassWord(ctx, req.(*QQRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QQ_serviceDesc = grpc.ServiceDesc{
	ServiceName: "QQ.QQ",
	HandlerType: (*QQServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQQPassWord",
			Handler:    _QQ_GetQQPassWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "FindQQPassword.proto",
}