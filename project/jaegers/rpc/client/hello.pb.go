// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package client

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

type SayHelloReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloReq) Reset()         { *m = SayHelloReq{} }
func (m *SayHelloReq) String() string { return proto.CompactTextString(m) }
func (*SayHelloReq) ProtoMessage()    {}
func (*SayHelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *SayHelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloReq.Unmarshal(m, b)
}
func (m *SayHelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloReq.Marshal(b, m, deterministic)
}
func (m *SayHelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloReq.Merge(m, src)
}
func (m *SayHelloReq) XXX_Size() int {
	return xxx_messageInfo_SayHelloReq.Size(m)
}
func (m *SayHelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloReq proto.InternalMessageInfo

func (m *SayHelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SayHelloReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloReply) Reset()         { *m = SayHelloReply{} }
func (m *SayHelloReply) String() string { return proto.CompactTextString(m) }
func (*SayHelloReply) ProtoMessage()    {}
func (*SayHelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *SayHelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloReply.Unmarshal(m, b)
}
func (m *SayHelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloReply.Marshal(b, m, deterministic)
}
func (m *SayHelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloReply.Merge(m, src)
}
func (m *SayHelloReply) XXX_Size() int {
	return xxx_messageInfo_SayHelloReply.Size(m)
}
func (m *SayHelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloReply proto.InternalMessageInfo

func (m *SayHelloReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SayHelloReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*SayHelloReq)(nil), "client.SayHelloReq")
	proto.RegisterType((*SayHelloReply)(nil), "client.SayHelloReply")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x51,
	0x52, 0xe4, 0xe2, 0x0e, 0x4e, 0xac, 0xf4, 0x00, 0xc9, 0x04, 0xa5, 0x16, 0x0a, 0x09, 0x71, 0xb1,
	0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0xe6, 0x5c,
	0xbc, 0x08, 0x25, 0x05, 0x39, 0x95, 0x20, 0x45, 0xc9, 0xf9, 0x29, 0x10, 0x45, 0xac, 0x41, 0x60,
	0x36, 0x48, 0x2c, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x09, 0xa2, 0x11, 0xc4, 0x36, 0x72, 0xe1, 0xe2,
	0x00, 0x5b, 0xe9, 0x58, 0x90, 0x29, 0x64, 0xc1, 0xc5, 0x01, 0x33, 0x44, 0x48, 0x58, 0x0f, 0x62,
	0xb9, 0x1e, 0x92, 0xcd, 0x52, 0xa2, 0x98, 0x82, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60,
	0x07, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x0f, 0xdd, 0x94, 0xbf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloApiClient is the client API for HelloApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloApiClient interface {
	SayHello(ctx context.Context, in *SayHelloReq, opts ...grpc.CallOption) (*SayHelloReply, error)
}

type helloApiClient struct {
	cc *grpc.ClientConn
}

func NewHelloApiClient(cc *grpc.ClientConn) HelloApiClient {
	return &helloApiClient{cc}
}

func (c *helloApiClient) SayHello(ctx context.Context, in *SayHelloReq, opts ...grpc.CallOption) (*SayHelloReply, error) {
	out := new(SayHelloReply)
	err := c.cc.Invoke(ctx, "/client.helloApi/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloApiServer is the server API for HelloApi service.
type HelloApiServer interface {
	SayHello(context.Context, *SayHelloReq) (*SayHelloReply, error)
}

// UnimplementedHelloApiServer can be embedded to have forward compatible implementations.
type UnimplementedHelloApiServer struct {
}

func (*UnimplementedHelloApiServer) SayHello(ctx context.Context, req *SayHelloReq) (*SayHelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterHelloApiServer(s *grpc.Server, srv HelloApiServer) {
	s.RegisterService(&_HelloApi_serviceDesc, srv)
}

func _HelloApi_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloApiServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.helloApi/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloApiServer).SayHello(ctx, req.(*SayHelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client.helloApi",
	HandlerType: (*HelloApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloApi_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}