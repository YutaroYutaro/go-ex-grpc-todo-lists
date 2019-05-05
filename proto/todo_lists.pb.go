// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo_lists.proto

package todolists

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Todo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Assignee             string   `protobuf:"bytes,2,opt,name=assignee,proto3" json:"assignee,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_31d64610e8c524ba, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Todo) GetAssignee() string {
	if m != nil {
		return m.Assignee
	}
	return ""
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type GetTodoRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoRequest) Reset()         { *m = GetTodoRequest{} }
func (m *GetTodoRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodoRequest) ProtoMessage()    {}
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31d64610e8c524ba, []int{1}
}

func (m *GetTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoRequest.Unmarshal(m, b)
}
func (m *GetTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoRequest.Marshal(b, m, deterministic)
}
func (m *GetTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoRequest.Merge(m, src)
}
func (m *GetTodoRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodoRequest.Size(m)
}
func (m *GetTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoRequest proto.InternalMessageInfo

func (m *GetTodoRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Todo)(nil), "todolists.Todo")
	proto.RegisterType((*GetTodoRequest)(nil), "todolists.GetTodoRequest")
}

func init() { proto.RegisterFile("todo_lists.proto", fileDescriptor_31d64610e8c524ba) }

var fileDescriptor_31d64610e8c524ba = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xc9, 0x4f, 0xc9,
	0x8f, 0xcf, 0xc9, 0x2c, 0x2e, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x89,
	0x80, 0x05, 0xa4, 0xa4, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0x12, 0x49, 0xa5, 0x69,
	0xfa, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x10, 0x75, 0x4a, 0x31, 0x5c, 0x2c, 0x21, 0xf9, 0x29, 0xf9,
	0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x4c, 0x99, 0x29,
	0x42, 0x52, 0x5c, 0x1c, 0x89, 0xc5, 0xc5, 0x99, 0xe9, 0x79, 0xa9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x08, 0x17, 0x6b, 0x49, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0x33,
	0x58, 0x02, 0xc2, 0x11, 0x12, 0xe2, 0x62, 0x49, 0xca, 0x4f, 0xa9, 0x94, 0x60, 0x01, 0x0b, 0x82,
	0xd9, 0x4a, 0x0a, 0x5c, 0x7c, 0xee, 0xa9, 0x25, 0x20, 0x0b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0xd0, 0xed, 0x31, 0xaa, 0xe7, 0xe2, 0x06, 0x49, 0xfb, 0x26, 0xe6, 0x25, 0xa6, 0xa7, 0x16,
	0x09, 0x99, 0x72, 0xb1, 0x43, 0x35, 0x08, 0x49, 0xea, 0xc1, 0xbd, 0xa0, 0x87, 0x6a, 0x88, 0x14,
	0x3f, 0x92, 0x14, 0x58, 0xad, 0x25, 0x17, 0x0f, 0x54, 0x89, 0x0f, 0x48, 0x50, 0x48, 0x4c, 0x0f,
	0xe2, 0x67, 0x3d, 0x98, 0x9f, 0xf5, 0x5c, 0x41, 0x7e, 0xc6, 0xd0, 0x68, 0xc0, 0x98, 0xc4, 0x06,
	0x56, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x91, 0x5a, 0x88, 0x43, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoManagerClient is the client API for TodoManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoManagerClient interface {
	GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*Todo, error)
	GetTodoLists(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (TodoManager_GetTodoListsClient, error)
}

type todoManagerClient struct {
	cc *grpc.ClientConn
}

func NewTodoManagerClient(cc *grpc.ClientConn) TodoManagerClient {
	return &todoManagerClient{cc}
}

func (c *todoManagerClient) GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todolists.TodoManager/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) GetTodoLists(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (TodoManager_GetTodoListsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TodoManager_serviceDesc.Streams[0], "/todolists.TodoManager/GetTodoLists", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoManagerGetTodoListsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoManager_GetTodoListsClient interface {
	Recv() (*Todo, error)
	grpc.ClientStream
}

type todoManagerGetTodoListsClient struct {
	grpc.ClientStream
}

func (x *todoManagerGetTodoListsClient) Recv() (*Todo, error) {
	m := new(Todo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoManagerServer is the server API for TodoManager service.
type TodoManagerServer interface {
	GetTodo(context.Context, *GetTodoRequest) (*Todo, error)
	GetTodoLists(*empty.Empty, TodoManager_GetTodoListsServer) error
}

// UnimplementedTodoManagerServer can be embedded to have forward compatible implementations.
type UnimplementedTodoManagerServer struct {
}

func (*UnimplementedTodoManagerServer) GetTodo(ctx context.Context, req *GetTodoRequest) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (*UnimplementedTodoManagerServer) GetTodoLists(req *empty.Empty, srv TodoManager_GetTodoListsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTodoLists not implemented")
}

func RegisterTodoManagerServer(s *grpc.Server, srv TodoManagerServer) {
	s.RegisterService(&_TodoManager_serviceDesc, srv)
}

func _TodoManager_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolists.TodoManager/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).GetTodo(ctx, req.(*GetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_GetTodoLists_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoManagerServer).GetTodoLists(m, &todoManagerGetTodoListsServer{stream})
}

type TodoManager_GetTodoListsServer interface {
	Send(*Todo) error
	grpc.ServerStream
}

type todoManagerGetTodoListsServer struct {
	grpc.ServerStream
}

func (x *todoManagerGetTodoListsServer) Send(m *Todo) error {
	return x.ServerStream.SendMsg(m)
}

var _TodoManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todolists.TodoManager",
	HandlerType: (*TodoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTodo",
			Handler:    _TodoManager_GetTodo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTodoLists",
			Handler:       _TodoManager_GetTodoLists_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo_lists.proto",
}
