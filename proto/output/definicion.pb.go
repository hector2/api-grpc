// Code generated by protoc-gen-go. DO NOT EDIT.
// source: definicion.proto

package definicion

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Id struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff1084f9dbe647f9, []int{0}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff1084f9dbe647f9, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Task struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Tags                 []*Tag   `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff1084f9dbe647f9, []int{2}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Task) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Task) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Tag struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff1084f9dbe647f9, []int{3}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Id)(nil), "definicion.Id")
	proto.RegisterType((*User)(nil), "definicion.User")
	proto.RegisterType((*Task)(nil), "definicion.Task")
	proto.RegisterType((*Tag)(nil), "definicion.Tag")
}

func init() { proto.RegisterFile("definicion.proto", fileDescriptor_ff1084f9dbe647f9) }

var fileDescriptor_ff1084f9dbe647f9 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0x87, 0x37, 0x69, 0x54, 0x76, 0xc4, 0x75, 0x19, 0xf6, 0x50, 0x7a, 0x5a, 0xe2, 0xa5, 0x1e,
	0x2c, 0xd2, 0x3e, 0x81, 0x27, 0x29, 0x78, 0xb1, 0xc6, 0x07, 0x88, 0x26, 0x96, 0x50, 0x6d, 0x25,
	0x09, 0xbe, 0x90, 0x2f, 0x2a, 0x4d, 0x91, 0xfe, 0xdb, 0x5b, 0x92, 0xf9, 0xe6, 0xf7, 0x4d, 0x06,
	0xf6, 0x4a, 0x7f, 0x98, 0xd6, 0xbc, 0x9b, 0xae, 0xcd, 0xbe, 0x6d, 0xe7, 0x3b, 0x84, 0xf1, 0x85,
	0x1f, 0x80, 0x96, 0x0a, 0x77, 0x40, 0x8d, 0x8a, 0xc9, 0x91, 0xa4, 0x57, 0x15, 0x35, 0x8a, 0x27,
	0xc0, 0x5e, 0x9d, 0xb6, 0x88, 0xc0, 0x5a, 0xf9, 0xa5, 0x43, 0x65, 0x5b, 0x85, 0x33, 0x7f, 0x06,
	0x26, 0xa4, 0x6b, 0x96, 0x3d, 0x78, 0x80, 0x33, 0x6f, 0xfc, 0xa7, 0x8e, 0x69, 0x80, 0x87, 0x0b,
	0xde, 0x00, 0xf3, 0xb2, 0x76, 0x71, 0x74, 0x8c, 0xd2, 0xcb, 0xfc, 0x3a, 0x9b, 0x0c, 0x23, 0x64,
	0x5d, 0x85, 0x22, 0xbf, 0x85, 0x48, 0xc8, 0x7a, 0x95, 0xf8, 0x6f, 0xa7, 0xa3, 0x3d, 0xff, 0x25,
	0xb0, 0x7b, 0xd1, 0xf6, 0xa7, 0x8f, 0x10, 0xd2, 0x6a, 0xe9, 0xf0, 0x0e, 0x2e, 0x1e, 0xb5, 0x1f,
	0x66, 0x9a, 0xe6, 0x97, 0x2a, 0xd9, 0xcf, 0x7d, 0xae, 0xe1, 0x1b, 0x2c, 0x60, 0xfb, 0x64, 0x5c,
	0xe0, 0x1d, 0xce, 0x80, 0xfe, 0xcb, 0xa7, 0x5a, 0xee, 0x49, 0xef, 0x78, 0x50, 0x2a, 0x38, 0x56,
	0x40, 0xb2, 0xb0, 0xf2, 0xcd, 0xdb, 0x79, 0x58, 0x74, 0xf1, 0x17, 0x00, 0x00, 0xff, 0xff, 0x59,
	0xc2, 0x25, 0x67, 0x7c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServicioTareasClient is the client API for ServicioTareas service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServicioTareasClient interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetTask(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Task, error)
	ListTasks(ctx context.Context, in *User, opts ...grpc.CallOption) (ServicioTareas_ListTasksClient, error)
	AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Id, error)
}

type servicioTareasClient struct {
	cc *grpc.ClientConn
}

func NewServicioTareasClient(cc *grpc.ClientConn) ServicioTareasClient {
	return &servicioTareasClient{cc}
}

func (c *servicioTareasClient) GetTask(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/definicion.ServicioTareas/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioTareasClient) ListTasks(ctx context.Context, in *User, opts ...grpc.CallOption) (ServicioTareas_ListTasksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServicioTareas_serviceDesc.Streams[0], "/definicion.ServicioTareas/ListTasks", opts...)
	if err != nil {
		return nil, err
	}
	x := &servicioTareasListTasksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServicioTareas_ListTasksClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type servicioTareasListTasksClient struct {
	grpc.ClientStream
}

func (x *servicioTareasListTasksClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *servicioTareasClient) AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/definicion.ServicioTareas/AddTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicioTareasServer is the server API for ServicioTareas service.
type ServicioTareasServer interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetTask(context.Context, *Id) (*Task, error)
	ListTasks(*User, ServicioTareas_ListTasksServer) error
	AddTask(context.Context, *Task) (*Id, error)
}

// UnimplementedServicioTareasServer can be embedded to have forward compatible implementations.
type UnimplementedServicioTareasServer struct {
}

func (*UnimplementedServicioTareasServer) GetTask(ctx context.Context, req *Id) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (*UnimplementedServicioTareasServer) ListTasks(req *User, srv ServicioTareas_ListTasksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (*UnimplementedServicioTareasServer) AddTask(ctx context.Context, req *Task) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTask not implemented")
}

func RegisterServicioTareasServer(s *grpc.Server, srv ServicioTareasServer) {
	s.RegisterService(&_ServicioTareas_serviceDesc, srv)
}

func _ServicioTareas_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioTareasServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definicion.ServicioTareas/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioTareasServer).GetTask(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioTareas_ListTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServicioTareasServer).ListTasks(m, &servicioTareasListTasksServer{stream})
}

type ServicioTareas_ListTasksServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type servicioTareasListTasksServer struct {
	grpc.ServerStream
}

func (x *servicioTareasListTasksServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func _ServicioTareas_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioTareasServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definicion.ServicioTareas/AddTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioTareasServer).AddTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServicioTareas_serviceDesc = grpc.ServiceDesc{
	ServiceName: "definicion.ServicioTareas",
	HandlerType: (*ServicioTareasServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTask",
			Handler:    _ServicioTareas_GetTask_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _ServicioTareas_AddTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTasks",
			Handler:       _ServicioTareas_ListTasks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "definicion.proto",
}