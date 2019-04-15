// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package server

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

type Executable struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Args                 []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Executable) Reset()         { *m = Executable{} }
func (m *Executable) String() string { return proto.CompactTextString(m) }
func (*Executable) ProtoMessage()    {}
func (*Executable) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *Executable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Executable.Unmarshal(m, b)
}
func (m *Executable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Executable.Marshal(b, m, deterministic)
}
func (m *Executable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Executable.Merge(m, src)
}
func (m *Executable) XXX_Size() int {
	return xxx_messageInfo_Executable.Size(m)
}
func (m *Executable) XXX_DiscardUnknown() {
	xxx_messageInfo_Executable.DiscardUnknown(m)
}

var xxx_messageInfo_Executable proto.InternalMessageInfo

func (m *Executable) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *Executable) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type Response struct {
	ExitCode             int32    `protobuf:"varint,1,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	Stdout               []byte   `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr               []byte   `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *Response) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *Response) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func init() {
	proto.RegisterType((*Executable)(nil), "server.Executable")
	proto.RegisterType((*Response)(nil), "server.Response")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xcc, 0xb8, 0xb8,
	0x5c, 0x2b, 0x52, 0x93, 0x4b, 0x4b, 0x12, 0x93, 0x72, 0x52, 0x85, 0x44, 0xb8, 0x58, 0x93, 0x33,
	0x4a, 0xf3, 0xb2, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x20, 0x1c, 0x21, 0x21, 0x2e, 0x96,
	0xc4, 0xa2, 0xf4, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x29, 0x8c, 0x8b,
	0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8a, 0x8b, 0x23, 0xb5, 0x22, 0xb3,
	0xc4, 0x39, 0x3f, 0x25, 0x15, 0xac, 0x91, 0x35, 0x08, 0xce, 0x17, 0x12, 0xe3, 0x62, 0x2b, 0x2e,
	0x49, 0xc9, 0x2f, 0x2d, 0x91, 0x60, 0x02, 0x1b, 0x09, 0xe5, 0x41, 0xc5, 0x53, 0x8b, 0x8a, 0x24,
	0x98, 0xe1, 0xe2, 0xa9, 0x45, 0x45, 0x46, 0x16, 0x5c, 0x2c, 0x20, 0xf7, 0x08, 0x19, 0x40, 0x69,
	0x21, 0x3d, 0xa8, 0xb3, 0x11, 0xae, 0x94, 0x12, 0x80, 0x89, 0xc1, 0x5c, 0xa0, 0xc4, 0xa0, 0xc1,
	0x98, 0xc4, 0x06, 0xf6, 0x98, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x52, 0x5f, 0xcb, 0x33, 0xe8,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecClient is the client API for Exec service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecClient interface {
	Exec(ctx context.Context, opts ...grpc.CallOption) (Exec_ExecClient, error)
}

type execClient struct {
	cc *grpc.ClientConn
}

func NewExecClient(cc *grpc.ClientConn) ExecClient {
	return &execClient{cc}
}

func (c *execClient) Exec(ctx context.Context, opts ...grpc.CallOption) (Exec_ExecClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Exec_serviceDesc.Streams[0], "/server.Exec/Exec", opts...)
	if err != nil {
		return nil, err
	}
	x := &execExecClient{stream}
	return x, nil
}

type Exec_ExecClient interface {
	Send(*Executable) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type execExecClient struct {
	grpc.ClientStream
}

func (x *execExecClient) Send(m *Executable) error {
	return x.ClientStream.SendMsg(m)
}

func (x *execExecClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExecServer is the server API for Exec service.
type ExecServer interface {
	Exec(Exec_ExecServer) error
}

// UnimplementedExecServer can be embedded to have forward compatible implementations.
type UnimplementedExecServer struct {
}

func (*UnimplementedExecServer) Exec(srv Exec_ExecServer) error {
	return status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

func RegisterExecServer(s *grpc.Server, srv ExecServer) {
	s.RegisterService(&_Exec_serviceDesc, srv)
}

func _Exec_Exec_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExecServer).Exec(&execExecServer{stream})
}

type Exec_ExecServer interface {
	SendAndClose(*Response) error
	Recv() (*Executable, error)
	grpc.ServerStream
}

type execExecServer struct {
	grpc.ServerStream
}

func (x *execExecServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *execExecServer) Recv() (*Executable, error) {
	m := new(Executable)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Exec_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Exec",
	HandlerType: (*ExecServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Exec",
			Handler:       _Exec_Exec_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "server.proto",
}