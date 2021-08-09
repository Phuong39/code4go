// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

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

type MyOrderReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyOrderReq) Reset()         { *m = MyOrderReq{} }
func (m *MyOrderReq) String() string { return proto.CompactTextString(m) }
func (*MyOrderReq) ProtoMessage()    {}
func (*MyOrderReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *MyOrderReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyOrderReq.Unmarshal(m, b)
}
func (m *MyOrderReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyOrderReq.Marshal(b, m, deterministic)
}
func (m *MyOrderReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyOrderReq.Merge(m, src)
}
func (m *MyOrderReq) XXX_Size() int {
	return xxx_messageInfo_MyOrderReq.Size(m)
}
func (m *MyOrderReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MyOrderReq.DiscardUnknown(m)
}

var xxx_messageInfo_MyOrderReq proto.InternalMessageInfo

func (m *MyOrderReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type MyOrderReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 []*Order `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyOrderReply) Reset()         { *m = MyOrderReply{} }
func (m *MyOrderReply) String() string { return proto.CompactTextString(m) }
func (*MyOrderReply) ProtoMessage()    {}
func (*MyOrderReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *MyOrderReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyOrderReply.Unmarshal(m, b)
}
func (m *MyOrderReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyOrderReply.Marshal(b, m, deterministic)
}
func (m *MyOrderReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyOrderReply.Merge(m, src)
}
func (m *MyOrderReply) XXX_Size() int {
	return xxx_messageInfo_MyOrderReply.Size(m)
}
func (m *MyOrderReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MyOrderReply.DiscardUnknown(m)
}

var xxx_messageInfo_MyOrderReply proto.InternalMessageInfo

func (m *MyOrderReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *MyOrderReply) GetData() []*Order {
	if m != nil {
		return m.Data
	}
	return nil
}

type Order struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AmountMoney          float32  `protobuf:"fixed32,3,opt,name=amount_money,json=amountMoney,proto3" json:"amount_money,omitempty"`
	IsPay                bool     `protobuf:"varint,4,opt,name=is_pay,json=isPay,proto3" json:"is_pay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Order) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Order) GetAmountMoney() float32 {
	if m != nil {
		return m.AmountMoney
	}
	return 0
}

func (m *Order) GetIsPay() bool {
	if m != nil {
		return m.IsPay
	}
	return false
}

type InfoReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoReq) Reset()         { *m = InfoReq{} }
func (m *InfoReq) String() string { return proto.CompactTextString(m) }
func (*InfoReq) ProtoMessage()    {}
func (*InfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *InfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoReq.Unmarshal(m, b)
}
func (m *InfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoReq.Marshal(b, m, deterministic)
}
func (m *InfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoReq.Merge(m, src)
}
func (m *InfoReq) XXX_Size() int {
	return xxx_messageInfo_InfoReq.Size(m)
}
func (m *InfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_InfoReq proto.InternalMessageInfo

func (m *InfoReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type InfoReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	HeadImage            string   `protobuf:"bytes,4,opt,name=head_image,json=headImage,proto3" json:"head_image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoReply) Reset()         { *m = InfoReply{} }
func (m *InfoReply) String() string { return proto.CompactTextString(m) }
func (*InfoReply) ProtoMessage()    {}
func (*InfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *InfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoReply.Unmarshal(m, b)
}
func (m *InfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoReply.Marshal(b, m, deterministic)
}
func (m *InfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoReply.Merge(m, src)
}
func (m *InfoReply) XXX_Size() int {
	return xxx_messageInfo_InfoReply.Size(m)
}
func (m *InfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_InfoReply proto.InternalMessageInfo

func (m *InfoReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InfoReply) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *InfoReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InfoReply) GetHeadImage() string {
	if m != nil {
		return m.HeadImage
	}
	return ""
}

func init() {
	proto.RegisterType((*MyOrderReq)(nil), "client.MyOrderReq")
	proto.RegisterType((*MyOrderReply)(nil), "client.MyOrderReply")
	proto.RegisterType((*Order)(nil), "client.Order")
	proto.RegisterType((*InfoReq)(nil), "client.InfoReq")
	proto.RegisterType((*InfoReply)(nil), "client.InfoReply")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4e, 0xc3, 0x30,
	0x0c, 0xc6, 0xe9, 0xd6, 0xb5, 0xab, 0x37, 0x84, 0xb0, 0x40, 0x94, 0x49, 0x48, 0x5d, 0x24, 0xa4,
	0x1e, 0x50, 0x0f, 0x43, 0x3c, 0x00, 0x07, 0x0e, 0x3d, 0x4c, 0xa0, 0xbc, 0x40, 0x15, 0x96, 0x00,
	0x11, 0x6d, 0x52, 0xfa, 0xe7, 0x90, 0xb7, 0x47, 0x49, 0x37, 0x36, 0x84, 0xc6, 0xcd, 0xfe, 0xfc,
	0x25, 0xfe, 0xd9, 0x06, 0xe8, 0x5b, 0xd1, 0x64, 0x75, 0xa3, 0x3b, 0x8d, 0xc1, 0xa6, 0x94, 0x42,
	0x75, 0xe4, 0x16, 0x60, 0x6d, 0x9e, 0x1b, 0x2e, 0x1a, 0x2a, 0xbe, 0xf0, 0x0a, 0x42, 0xeb, 0x29,
	0x24, 0x8f, 0xbd, 0xc4, 0x4b, 0xc7, 0x34, 0xb0, 0x69, 0xce, 0xc9, 0x13, 0xcc, 0x7f, 0x6c, 0x75,
	0x69, 0x10, 0xc1, 0xdf, 0x68, 0x2e, 0x9c, 0x6b, 0x42, 0x5d, 0x8c, 0x4b, 0xf0, 0x39, 0xeb, 0x58,
	0x3c, 0x4a, 0xc6, 0xe9, 0x6c, 0x75, 0x9a, 0x0d, 0x1d, 0xb2, 0xe1, 0x95, 0x2b, 0x91, 0x1e, 0x26,
	0x2e, 0xc5, 0x6b, 0x98, 0x6a, 0x1b, 0xec, 0x3b, 0x85, 0x2e, 0xcf, 0xf9, 0x21, 0xc3, 0xe8, 0x90,
	0x01, 0x97, 0x30, 0x67, 0x95, 0xee, 0x55, 0x57, 0x54, 0x5a, 0x09, 0x13, 0x8f, 0x13, 0x2f, 0x1d,
	0xd1, 0xd9, 0xa0, 0xad, 0xad, 0x84, 0x97, 0x10, 0xc8, 0xb6, 0xa8, 0x99, 0x89, 0xfd, 0xc4, 0x4b,
	0xa7, 0x74, 0x22, 0xdb, 0x17, 0x66, 0x08, 0x81, 0x30, 0x57, 0x6f, 0xfa, 0xdf, 0x09, 0x3f, 0x21,
	0x1a, 0x3c, 0xc7, 0xc6, 0x3b, 0xca, 0x85, 0xe0, 0x2b, 0x56, 0x09, 0xc7, 0x13, 0x51, 0x17, 0xe3,
	0x0d, 0xc0, 0x87, 0x60, 0xbc, 0x90, 0x15, 0x7b, 0x17, 0x0e, 0x26, 0xa2, 0x91, 0x55, 0x72, 0x2b,
	0xac, 0xd4, 0xf0, 0xd7, 0x63, 0x2d, 0xf1, 0x01, 0xc2, 0xed, 0x66, 0x11, 0x77, 0x2b, 0xdb, 0x5f,
	0x64, 0x71, 0xf1, 0x47, 0xab, 0x4b, 0x43, 0x4e, 0xf0, 0x0e, 0x7c, 0x8b, 0x8b, 0x67, 0xbb, 0xfa,
	0x76, 0xc0, 0xc5, 0xf9, 0x6f, 0xc1, 0xb9, 0x5f, 0x03, 0x77, 0xf4, 0xfb, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x37, 0xae, 0xc1, 0x3a, 0x02, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserApiClient is the client API for UserApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserApiClient interface {
	// 我的订单
	MyOrder(ctx context.Context, in *MyOrderReq, opts ...grpc.CallOption) (*MyOrderReply, error)
	// 我的信息
	Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoReply, error)
}

type userApiClient struct {
	cc *grpc.ClientConn
}

func NewUserApiClient(cc *grpc.ClientConn) UserApiClient {
	return &userApiClient{cc}
}

func (c *userApiClient) MyOrder(ctx context.Context, in *MyOrderReq, opts ...grpc.CallOption) (*MyOrderReply, error) {
	out := new(MyOrderReply)
	err := c.cc.Invoke(ctx, "/client.userApi/MyOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := c.cc.Invoke(ctx, "/client.userApi/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserApiServer is the server API for UserApi service.
type UserApiServer interface {
	// 我的订单
	MyOrder(context.Context, *MyOrderReq) (*MyOrderReply, error)
	// 我的信息
	Info(context.Context, *InfoReq) (*InfoReply, error)
}

// UnimplementedUserApiServer can be embedded to have forward compatible implementations.
type UnimplementedUserApiServer struct {
}

func (*UnimplementedUserApiServer) MyOrder(ctx context.Context, req *MyOrderReq) (*MyOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MyOrder not implemented")
}
func (*UnimplementedUserApiServer) Info(ctx context.Context, req *InfoReq) (*InfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}

func RegisterUserApiServer(s *grpc.Server, srv UserApiServer) {
	s.RegisterService(&_UserApi_serviceDesc, srv)
}

func _UserApi_MyOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).MyOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.userApi/MyOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).MyOrder(ctx, req.(*MyOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApi_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.userApi/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).Info(ctx, req.(*InfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client.userApi",
	HandlerType: (*UserApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MyOrder",
			Handler:    _UserApi_MyOrder_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _UserApi_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}