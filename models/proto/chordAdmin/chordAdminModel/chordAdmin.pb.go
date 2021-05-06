// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: chordAdmin.proto

package chordAdminModel

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

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chordAdmin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_chordAdmin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_chordAdmin_proto_rawDescGZIP(), []int{0}
}

func (x *Success) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ActiveNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActiveNodesRequest) Reset() {
	*x = ActiveNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chordAdmin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveNodesRequest) ProtoMessage() {}

func (x *ActiveNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chordAdmin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveNodesRequest.ProtoReflect.Descriptor instead.
func (*ActiveNodesRequest) Descriptor() ([]byte, []int) {
	return file_chordAdmin_proto_rawDescGZIP(), []int{1}
}

type ActiveNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *ActiveNodes) Reset() {
	*x = ActiveNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chordAdmin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveNodes) ProtoMessage() {}

func (x *ActiveNodes) ProtoReflect() protoreflect.Message {
	mi := &file_chordAdmin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveNodes.ProtoReflect.Descriptor instead.
func (*ActiveNodes) Descriptor() ([]byte, []int) {
	return file_chordAdmin_proto_rawDescGZIP(), []int{2}
}

func (x *ActiveNodes) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chordAdmin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_chordAdmin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_chordAdmin_proto_rawDescGZIP(), []int{3}
}

func (x *Node) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_chordAdmin_proto protoreflect.FileDescriptor

var file_chordAdmin_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x23, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a,
	0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x16, 0x0a, 0x04, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x64, 0x32, 0x63, 0x0a, 0x0a, 0x43, 0x68, 0x6f, 0x72, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x20, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x05, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x33, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x63, 0x68, 0x6f, 0x72,
	0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chordAdmin_proto_rawDescOnce sync.Once
	file_chordAdmin_proto_rawDescData = file_chordAdmin_proto_rawDesc
)

func file_chordAdmin_proto_rawDescGZIP() []byte {
	file_chordAdmin_proto_rawDescOnce.Do(func() {
		file_chordAdmin_proto_rawDescData = protoimpl.X.CompressGZIP(file_chordAdmin_proto_rawDescData)
	})
	return file_chordAdmin_proto_rawDescData
}

var file_chordAdmin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chordAdmin_proto_goTypes = []interface{}{
	(*Success)(nil),            // 0: Success
	(*ActiveNodesRequest)(nil), // 1: ActiveNodesRequest
	(*ActiveNodes)(nil),        // 2: ActiveNodes
	(*Node)(nil),               // 3: Node
}
var file_chordAdmin_proto_depIdxs = []int32{
	3, // 0: ActiveNodes.nodes:type_name -> Node
	3, // 1: ChordAdmin.AddActiveNode:input_type -> Node
	1, // 2: ChordAdmin.GetActiveNodes:input_type -> ActiveNodesRequest
	0, // 3: ChordAdmin.AddActiveNode:output_type -> Success
	2, // 4: ChordAdmin.GetActiveNodes:output_type -> ActiveNodes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chordAdmin_proto_init() }
func file_chordAdmin_proto_init() {
	if File_chordAdmin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chordAdmin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
		file_chordAdmin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveNodesRequest); i {
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
		file_chordAdmin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveNodes); i {
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
		file_chordAdmin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
			RawDescriptor: file_chordAdmin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chordAdmin_proto_goTypes,
		DependencyIndexes: file_chordAdmin_proto_depIdxs,
		MessageInfos:      file_chordAdmin_proto_msgTypes,
	}.Build()
	File_chordAdmin_proto = out.File
	file_chordAdmin_proto_rawDesc = nil
	file_chordAdmin_proto_goTypes = nil
	file_chordAdmin_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChordAdminClient is the client API for ChordAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChordAdminClient interface {
	AddActiveNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Success, error)
	GetActiveNodes(ctx context.Context, in *ActiveNodesRequest, opts ...grpc.CallOption) (*ActiveNodes, error)
}

type chordAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewChordAdminClient(cc grpc.ClientConnInterface) ChordAdminClient {
	return &chordAdminClient{cc}
}

func (c *chordAdminClient) AddActiveNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/ChordAdmin/AddActiveNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordAdminClient) GetActiveNodes(ctx context.Context, in *ActiveNodesRequest, opts ...grpc.CallOption) (*ActiveNodes, error) {
	out := new(ActiveNodes)
	err := c.cc.Invoke(ctx, "/ChordAdmin/GetActiveNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChordAdminServer is the server API for ChordAdmin service.
type ChordAdminServer interface {
	AddActiveNode(context.Context, *Node) (*Success, error)
	GetActiveNodes(context.Context, *ActiveNodesRequest) (*ActiveNodes, error)
}

// UnimplementedChordAdminServer can be embedded to have forward compatible implementations.
type UnimplementedChordAdminServer struct {
}

func (*UnimplementedChordAdminServer) AddActiveNode(context.Context, *Node) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddActiveNode not implemented")
}
func (*UnimplementedChordAdminServer) GetActiveNodes(context.Context, *ActiveNodesRequest) (*ActiveNodes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveNodes not implemented")
}

func RegisterChordAdminServer(s *grpc.Server, srv ChordAdminServer) {
	s.RegisterService(&_ChordAdmin_serviceDesc, srv)
}

func _ChordAdmin_AddActiveNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordAdminServer).AddActiveNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChordAdmin/AddActiveNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordAdminServer).AddActiveNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordAdmin_GetActiveNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActiveNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordAdminServer).GetActiveNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChordAdmin/GetActiveNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordAdminServer).GetActiveNodes(ctx, req.(*ActiveNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChordAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ChordAdmin",
	HandlerType: (*ChordAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddActiveNode",
			Handler:    _ChordAdmin_AddActiveNode_Handler,
		},
		{
			MethodName: "GetActiveNodes",
			Handler:    _ChordAdmin_GetActiveNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chordAdmin.proto",
}
